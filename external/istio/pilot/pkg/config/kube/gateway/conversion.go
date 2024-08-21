// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gateway

import (
	"fmt"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	klabels "k8s.io/apimachinery/pkg/labels"
	gatewayapiV1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gatewayapiV1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	istio "istio.io/api/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/features"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/model/credentials"
	"istio.io/istio/pilot/pkg/model/kstatus"
	"istio.io/istio/pilot/pkg/util/sets"
	"istio.io/istio/pkg/config"
	"istio.io/istio/pkg/config/constants"
	"istio.io/istio/pkg/config/host"
	"istio.io/istio/pkg/config/schema/gvk"
)

const (
	DefaultClassName  = "higress"
	ControllerName    = "higress.io/gateway-controller"
	GATEWAY_CLASS_ENV = "GATEWAY_CLASS_ENV"
)

// KubernetesResources stores all inputs to our conversion
type KubernetesResources struct {
	GatewayClass   []config.Config
	Gateway        []config.Config
	HTTPRoute      []config.Config
	TCPRoute       []config.Config
	TLSRoute       []config.Config
	ReferenceGrant []config.Config
	// Namespaces stores all namespace in the cluster, keyed by name
	Namespaces map[string]*corev1.Namespace

	// Domain for the cluster. Typically, cluster.local
	Domain  string
	Context model.GatewayContext
}

// OutputResources stores all outputs of our conversion
type OutputResources struct {
	Gateway        []config.Config
	VirtualService []config.Config
	// AllowedReferences stores all allowed references, from Reference -> to Reference(s)
	AllowedReferences map[Reference]map[Reference]struct{}
	// ReferencedNamespaceKeys stores the label key of all namespace selections. This allows us to quickly
	// determine if a namespace update could have impacted any Gateways. See namespaceEvent.
	ReferencedNamespaceKeys sets.Set
}

// Reference stores a reference to a namespaced GVK, as used by ReferenceGrant
type Reference struct {
	Kind      config.GroupVersionKind
	Namespace gatewayapiV1beta1.Namespace
}

// convertResources is the top level entrypoint to our conversion logic, computing the full state based
// on KubernetesResources inputs.
func convertResources(r *KubernetesResources) OutputResources {
	result := OutputResources{}
	gw, gwMap, nsReferences := convertGateways(r)
	result.Gateway = gw
	result.VirtualService = convertVirtualService(r, gwMap)

	// Once we have gone through all route computation, we will know how many routes bound to each gateway.
	// Report this in the status.
	for _, dm := range gwMap {
		for _, pri := range dm {
			if pri.ReportAttachedRoutes != nil {
				pri.ReportAttachedRoutes()
			}
		}
	}
	result.AllowedReferences = convertReferencePolicies(r)
	result.ReferencedNamespaceKeys = nsReferences
	return result
}

// convertReferencePolicies extracts all ReferenceGrant into an easily accessibly index.
// The currently supported references are:
// * Gateway -> Secret
func convertReferencePolicies(r *KubernetesResources) map[Reference]map[Reference]struct{} {
	// TODO support Name in ReferenceGrantTo
	res := map[Reference]map[Reference]struct{}{}
	for _, obj := range r.ReferenceGrant {
		rp := obj.Spec.(*gatewayapiV1alpha2.ReferenceGrantSpec)
		for _, from := range rp.From {
			fromKey := Reference{
				Namespace: gatewayapiV1beta1.Namespace(from.Namespace),
			}
			if string(from.Group) == gvk.KubernetesGateway.Group && string(from.Kind) == gvk.KubernetesGateway.Kind {
				fromKey.Kind = gvk.KubernetesGateway
			} else {
				// Not supported type. Not an error; may be for another controller
				continue
			}
			for _, to := range rp.To {
				toKey := Reference{
					Namespace: gatewayapiV1beta1.Namespace(from.Namespace),
				}
				if to.Group == "" && string(to.Kind) == gvk.Secret.Kind {
					toKey.Kind = gvk.Secret
				} else {
					// Not supported type. Not an error; may be for another controller
					continue
				}
				if _, f := res[fromKey]; !f {
					res[fromKey] = map[Reference]struct{}{}
				}
				res[fromKey][toKey] = struct{}{}
			}
		}
	}
	return res
}

// convertVirtualService takes all xRoute types and generates corresponding VirtualServices.
func convertVirtualService(r *KubernetesResources, gatewayMap map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo) []config.Config {
	result := []config.Config{}
	for _, obj := range r.TCPRoute {
		if vsConfig := buildTCPVirtualService(obj, gatewayMap, r.Domain); vsConfig != nil {
			result = append(result, *vsConfig)
		}
	}

	for _, obj := range r.TLSRoute {
		if vsConfig := buildTLSVirtualService(obj, gatewayMap, r.Domain); vsConfig != nil {
			result = append(result, *vsConfig)
		}
	}

	// for gateway routes, build one VS per gateway+host
	gatewayRoutes := make(map[string]map[string]*config.Config)

	for _, obj := range r.HTTPRoute {
		buildHTTPVirtualServices(r, obj, gatewayMap, gatewayRoutes, r.Domain)
	}
	for _, vsByHost := range gatewayRoutes {
		for _, vsConfig := range vsByHost {
			result = append(result, *vsConfig)
		}
	}
	return result
}

// isCatchAll returns true if HTTPMatchRequest is a catchall match otherwise
// false. Note - this may not be exactly "catch all" as we don't know the full
// class of possible inputs As such, this is used only for optimization.
func isCatchAllMatch(m *istio.HTTPMatchRequest) bool {
	catchall := false
	if m.Uri != nil {
		switch m := m.Uri.MatchType.(type) {
		case *istio.StringMatch_Prefix:
			catchall = m.Prefix == "/"
		case *istio.StringMatch_Regex:
			catchall = m.Regex == "*"
		}
	}
	// A Match is catch all if and only if it has no match set
	// and URI has a prefix / or regex *.
	return catchall &&
		len(m.Headers) == 0 &&
		len(m.QueryParams) == 0 &&
		len(m.SourceLabels) == 0 &&
		len(m.WithoutHeaders) == 0 &&
		len(m.Gateways) == 0 &&
		m.Method == nil &&
		m.Scheme == nil &&
		m.Port == 0 &&
		m.Authority == nil &&
		m.SourceNamespace == ""
}

// getURIRank ranks a URI match type. Exact > Prefix > Regex
func getURIRank(match *istio.HTTPMatchRequest) int {
	if match.Uri == nil {
		return -1
	}
	switch match.Uri.MatchType.(type) {
	case *istio.StringMatch_Exact:
		return 3
	case *istio.StringMatch_Prefix:
		return 2
	case *istio.StringMatch_Regex:
		// TODO optimize in new verison envoy
		if strings.HasSuffix(match.Uri.GetRegex(), prefixMatchRegex) &&
			!strings.ContainsAny(strings.TrimSuffix(match.Uri.GetRegex(), prefixMatchRegex), `\.+*?()|[]{}^$`) {
			return 2
		}
		return 1
	}
	// should not happen
	return -1
}

func getURILength(match *istio.HTTPMatchRequest) int {
	if match.Uri == nil {
		return 0
	}
	switch match.Uri.MatchType.(type) {
	case *istio.StringMatch_Prefix:
		return len(match.Uri.GetPrefix())
	case *istio.StringMatch_Exact:
		return len(match.Uri.GetExact())
	case *istio.StringMatch_Regex:
		return len(match.Uri.GetRegex())
	}
	// should not happen
	return -1
}

// sortHTTPRoutes sorts generated vs routes to meet gateway-api requirements
// see https://gateway-api.sigs.k8s.io/v1alpha2/references/spec/#gateway.networking.k8s.io/v1alpha2.HTTPRouteRule
func sortHTTPRoutes(routes []*istio.HTTPRoute) {
	sort.SliceStable(routes, func(i, j int) bool {
		if len(routes[i].Match) == 0 {
			return false
		} else if len(routes[j].Match) == 0 {
			return true
		}
		if isCatchAllMatch(routes[i].Match[0]) {
			return false
		} else if isCatchAllMatch(routes[j].Match[0]) {
			return true
		}
		// Only look at match[0], we always generate only one match
		m1, m2 := routes[i].Match[0], routes[j].Match[0]
		r1, r2 := getURIRank(m1), getURIRank(m2)
		len1, len2 := getURILength(m1), getURILength(m2)
		switch {
		// 1: Exact/Prefix/Regex
		case r1 != r2:
			return r1 > r2
		case len1 != len2:
			return len1 > len2
			// 2: method math
		case (m1.Method == nil) != (m2.Method == nil):
			return m1.Method != nil
			// 3: number of header matches
		case len(m1.Headers) != len(m2.Headers):
			return len(m1.Headers) > len(m2.Headers)
			// 4: number of query matches
		default:
			return len(m1.QueryParams) > len(m2.QueryParams)
		}
	})
}

func routeMeta(obj config.Config) map[string]string {
	m := parentMeta(obj, nil)
	m[constants.InternalRouteSemantics] = constants.RouteSemanticsGateway
	return m
}

func filteredReferences(parents []routeParentReference) []routeParentReference {
	ret := make([]routeParentReference, 0, len(parents))
	for _, p := range parents {
		if p.DeniedReason != nil {
			// We should filter this out
			continue
		}
		ret = append(ret, p)
	}
	// To ensure deterministic order, sort them
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].InternalName < ret[j].InternalName
	})
	return ret
}

// Added by ingress
func generateRouteName(obj config.Config) string {
	if obj.Namespace == features.HigressSystemNs {
		return obj.Name
	}
	return path.Join(obj.Namespace, obj.Name)
}

// End added by ingress

func buildHTTPVirtualServices(ctx *KubernetesResources, obj config.Config, gateways map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo, gatewayRoutes map[string]map[string]*config.Config, domain string) {
	route := obj.Spec.(*gatewayapiV1beta1.HTTPRouteSpec)

	parentRefs := extractParentReferenceInfo(gateways, route.ParentRefs, route.Hostnames, gvk.HTTPRoute, obj.Namespace)

	reportError := func(routeErr *ConfigError) {
		obj.Status.(*kstatus.WrappedStatus).Mutate(func(s config.Status) config.Status {
			rs := s.(*gatewayapiV1beta1.HTTPRouteStatus)
			rs.Parents = createRouteStatus(parentRefs, obj, rs.Parents, routeErr)
			return rs
		})
	}

	httproutes := []*istio.HTTPRoute{}
	for _, r := range route.Rules {
		// TODO: implement rewrite, timeout, mirror, corspolicy, retries
		vs := &istio.HTTPRoute{
			Name: generateRouteName(obj),
		}
		for _, match := range r.Matches {
			uri, err := createURIMatch(match)
			if err != nil {
				reportError(err)
				return
			}
			headers, err := createHeadersMatch(match)
			if err != nil {
				reportError(err)
				return
			}
			qp, err := createQueryParamsMatch(match)
			if err != nil {
				reportError(err)
				return
			}
			method, err := createMethodMatch(match)
			if err != nil {
				reportError(err)
				return
			}
			vs.Match = append(vs.Match, &istio.HTTPMatchRequest{
				Uri:         uri,
				Headers:     headers,
				QueryParams: qp,
				Method:      method,
			})
		}
		for _, filter := range r.Filters {
			switch filter.Type {
			case gatewayapiV1beta1.HTTPRouteFilterRequestHeaderModifier:
				vs.Headers = createHeadersFilter(filter.RequestHeaderModifier)
			case gatewayapiV1beta1.HTTPRouteFilterRequestRedirect:
				vs.Redirect = createRedirectFilter(filter.RequestRedirect)
			case gatewayapiV1beta1.HTTPRouteFilterRequestMirror:
				mirror, err := createMirrorFilter(filter.RequestMirror, obj.Namespace, domain)
				if err != nil {
					reportError(err)
					return
				}
				vs.Mirror = mirror
			case gatewayapiV1beta1.HTTPRouteFilterURLRewrite:
				vs.Rewrite = createRewriteFilter(filter.URLRewrite)
			default:
				reportError(&ConfigError{
					Reason:  InvalidFilter,
					Message: fmt.Sprintf("unsupported filter type %q", filter.Type),
				})
				return
			}
		}

		zero := true
		for _, w := range r.BackendRefs {
			if w.Weight == nil || (w.Weight != nil && int(*w.Weight) != 0) {
				zero = false
				break
			}
		}
		if zero && vs.Redirect == nil {
			// The spec requires us to 503 when there are no >0 weight backends
			vs.Fault = &istio.HTTPFaultInjection{Abort: &istio.HTTPFaultInjection_Abort{
				Percentage: &istio.Percent{
					Value: 100,
				},
				ErrorType: &istio.HTTPFaultInjection_Abort_HttpStatus{
					HttpStatus: 503,
				},
			}}
		}

		fallbackCluster := obj.Annotations["higress.io/fallback-service"]

		route, err := buildHTTPDestination(r.BackendRefs, obj.Namespace, domain, zero, fallbackCluster)
		if err != nil {
			reportError(err)
			return
		}
		vs.Route = route

		httproutes = append(httproutes, vs)
	}
	reportError(nil)

	count := 0
	for _, parent := range filteredReferences(parentRefs) {
		// for gateway routes, build one VS per gateway+host
		routeMap := gatewayRoutes
		routeKey := parent.InternalName
		vsHosts := hostnameToStringList(route.Hostnames)
		routes := httproutes
		if len(routes) == 0 {
			continue
		}
		if _, f := routeMap[routeKey]; !f {
			routeMap[routeKey] = make(map[string]*config.Config)
		}

		// Create one VS per hostname with a single hostname.
		// This ensures we can treat each hostname independently, as the spec requires
		for _, h := range vsHosts {
			if cfg := routeMap[routeKey][h]; cfg != nil {
				// merge http routes
				vs := cfg.Spec.(*istio.VirtualService)
				vs.Http = append(vs.Http, routes...)
				// append parents
				cfg.Annotations[constants.InternalParentNames] = fmt.Sprintf("%s,%s/%s.%s",
					cfg.Annotations[constants.InternalParentNames], obj.GroupVersionKind.Kind, obj.Name, obj.Namespace)
			} else {
				name := fmt.Sprintf("%s-%d-%s", obj.Name, count, constants.KubernetesGatewayName)
				routeMap[routeKey][h] = &config.Config{
					Meta: config.Meta{
						CreationTimestamp: obj.CreationTimestamp,
						GroupVersionKind:  gvk.VirtualService,
						Name:              name,
						Annotations:       routeMeta(obj),
						Namespace:         obj.Namespace,
						Domain:            ctx.Domain,
					},
					Spec: &istio.VirtualService{
						Hosts:    []string{h},
						Gateways: []string{parent.InternalName},
						Http:     routes,
					},
				}
				count++
			}
		}
	}
	for _, vsByHost := range gatewayRoutes {
		for _, cfg := range vsByHost {
			vs := cfg.Spec.(*istio.VirtualService)
			sortHTTPRoutes(vs.Http)
		}
	}
}

func createRewriteFilter(filter *gatewayapiV1beta1.HTTPURLRewriteFilter) *istio.HTTPRewrite {
	if filter == nil {
		return nil
	}
	rewrite := &istio.HTTPRewrite{}
	if filter.Path != nil {
		switch filter.Path.Type {
		case gatewayapiV1beta1.PrefixMatchHTTPPathModifier:
			rewrite.Uri = *filter.Path.ReplacePrefixMatch
		case gatewayapiV1beta1.FullPathHTTPPathModifier:
			rewrite.UriRegex = &istio.RegexMatchAndSubstitute{
				Pattern:      "/.*",
				Substitution: *filter.Path.ReplaceFullPath,
			}
		}
	}
	if filter.Hostname != nil {
		rewrite.Authority = string(*filter.Hostname)
	}
	// Nothing done
	if rewrite.Uri == "" && rewrite.UriRegex == nil && rewrite.Authority == "" {
		return nil
	}
	return rewrite
}

func parentMeta(obj config.Config, sectionName *gatewayapiV1beta1.SectionName) map[string]string {
	name := fmt.Sprintf("%s/%s.%s", obj.GroupVersionKind.Kind, obj.Name, obj.Namespace)
	if sectionName != nil {
		name = fmt.Sprintf("%s/%s/%s.%s", obj.GroupVersionKind.Kind, obj.Name, *sectionName, obj.Namespace)
	}
	return map[string]string{
		constants.InternalParentNames: name,
	}
}

func hostnameToStringList(h []gatewayapiV1beta1.Hostname) []string {
	// In the Istio API, empty hostname is not allowed. In the Kubernetes API hosts means "any"
	if len(h) == 0 {
		return []string{"*"}
	}
	res := make([]string, 0, len(h))
	for _, i := range h {
		res = append(res, string(i))
	}
	return res
}

func toInternalParentReference(p gatewayapiV1beta1.ParentReference, localNamespace string) (parentKey, error) {
	empty := parentKey{}
	grp := defaultIfNil((*string)(p.Group), gvk.KubernetesGateway.Group)
	kind := defaultIfNil((*string)(p.Kind), gvk.KubernetesGateway.Kind)
	var ik config.GroupVersionKind
	var ns string
	// Currently supported types are Gateway and Mesh
	if kind == gvk.KubernetesGateway.Kind && grp == gvk.KubernetesGateway.Group {
		// Unset namespace means "same namespace"
		ns = defaultIfNil((*string)(p.Namespace), localNamespace)
		ik = gvk.KubernetesGateway
	} else if kind == meshGVK.Kind && grp == meshGVK.Group {
		ik = meshGVK
	} else {
		return empty, fmt.Errorf("unsupported parentKey: %v/%v", grp, kind)
	}
	return parentKey{
		Kind:      ik,
		Name:      string(p.Name),
		Namespace: ns,
	}, nil
}

func referenceAllowed(p *parentInfo, routeKind config.GroupVersionKind, parentKind config.GroupVersionKind, hostnames []gatewayapiV1beta1.Hostname, namespace string) error {
	// First check the hostnames are a match. This is a bi-directional wildcard match. Only one route
	// hostname must match for it to be allowed (but the others will be filtered at runtime)
	// If either is empty its treated as a wildcard which always matches

	if len(hostnames) == 0 {
		hostnames = []gatewayapiV1beta1.Hostname{"*"}
	}
	if len(p.Hostnames) > 0 {
		// TODO: the spec actually has a label match, not a string match. That is, *.com does not match *.apple.com
		// We are doing a string match here
		matched := false
		hostMatched := false
		for _, routeHostname := range hostnames {
			for _, parentHostNamespace := range p.Hostnames {
				spl := strings.Split(parentHostNamespace, "/")
				parentNamespace, parentHostname := spl[0], spl[1]
				hostnameMatch := host.Name(parentHostname).Matches(host.Name(routeHostname))
				namespaceMatch := parentNamespace == "*" || parentNamespace == namespace
				hostMatched = hostMatched || hostnameMatch
				if hostnameMatch && namespaceMatch {
					matched = true
					break
				}
			}
		}
		if !matched {
			if hostMatched {
				return fmt.Errorf("hostnames matched parent hostname %q, but namespace %q is not allowed by the parent", p.OriginalHostname, namespace)
			}
			return fmt.Errorf("no hostnames matched parent hostname %q", p.OriginalHostname)
		}
	}
	// Also make sure this route kind is allowed
	if len(p.AllowedKinds) > 0 {
		matched := false
		for _, ak := range p.AllowedKinds {
			if string(ak.Kind) == routeKind.Kind && defaultIfNil((*string)(ak.Group), gvk.GatewayClass.Group) == routeKind.Group {
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Errorf("kind %v is not allowed", routeKind)
		}
	}

	if parentKind == meshGVK {
		for _, h := range hostnames {
			if h == "*" {
				return fmt.Errorf("mesh requires hostname to be set")
			}
		}
	}
	return nil
}

func BackendRefsBeta1ToAlpha1(BackendRefs []gatewayapiV1alpha2.BackendRef) []gatewayapiV1beta1.BackendRef {
	newbakends := []gatewayapiV1beta1.BackendRef{}
	for _, bk := range BackendRefs {
		newref := gatewayapiV1beta1.BackendRef{
			BackendObjectReference: BackendObjectReferenceBeta1ToAlpha1(bk.BackendObjectReference),
			Weight:                 bk.Weight,
		}
		newbakends = append(newbakends, newref)
	}
	return newbakends
}
func BackendObjectReferenceBeta1ToAlpha1(backend gatewayapiV1alpha2.BackendObjectReference) gatewayapiV1beta1.BackendObjectReference {
	return gatewayapiV1beta1.BackendObjectReference{
		Group:     (*gatewayapiV1beta1.Group)(backend.Group),
		Kind:      (*gatewayapiV1beta1.Kind)(backend.Kind),
		Name:      gatewayapiV1beta1.ObjectName(backend.Name),
		Namespace: (*gatewayapiV1beta1.Namespace)(backend.Namespace),
		Port:      (*gatewayapiV1beta1.PortNumber)(backend.Port),
	}
}
func beta1ToAlpha1(routeRefs []gatewayapiV1alpha2.ParentReference) []gatewayapiV1beta1.ParentReference {
	newRefs := []gatewayapiV1beta1.ParentReference{}
	for _, routeref := range routeRefs {
		newref := parentRefAlpha2Tobeta1(routeref)
		newRefs = append(newRefs, newref)
	}
	return newRefs
}
func parentRefAlpha2Tobeta1(routeRef gatewayapiV1alpha2.ParentReference) gatewayapiV1beta1.ParentReference {
	newref := gatewayapiV1beta1.ParentReference{
		Group:       (*gatewayapiV1beta1.Group)(routeRef.Group),
		Kind:        (*gatewayapiV1beta1.Kind)(routeRef.Kind),
		Namespace:   (*gatewayapiV1beta1.Namespace)(routeRef.Namespace),
		Name:        gatewayapiV1beta1.ObjectName(routeRef.Name),
		SectionName: (*gatewayapiV1beta1.SectionName)(routeRef.SectionName),
		Port:        (*gatewayapiV1beta1.PortNumber)(routeRef.Port),
	}
	return newref
}
func parentRefBeta1ToAlpha2(routeRef gatewayapiV1beta1.ParentReference) gatewayapiV1alpha2.ParentReference {
	newref := gatewayapiV1alpha2.ParentReference{
		Group:       (*gatewayapiV1alpha2.Group)(routeRef.Group),
		Kind:        (*gatewayapiV1alpha2.Kind)(routeRef.Kind),
		Namespace:   (*gatewayapiV1alpha2.Namespace)(routeRef.Namespace),
		Name:        gatewayapiV1alpha2.ObjectName(routeRef.Name),
		SectionName: (*gatewayapiV1alpha2.SectionName)(routeRef.SectionName),
		Port:        (*gatewayapiV1alpha2.PortNumber)(routeRef.Port),
	}
	return newref
}

func statusRoutesAlpha2ToBeta1(routeRefs []gatewayapiV1alpha2.RouteParentStatus) []gatewayapiV1beta1.RouteParentStatus {
	newRefs := []gatewayapiV1beta1.RouteParentStatus{}
	for _, routeref := range routeRefs {
		newref := statusParentRefAlpha2Tobeta1(routeref)
		newRefs = append(newRefs, newref)
	}
	return newRefs
}
func statusParentRefAlpha2Tobeta1(routeParentStatus gatewayapiV1alpha2.RouteParentStatus) gatewayapiV1beta1.RouteParentStatus {
	parentStatus := gatewayapiV1beta1.RouteParentStatus{
		ParentRef:      parentRefAlpha2Tobeta1(routeParentStatus.ParentRef),
		ControllerName: gatewayapiV1beta1.GatewayController(routeParentStatus.ControllerName),
		Conditions:     routeParentStatus.Conditions,
	}
	return parentStatus
}
func statusRoutesBeta1ToAlpha1(routeRefs []gatewayapiV1beta1.RouteParentStatus) []gatewayapiV1alpha2.RouteParentStatus {
	newRefs := []gatewayapiV1alpha2.RouteParentStatus{}
	for _, routeref := range routeRefs {
		newref := gatewayapiV1alpha2.RouteParentStatus{
			ParentRef:      parentRefBeta1ToAlpha2(routeref.ParentRef),
			ControllerName: gatewayapiV1alpha2.GatewayController(routeref.ControllerName),
			Conditions:     routeref.Conditions,
		}
		newRefs = append(newRefs, newref)
	}
	return newRefs
}

func extractParentReferenceInfo(gateways map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo, routeRefs []gatewayapiV1beta1.ParentReference,
	hostnames []gatewayapiV1beta1.Hostname, kind config.GroupVersionKind, localNamespace string) []routeParentReference {
	parentRefs := []routeParentReference{}
	for _, ref := range routeRefs {
		ir, err := toInternalParentReference(ref, localNamespace)
		if err != nil {
			// Cannot handle the reference. Maybe it is for another controller, so we just ignore it
			continue
		}
		appendParent := func(pr *parentInfo, pk parentKey) {
			rpi := routeParentReference{
				InternalName:      pr.InternalName,
				DeniedReason:      referenceAllowed(pr, kind, pk.Kind, hostnames, localNamespace),
				OriginalReference: ref,
			}
			if rpi.DeniedReason == nil {
				// Record that we were able to bind to the parent
				pr.AttachedRoutes++
			}
			parentRefs = append(parentRefs, rpi)
		}
		if ref.SectionName != nil {
			// We are selecting a specific section, so attach just that section
			if pr, f := gateways[ir][*ref.SectionName]; f {
				appendParent(pr, ir)
			}
		} else {
			// no section name set, match all sections
			for _, pr := range gateways[ir] {
				appendParent(pr, ir)
			}
		}
	}
	return parentRefs
}

func buildTCPVirtualService(obj config.Config, gateways map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo, domain string) *config.Config {
	route := obj.Spec.(*gatewayapiV1alpha2.TCPRouteSpec)

	parentRefs := extractParentReferenceInfo(gateways, beta1ToAlpha1(route.ParentRefs), nil, gvk.TCPRoute, obj.Namespace)

	reportError := func(routeErr *ConfigError) {
		obj.Status.(*kstatus.WrappedStatus).Mutate(func(s config.Status) config.Status {
			rs := s.(*gatewayapiV1alpha2.TCPRouteStatus)
			rs.Parents = statusRoutesBeta1ToAlpha1(createRouteStatus(parentRefs, obj, statusRoutesAlpha2ToBeta1(rs.Parents), routeErr))
			return rs
		})
	}
	gatewayNames := referencesToInternalNames(parentRefs)
	if len(gatewayNames) == 0 {
		reportError(nil)
		return nil
	}

	routes := []*istio.TCPRoute{}
	for _, r := range route.Rules {
		route, err := buildTCPDestination(BackendRefsBeta1ToAlpha1(r.BackendRefs), obj.Namespace, domain)
		if err != nil {
			reportError(err)
			return nil
		}
		ir := &istio.TCPRoute{
			Route: route,
		}
		routes = append(routes, ir)
	}

	reportError(nil)
	vsConfig := config.Config{
		Meta: config.Meta{
			CreationTimestamp: obj.CreationTimestamp,
			GroupVersionKind:  gvk.VirtualService,
			Name:              fmt.Sprintf("%s-tcp-%s", obj.Name, constants.KubernetesGatewayName),
			Annotations:       parentMeta(obj, nil),
			Namespace:         obj.Namespace,
			Domain:            domain,
		},
		Spec: &istio.VirtualService{
			// We can use wildcard here since each listener can have at most one route bound to it, so we have
			// a single VS per Gateway.
			Hosts:    []string{"*"},
			Gateways: gatewayNames,
			Tcp:      routes,
		},
	}
	return &vsConfig
}

func buildTLSVirtualService(obj config.Config, gateways map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo, domain string) *config.Config {
	route := obj.Spec.(*gatewayapiV1alpha2.TLSRouteSpec)

	parentRefs := extractParentReferenceInfo(gateways, beta1ToAlpha1(route.ParentRefs), nil, gvk.TLSRoute, obj.Namespace)

	reportError := func(routeErr *ConfigError) {
		obj.Status.(*kstatus.WrappedStatus).Mutate(func(s config.Status) config.Status {
			rs := s.(*gatewayapiV1alpha2.TLSRouteStatus)
			rs.Parents = statusRoutesBeta1ToAlpha1(createRouteStatus(parentRefs, obj, statusRoutesAlpha2ToBeta1(rs.Parents), routeErr))
			return rs
		})
	}

	routes := []*istio.TLSRoute{}
	for _, r := range route.Rules {
		dest, err := buildTCPDestination(BackendRefsBeta1ToAlpha1(r.BackendRefs), obj.Namespace, domain)
		if err != nil {
			reportError(err)
			return nil
		}
		if len(dest) == 0 {
			return nil
		}
		ir := &istio.TLSRoute{
			Match: buildTLSMatch(route.Hostnames),
			Route: dest,
		}
		routes = append(routes, ir)
	}

	reportError(nil)
	gatewayNames := referencesToInternalNames(parentRefs)
	if len(gatewayNames) == 0 {
		// TODO we need to properly return not admitted here
		return nil
	}
	vsConfig := config.Config{
		Meta: config.Meta{
			CreationTimestamp: obj.CreationTimestamp,
			GroupVersionKind:  gvk.VirtualService,
			Name:              fmt.Sprintf("%s-tls-%s", obj.Name, constants.KubernetesGatewayName),
			Annotations:       parentMeta(obj, nil),
			Namespace:         obj.Namespace,
			Domain:            domain,
		},
		Spec: &istio.VirtualService{
			Hosts:    hostnamesToStringListWithWildcard(route.Hostnames),
			Gateways: gatewayNames,
			Tls:      routes,
		},
	}
	return &vsConfig
}

func buildTCPDestination(forwardTo []gatewayapiV1beta1.BackendRef, ns, domain string) ([]*istio.RouteDestination, *ConfigError) {
	if forwardTo == nil {
		return nil, nil
	}

	weights := []int{}
	action := []gatewayapiV1beta1.BackendRef{}
	for i, w := range forwardTo {
		wt := 1
		if w.Weight != nil {
			wt = int(*w.Weight)
		}
		if wt == 0 {
			continue
		}
		action = append(action, forwardTo[i])
		weights = append(weights, wt)
	}
	weights = standardizeWeights(weights)
	res := []*istio.RouteDestination{}
	for i, fwd := range action {
		dst, err := buildDestination(fwd, ns, domain)
		if err != nil {
			return nil, err
		}
		res = append(res, &istio.RouteDestination{
			Destination: dst,
			Weight:      int32(weights[i]),
		})
	}
	return res, nil
}

func buildTLSMatch(hostnames []gatewayapiV1alpha2.Hostname) []*istio.TLSMatchAttributes {
	// Currently, the spec only supports extensions beyond hostname, which are not currently implemented by Istio.
	return []*istio.TLSMatchAttributes{{
		SniHosts: hostnamesToStringListWithWildcard(hostnames),
	}}
}

func hostnamesToStringListWithWildcard(h []gatewayapiV1alpha2.Hostname) []string {
	if len(h) == 0 {
		return []string{"*"}
	}
	res := make([]string, 0, len(h))
	for _, i := range h {
		res = append(res, string(i))
	}
	return res
}

func intSum(n []int) int {
	r := 0
	for _, i := range n {
		r += i
	}
	return r
}

func buildHTTPDestination(forwardTo []gatewayapiV1beta1.HTTPBackendRef, ns string, domain string, totalZero bool, fallbackCluster string) ([]*istio.HTTPRouteDestination, *ConfigError) {
	if forwardTo == nil {
		return nil, nil
	}

	var fallbackDest *istio.Destination
	if fallbackCluster != "" {
		var port uint64
		host := fallbackCluster
		colon := strings.LastIndex(fallbackCluster, ":")
		if colon != -1 {
			var err error
			port, err = strconv.ParseUint(fallbackCluster[colon+1:], 10, 32)
			if err == nil && port > 0 && port < 65536 {
				host = fallbackCluster[:colon]
			}
		}
		fallbackDest = &istio.Destination{
			Host: host,
		}
		if port > 0 {
			fallbackDest.Port = &istio.PortSelector{
				Number: uint32(port),
			}
		}
	}

	weights := []int{}
	action := []gatewayapiV1beta1.HTTPBackendRef{}
	for i, w := range forwardTo {
		wt := 1
		if w.Weight != nil {
			wt = int(*w.Weight)
		}
		// When total weight is zero, create destination to add falutInjection.
		// When total weight is not zero, do not create the destination.
		if wt == 0 && !totalZero {
			continue
		}
		action = append(action, forwardTo[i])
		weights = append(weights, wt)
	}
	weights = standardizeWeights(weights)
	res := []*istio.HTTPRouteDestination{}
	for i, fwd := range action {
		dst, err := buildDestination(fwd.BackendRef, ns, domain)
		if err != nil {
			return nil, err
		}
		rd := &istio.HTTPRouteDestination{
			Destination: dst,
			Weight:      int32(weights[i]),
		}
		if fallbackDest != nil {
			rd.FallbackClusters = append(rd.FallbackClusters, fallbackDest)
		}
		for _, filter := range fwd.Filters {
			switch filter.Type {
			case gatewayapiV1beta1.HTTPRouteFilterRequestHeaderModifier:
				rd.Headers = createHeadersFilter(filter.RequestHeaderModifier)
			default:
				return nil, &ConfigError{Reason: InvalidFilter, Message: fmt.Sprintf("unsupported filter type %q", filter.Type)}
			}
		}
		res = append(res, rd)
	}
	return res, nil
}

func buildDestination(to gatewayapiV1beta1.BackendRef, ns, domain string) (*istio.Destination, *ConfigError) {
	namespace := defaultIfNil((*string)(to.Namespace), ns)
	if nilOrEqual((*string)(to.Group), "") && nilOrEqual((*string)(to.Kind), gvk.Service.Kind) {
		// Service
		if to.Port == nil {
			// "Port is required when the referent is a Kubernetes Service."
			return nil, &ConfigError{Reason: InvalidDestination, Message: "port is required in backendRef"}
		}
		if strings.Contains(string(to.Name), ".") {
			return nil, &ConfigError{Reason: InvalidDestination, Message: "serviceName invalid; the name of the Service must be used, not the hostname."}
		}
		return &istio.Destination{
			// TODO: implement ReferenceGrant for cross namespace
			Host: fmt.Sprintf("%s.%s.svc.%s", to.Name, namespace, domain),
			Port: &istio.PortSelector{Number: uint32(*to.Port)},
		}, nil
	}
	if nilOrEqual((*string)(to.Group), gvk.ServiceEntry.Group) && nilOrEqual((*string)(to.Kind), "Hostname") {
		// Hostname synthetic type
		if to.Port == nil {
			// We don't know where to send without port
			return nil, &ConfigError{Reason: InvalidDestination, Message: "port is required in backendRef"}
		}
		if to.Namespace != nil {
			return nil, &ConfigError{Reason: InvalidDestination, Message: "namespace may not be set with Hostname type"}
		}
		return &istio.Destination{
			Host: string(to.Name),
			Port: &istio.PortSelector{Number: uint32(*to.Port)},
		}, nil
	}
	if equal((*string)(to.Group), "networking.higress.io") && nilOrEqual((*string)(to.Kind), "Service") {
		var port *istio.PortSelector
		if to.Port != nil {
			port = &istio.PortSelector{Number: uint32(*to.Port)}
		}
		return &istio.Destination{
			Host: string(to.Name),
			Port: port,
		}, nil
	}
	return nil, &ConfigError{
		Reason:  InvalidDestination,
		Message: fmt.Sprintf("referencing unsupported backendRef: group %q kind %q", emptyIfNil((*string)(to.Group)), emptyIfNil((*string)(to.Kind))),
	}
}

// standardizeWeights migrates a list of weights from relative weights, to weights out of 100
// In the event we cannot cleanly move to 100 denominator, we will round up weights in order. See test for details.
// TODO in the future we should probably just make VirtualService support relative weights directly
func standardizeWeights(weights []int) []int {
	if len(weights) == 1 {
		// Instead of setting weight=100 for a single destination, we will not set weight at all
		return []int{0}
	}
	total := intSum(weights)
	if total == 0 {
		// All empty, fallback to even weight
		for i := range weights {
			weights[i] = 1
		}
		total = len(weights)
	}
	results := make([]int, 0, len(weights))
	remainders := make([]float64, 0, len(weights))
	for _, w := range weights {
		perc := float64(w) / float64(total)
		rounded := int(perc * 100)
		remainders = append(remainders, (perc*100)-float64(rounded))
		results = append(results, rounded)
	}
	remaining := 100 - intSum(results)
	order := argsort(remainders)
	for _, idx := range order {
		if remaining <= 0 {
			break
		}
		remaining--
		results[idx]++
	}
	return results
}

type argSlice struct {
	sort.Interface
	idx []int
}

func (s argSlice) Swap(i, j int) {
	s.Interface.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func argsort(n []float64) []int {
	s := &argSlice{Interface: sort.Float64Slice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	sort.Sort(sort.Reverse(s))
	return s.idx
}

func headerListToMap(hl []gatewayapiV1beta1.HTTPHeader) map[string]string {
	if len(hl) == 0 {
		return nil
	}
	res := map[string]string{}
	for _, e := range hl {
		k := strings.ToLower(string(e.Name))
		if _, f := res[k]; f {
			// "Subsequent entries with an equivalent header name MUST be ignored"
			continue
		}
		res[k] = e.Value
	}
	return res
}

func createMirrorFilter(filter *gatewayapiV1beta1.HTTPRequestMirrorFilter, ns, domain string) (*istio.Destination, *ConfigError) {
	if filter == nil {
		return nil, nil
	}
	var weightOne int32 = 1
	return buildDestination(gatewayapiV1beta1.BackendRef{
		BackendObjectReference: filter.BackendRef,
		Weight:                 &weightOne,
	}, ns, domain)
}

func createRedirectFilter(filter *gatewayapiV1beta1.HTTPRequestRedirectFilter) *istio.HTTPRedirect {
	if filter == nil {
		return nil
	}
	resp := &istio.HTTPRedirect{}
	if filter.StatusCode != nil {
		// Istio allows 301, 302, 303, 307, 308.
		// Gateway allows only 301 and 302.
		resp.RedirectCode = uint32(*filter.StatusCode)
	}
	if filter.Hostname != nil {
		resp.Authority = string(*filter.Hostname)
	}
	if filter.Scheme != nil {
		// Both allow http and https
		resp.Scheme = *filter.Scheme
	}
	if filter.Port != nil {
		resp.RedirectPort = &istio.HTTPRedirect_Port{Port: uint32(*filter.Port)}
	} else {
		// "When empty, port (if specified) of the request is used."
		// this differs from Istio default
		resp.RedirectPort = &istio.HTTPRedirect_DerivePort{DerivePort: istio.HTTPRedirect_FROM_REQUEST_PORT}
	}
	return resp
}

func createHeadersFilter(filter *gatewayapiV1beta1.HTTPRequestHeaderFilter) *istio.Headers {
	if filter == nil {
		return nil
	}
	return &istio.Headers{
		Request: &istio.Headers_HeaderOperations{
			Add:    headerListToMap(filter.Add),
			Remove: filter.Remove,
			Set:    headerListToMap(filter.Set),
		},
	}
}

// nolint: unparam
func createMethodMatch(match gatewayapiV1beta1.HTTPRouteMatch) (*istio.StringMatch, *ConfigError) {
	if match.Method == nil {
		return nil, nil
	}
	return &istio.StringMatch{
		MatchType: &istio.StringMatch_Exact{Exact: string(*match.Method)},
	}, nil
}

func createQueryParamsMatch(match gatewayapiV1beta1.HTTPRouteMatch) (map[string]*istio.StringMatch, *ConfigError) {
	res := map[string]*istio.StringMatch{}
	for _, qp := range match.QueryParams {
		tp := gatewayapiV1beta1.QueryParamMatchExact
		if qp.Type != nil {
			tp = *qp.Type
		}
		switch tp {
		case gatewayapiV1beta1.QueryParamMatchExact:
			res[string(qp.Name)] = &istio.StringMatch{
				MatchType: &istio.StringMatch_Exact{Exact: qp.Value},
			}
		case gatewayapiV1beta1.QueryParamMatchRegularExpression:
			res[string(qp.Name)] = &istio.StringMatch{
				MatchType: &istio.StringMatch_Regex{Regex: qp.Value},
			}
		default:
			// Should never happen, unless a new field is added
			return nil, &ConfigError{Reason: InvalidConfiguration, Message: fmt.Sprintf("unknown type: %q is not supported QueryParams type", tp)}
		}
	}

	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

func createHeadersMatch(match gatewayapiV1beta1.HTTPRouteMatch) (map[string]*istio.StringMatch, *ConfigError) {
	res := map[string]*istio.StringMatch{}
	for _, header := range match.Headers {
		tp := gatewayapiV1beta1.HeaderMatchExact
		if header.Type != nil {
			tp = *header.Type
		}
		switch tp {
		case gatewayapiV1beta1.HeaderMatchExact:
			res[string(header.Name)] = &istio.StringMatch{
				MatchType: &istio.StringMatch_Exact{Exact: header.Value},
			}
		case gatewayapiV1beta1.HeaderMatchRegularExpression:
			res[string(header.Name)] = &istio.StringMatch{
				MatchType: &istio.StringMatch_Regex{Regex: header.Value},
			}
		default:
			// Should never happen, unless a new field is added
			return nil, &ConfigError{Reason: InvalidConfiguration, Message: fmt.Sprintf("unknown type: %q is not supported HeaderMatch type", tp)}
		}
	}

	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

// prefixMatchRegex optionally matches "/..." at the end of a path.
// regex taken from https://github.com/projectcontour/contour/blob/2b3376449bedfea7b8cea5fbade99fb64009c0f6/internal/envoy/v3/route.go#L59
const prefixMatchRegex = `((\/).*)?`

func createURIMatch(match gatewayapiV1beta1.HTTPRouteMatch) (*istio.StringMatch, *ConfigError) {
	tp := gatewayapiV1beta1.PathMatchPathPrefix
	if match.Path.Type != nil {
		tp = *match.Path.Type
	}
	dest := "/"
	if match.Path.Value != nil {
		dest = *match.Path.Value
	}
	switch tp {
	case gatewayapiV1beta1.PathMatchPathPrefix:
		path := *match.Path.Value
		if path == "/" {
			// Optimize common case of / to not needed regex
			return &istio.StringMatch{
				MatchType: &istio.StringMatch_Prefix{Prefix: path},
			}, nil
		}
		path = strings.TrimSuffix(path, "/")
		return &istio.StringMatch{
			MatchType: &istio.StringMatch_Regex{Regex: regexp.QuoteMeta(path) + prefixMatchRegex},
		}, nil
	case gatewayapiV1beta1.PathMatchExact:
		return &istio.StringMatch{
			MatchType: &istio.StringMatch_Exact{Exact: dest},
		}, nil
	case gatewayapiV1beta1.PathMatchRegularExpression:
		return &istio.StringMatch{
			MatchType: &istio.StringMatch_Regex{Regex: dest},
		}, nil
	default:
		// Should never happen, unless a new field is added
		return nil, &ConfigError{Reason: InvalidConfiguration, Message: fmt.Sprintf("unknown type: %q is not supported Path match type", tp)}
	}
}

// getGatewayClass finds all gateway class that are owned by Istio
func getGatewayClasses(r *KubernetesResources) map[string]struct{} {
	classes := map[string]struct{}{}
	builtinClassExists := false
	for _, obj := range r.GatewayClass {
		gwc := obj.Spec.(*gatewayapiV1beta1.GatewayClassSpec)
		if obj.Name == DefaultClassName {
			builtinClassExists = true
		}
		log.Errorf("HigressGatewayClass: %s", features.HigressGatewayClass)
		if string(gwc.ControllerName) == features.HigressGatewayClass {
			// TODO we can add any settings we need here needed for the controller
			// For now, we have none, so just add a struct
			log.Errorf("HigressGatewayClass equal")
			classes[obj.Name] = struct{}{}

			obj.Status.(*kstatus.WrappedStatus).Mutate(func(s config.Status) config.Status {
				gcs := s.(*gatewayapiV1beta1.GatewayClassStatus)
				gcs.Conditions = kstatus.UpdateConditionIfChanged(gcs.Conditions, metav1.Condition{
					Type:               string(gatewayapiV1beta1.GatewayClassConditionStatusAccepted),
					Status:             kstatus.StatusTrue,
					ObservedGeneration: obj.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatewayapiV1beta1.GatewayClassConditionStatusAccepted),
					Message:            "Handled by Higress controller",
				})
				return gcs
			})
		}
	}
	if !builtinClassExists {
		// Allow `istio` class without explicit GatewayClass. However, if it already exists then do not
		// add it here, in case it points to a different controller.
		classes[DefaultClassName] = struct{}{}
	}
	return classes
}

var meshGVK = config.GroupVersionKind{
	Group:   gvk.KubernetesGateway.Group,
	Version: gvk.KubernetesGateway.Version,
	Kind:    "Mesh",
}

// parentKey holds info about a parentRef (ie route binding to a Gateway). This is a mirror of
// gatewayapiV1beta1.ParentRef in a form that can be stored in a map
type parentKey struct {
	Kind config.GroupVersionKind
	// Name is the original name of the resource (ie Kubernetes Gateway name)
	Name string
	// Namespace is the namespace of the resource
	Namespace string
}

// parentInfo holds info about a "parent" - something that can be referenced as a ParentRef in the API.
// Today, this is just Gateway and Mesh.
type parentInfo struct {
	// InternalName refers to the internal name we can reference it by. For example, "mesh" or "my-ns/my-gateway"
	InternalName string
	// AllowedKinds indicates which kinds can be admitted by this parent
	AllowedKinds []gatewayapiV1beta1.RouteGroupKind
	// Hostnames is the hostnames that must be match to reference to the parent. For gateway this is listener hostname
	// Format is ns/hostname
	Hostnames []string
	// OriginalHostname is the unprocessed form of Hostnames; how it appeared in users' config
	OriginalHostname string

	// AttachedRoutes keeps track of how many routes are attached to this parent. This is tracked for status.
	// Because this is mutate in the route generation, parentInfo must be passed as a pointer
	AttachedRoutes int32
	// ReportAttachedRoutes is a callback that should be triggered once all AttachedRoutes are computed, to
	// actually store the attached route count in the status
	ReportAttachedRoutes func()
}

// routeParentReference holds information about a route's parent reference
type routeParentReference struct {
	// InternalName refers to the internal name of the parent we can reference it by. For example, "mesh" or "my-ns/my-gateway"
	InternalName string
	// DeniedReason, if present, indicates why the reference was not valid
	DeniedReason error
	// OriginalReference contains the original reference
	OriginalReference gatewayapiV1beta1.ParentReference
}

// referencesToInternalNames converts valid parent references to names that can be used in VirtualService
func referencesToInternalNames(parents []routeParentReference) []string {
	ret := make([]string, 0, len(parents))
	for _, p := range parents {
		if p.DeniedReason != nil {
			// We should filter this out
			continue
		}
		ret = append(ret, p.InternalName)
	}
	// To ensure deterministic order, sort them
	sort.Strings(ret)
	return ret
}

func convertGateways(r *KubernetesResources) ([]config.Config, map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo, sets.Set) {
	// result stores our generated Istio Gateways
	result := []config.Config{}
	// gwMap stores an index to access parentInfo (which corresponds to a Kubernetes Gateway)
	gwMap := map[parentKey]map[gatewayapiV1beta1.SectionName]*parentInfo{}
	// namespaceLabelReferences keeps track of all namespace label keys referenced by Gateways. This is
	// used to ensure we handle namespace updates for those keys.
	namespaceLabelReferences := sets.NewSet()
	classes := getGatewayClasses(r)
	for _, obj := range r.Gateway {
		obj := obj
		kgw := obj.Spec.(*gatewayapiV1beta1.GatewaySpec)
		if _, f := classes[string(kgw.GatewayClassName)]; !f {
			// No gateway class found, this may be meant for another controller; should be skipped.
			continue
		}

		// Setup initial conditions to the success state. If we encounter errors, we will update this.
		gatewayConditions := map[string]*condition{
			string(gatewayapiV1beta1.GatewayConditionReady): {
				reason:  "ListenersValid",
				message: "Listeners valid",
			},
		}
		if isManaged(kgw) {
			gatewayConditions[string(gatewayapiV1beta1.GatewayConditionScheduled)] = &condition{
				error: &ConfigError{
					Reason:  "ResourcesPending",
					Message: "Resources not yet deployed to the cluster",
				},
				setOnce: true,
			}
		} else {
			gatewayConditions[string(gatewayapiV1beta1.GatewayConditionScheduled)] = &condition{
				reason:  "ResourcesAvailable",
				message: "Resources available",
			}
		}
		servers := []*istio.Server{}

		// Extract the addresses. A gateway will bind to a specific Service
		gatewayServices, skippedAddresses := extractGatewayServices(r, kgw, obj)
		invalidListeners := []gatewayapiV1beta1.SectionName{}
		for i, l := range kgw.Listeners {
			i := i
			namespaceLabelReferences.Insert(getNamespaceLabelReferences(l.AllowedRoutes)...)
			server, ok := buildListener(r, obj, l, i)
			if !ok {
				invalidListeners = append(invalidListeners, l.Name)
				continue
			}
			meta := parentMeta(obj, &l.Name)
			meta[model.InternalGatewayServiceAnnotation] = strings.Join(gatewayServices, ",")
			// Each listener generates an Istio Gateway with a single Server. This allows binding to a specific listener.
			gatewayConfig := config.Config{
				Meta: config.Meta{
					CreationTimestamp: obj.CreationTimestamp,
					GroupVersionKind:  gvk.Gateway,
					Name:              fmt.Sprintf("%s-%s-%s", obj.Name, constants.KubernetesGatewayName, l.Name),
					Annotations:       meta,
					Namespace:         obj.Namespace,
					Domain:            r.Domain,
				},
				Spec: &istio.Gateway{
					Servers: []*istio.Server{server},
				},
			}
			ref := parentKey{
				Kind:      gvk.KubernetesGateway,
				Name:      obj.Name,
				Namespace: obj.Namespace,
			}
			if _, f := gwMap[ref]; !f {
				gwMap[ref] = map[gatewayapiV1beta1.SectionName]*parentInfo{}
			}

			pri := &parentInfo{
				InternalName:     obj.Namespace + "/" + gatewayConfig.Name,
				AllowedKinds:     generateSupportedKinds(l),
				Hostnames:        server.Hosts,
				OriginalHostname: emptyIfNil((*string)(l.Hostname)),
			}
			pri.ReportAttachedRoutes = func() {
				reportListenerAttachedRoutes(i, obj, pri.AttachedRoutes)
			}
			gwMap[ref][l.Name] = pri
			result = append(result, gatewayConfig)
			servers = append(servers, server)
		}

		internal, external, warnings := r.Context.ResolveGatewayInstances(obj.Namespace, gatewayServices, servers)
		if len(skippedAddresses) > 0 {
			warnings = append(warnings, fmt.Sprintf("Only Hostname is supported, ignoring %v", skippedAddresses))
		}
		if len(warnings) > 0 {
			var msg string
			if len(internal) > 0 {
				msg = fmt.Sprintf("Assigned to service(s) %s, but failed to assign to all requested addresses: %s",
					humanReadableJoin(internal), strings.Join(warnings, "; "))
			} else {
				msg = fmt.Sprintf("failed to assign to any requested addresses: %s", strings.Join(warnings, "; "))
			}
			gatewayConditions[string(gatewayapiV1beta1.GatewayConditionReady)].error = &ConfigError{
				Reason:  string(gatewayapiV1beta1.GatewayReasonAddressNotAssigned),
				Message: msg,
			}
		} else if len(invalidListeners) > 0 {
			gatewayConditions[string(gatewayapiV1beta1.GatewayConditionReady)].error = &ConfigError{
				Reason:  string(gatewayapiV1beta1.GatewayReasonListenersNotValid),
				Message: fmt.Sprintf("Invalid listeners: %v", invalidListeners),
			}
		} else {
			gatewayConditions[string(gatewayapiV1beta1.GatewayConditionReady)].message = fmt.Sprintf("Gateway valid, assigned to service(s) %s", humanReadableJoin(internal))
		}
		obj.Status.(*kstatus.WrappedStatus).Mutate(func(s config.Status) config.Status {
			gs := s.(*gatewayapiV1beta1.GatewayStatus)
			addressesToReport := external
			addrType := gatewayapiV1beta1.IPAddressType
			if len(addressesToReport) == 0 {
				// There are no external addresses, so report the internal ones
				// TODO: should we always report both?
				addressesToReport = internal
				addrType = gatewayapiV1beta1.HostnameAddressType
			}
			gs.Addresses = make([]gatewayapiV1beta1.GatewayAddress, 0, len(addressesToReport))
			for _, addr := range addressesToReport {
				addrPairs := strings.Split(addr, ":")
				gs.Addresses = append(gs.Addresses, gatewayapiV1beta1.GatewayAddress{
					Type: &addrType,
					// strip the port
					Value: addrPairs[0],
				})
			}
			return gs
		})
		reportGatewayCondition(obj, gatewayConditions)
	}
	// Insert a parent for Mesh references.
	gwMap[parentKey{
		Kind: meshGVK,
		Name: "istio",
	}] = map[gatewayapiV1beta1.SectionName]*parentInfo{
		"": {
			InternalName: "mesh",
		},
	}
	return result, gwMap, namespaceLabelReferences
}

// isManaged checks if a Gateway is managed (ie we create the Deployment and Service) or unmanaged.
// This is based on the address field of the spec. If address is set with a Hostname type, it should point to an existing
// Service that handles the gateway traffic. If it is not set, or refers to only a single IP, we will consider it managed and provision the Service.
// If there is an IP, we will set the `loadBalancerIP` type.
// While there is no defined standard for this in the API yet, it is tracked in https://github.com/kubernetes-sigs/gateway-api/issues/892.
// So far, this mirrors how out of clusters work (address set means to use existing IP, unset means to provision one),
// and there has been growing consensus on this model for in cluster deployments.
func isManaged(gw *gatewayapiV1beta1.GatewaySpec) bool {
	if len(gw.Addresses) == 0 {
		return true
	}
	if len(gw.Addresses) > 1 {
		return false
	}
	if t := gw.Addresses[0].Type; t == nil || *t == gatewayapiV1beta1.IPAddressType {
		return true
	}
	return false
}

func extractGatewayServices(r *KubernetesResources, kgw *gatewayapiV1beta1.GatewaySpec, obj config.Config) ([]string, []string) {
	if isManaged(kgw) {
		return []string{fmt.Sprintf("%s.%s.svc.%v", obj.Name, obj.Namespace, r.Domain)}, nil
	}
	gatewayServices := []string{}
	skippedAddresses := []string{}
	for _, addr := range kgw.Addresses {
		if addr.Type != nil && *addr.Type != gatewayapiV1beta1.HostnameAddressType {
			// We only support HostnameAddressType. Keep track of invalid ones so we can report in status.
			skippedAddresses = append(skippedAddresses, addr.Value)
			continue
		}
		// TODO: For now we are using Addresses. There has been some discussion of allowing inline
		// parameters on the class field like a URL, in which case we will probably just use that. See
		// https://github.com/kubernetes-sigs/gateway-api/pull/614
		fqdn := addr.Value
		if !strings.Contains(fqdn, ".") {
			// Short name, expand it
			fqdn = fmt.Sprintf("%s.%s.svc.%s", fqdn, obj.Namespace, r.Domain)
		}
		gatewayServices = append(gatewayServices, fqdn)
	}
	return gatewayServices, skippedAddresses
}

// getNamespaceLabelReferences fetches all label keys used in namespace selectors. Return order may not be stable.
func getNamespaceLabelReferences(routes *gatewayapiV1beta1.AllowedRoutes) []string {
	if routes == nil || routes.Namespaces == nil || routes.Namespaces.Selector == nil {
		return nil
	}
	res := []string{}
	for k := range routes.Namespaces.Selector.MatchLabels {
		res = append(res, k)
	}
	for _, me := range routes.Namespaces.Selector.MatchExpressions {
		res = append(res, me.Key)
	}
	return res
}

func buildListener(r *KubernetesResources, obj config.Config, l gatewayapiV1beta1.Listener, listenerIndex int) (*istio.Server, bool) {
	listenerConditions := map[string]*condition{
		string(gatewayapiV1beta1.ListenerConditionReady): {
			reason:  "ListenerReady",
			message: "No errors found",
		},
		string(gatewayapiV1beta1.ListenerConditionDetached): {
			reason:  "ListenerReady",
			message: "No errors found",
			status:  kstatus.StatusFalse,
		},
		string(gatewayapiV1beta1.ListenerConditionConflicted): {
			reason:  "ListenerReady",
			message: "No errors found",
			status:  kstatus.StatusFalse,
		},
		string(gatewayapiV1beta1.ListenerConditionResolvedRefs): {
			reason:  "ListenerReady",
			message: "No errors found",
		},
	}
	defer reportListenerCondition(listenerIndex, l, obj, listenerConditions)
	tls, err := buildTLS(l.TLS, obj.Namespace)
	if err != nil {
		listenerConditions[string(gatewayapiV1beta1.ListenerConditionReady)].error = &ConfigError{
			Reason:  string(gatewayapiV1beta1.ListenerReasonInvalid),
			Message: err.Message,
		}
		listenerConditions[string(gatewayapiV1beta1.ListenerConditionResolvedRefs)].error = &ConfigError{
			Reason:  string(gatewayapiV1beta1.ListenerReasonInvalidCertificateRef),
			Message: err.Message,
		}
		return nil, false
	}
	hostnames := buildHostnameMatch(obj.Namespace, r, l)
	server := &istio.Server{
		Port: &istio.Port{
			// Name is required. We only have one server per Gateway, so we can just name them all the same
			Name:     "default",
			Number:   uint32(l.Port),
			Protocol: listenerProtocolToIstio(l.Protocol),
		},
		Hosts: hostnames,
		Tls:   tls,
	}

	return server, true
}

func listenerProtocolToIstio(protocol gatewayapiV1beta1.ProtocolType) string {
	// Currently, all gateway-api protocols are valid Istio protocols.
	return string(protocol)
}

func buildTLS(tls *gatewayapiV1beta1.GatewayTLSConfig, namespace string) (*istio.ServerTLSSettings, *ConfigError) {
	if tls == nil {
		return nil, nil
	}
	// Explicitly not supported: file mounted
	// Not yet implemented: TLS mode, https redirect, max protocol version, SANs, CipherSuites, VerifyCertificate

	out := &istio.ServerTLSSettings{
		HttpsRedirect: false,
	}
	mode := gatewayapiV1beta1.TLSModeTerminate
	if tls.Mode != nil {
		mode = *tls.Mode
	}
	switch mode {
	case gatewayapiV1beta1.TLSModeTerminate:
		out.Mode = istio.ServerTLSSettings_SIMPLE
		if len(tls.CertificateRefs) != 1 {
			// This is required in the API, should be rejected in validation
			return nil, &ConfigError{Reason: InvalidConfiguration, Message: "exactly 1 certificateRefs should be present for TLS termination"}
		}
		cred, err := buildSecretReference(tls.CertificateRefs[0], namespace)
		if err != nil {
			return nil, err
		}
		out.CredentialName = cred
	case gatewayapiV1beta1.TLSModePassthrough:
		out.Mode = istio.ServerTLSSettings_PASSTHROUGH
	}
	return out, nil
}

func buildSecretReference(ref gatewayapiV1beta1.SecretObjectReference, defaultNamespace string) (string, *ConfigError) {
	if !nilOrEqual((*string)(ref.Group), gvk.Secret.Group) || !nilOrEqual((*string)(ref.Kind), gvk.Secret.Kind) {
		return "", &ConfigError{Reason: InvalidTLS, Message: fmt.Sprintf("invalid certificate reference %v, only secret is allowed", objectReferenceString(ref))}
	}
	return credentials.ToKubernetesGatewayResource(defaultIfNil((*string)(ref.Namespace), defaultNamespace), string(ref.Name)), nil
}

func objectReferenceString(ref gatewayapiV1beta1.SecretObjectReference) string {
	return fmt.Sprintf("%s/%s/%s.%s",
		emptyIfNil((*string)(ref.Group)),
		emptyIfNil((*string)(ref.Kind)),
		ref.Name,
		emptyIfNil((*string)(ref.Namespace)))
}

func parentRefString(ref gatewayapiV1beta1.ParentReference) string {
	return fmt.Sprintf("%s/%s/%s/%s.%s",
		emptyIfNil((*string)(ref.Group)),
		emptyIfNil((*string)(ref.Kind)),
		ref.Name,
		emptyIfNil((*string)(ref.SectionName)),
		emptyIfNil((*string)(ref.Namespace)))
}

// buildHostnameMatch generates a VirtualService.spec.hosts section from a listener
func buildHostnameMatch(localNamespace string, r *KubernetesResources, l gatewayapiV1beta1.Listener) []string {
	// We may allow all hostnames or a specific one
	hostname := "*"
	if l.Hostname != nil {
		hostname = string(*l.Hostname)
	}

	resp := []string{}
	for _, ns := range namespacesFromSelector(localNamespace, r, l.AllowedRoutes) {
		resp = append(resp, fmt.Sprintf("%s/%s", ns, hostname))
	}

	// If nothing matched use ~ namespace (match nothing). We need this since its illegal to have an
	// empty hostname list, but we still need the Gateway provisioned to ensure status is properly set and
	// SNI matches are established; we just don't want to actually match any routing rules (yet).
	if len(resp) == 0 {
		return []string{"~/" + hostname}
	}
	return resp
}

// namespacesFromSelector determines a list of allowed namespaces for a given AllowedRoutes
func namespacesFromSelector(localNamespace string, r *KubernetesResources, lr *gatewayapiV1beta1.AllowedRoutes) []string {
	// Default is to allow only the same namespace
	if lr == nil || lr.Namespaces == nil || lr.Namespaces.From == nil || *lr.Namespaces.From == gatewayapiV1beta1.NamespacesFromSame {
		return []string{localNamespace}
	}
	if *lr.Namespaces.From == gatewayapiV1beta1.NamespacesFromAll {
		return []string{"*"}
	}

	if lr.Namespaces.Selector == nil {
		// Should never happen, invalid config
		return []string{"*"}
	}

	// gateway-api has selectors, but Istio Gateway just has a list of names. We will run the selector
	// against all namespaces and get a list of matching namespaces that can be converted into a list
	// Istio can handle.
	ls, err := metav1.LabelSelectorAsSelector(lr.Namespaces.Selector)
	if err != nil {
		return nil
	}
	namespaces := []string{}
	for _, ns := range r.Namespaces {
		if ls.Matches(toNamespaceSet(ns.Name, ns.Labels)) {
			namespaces = append(namespaces, ns.Name)
		}
	}
	// Ensure stable order
	sort.Strings(namespaces)
	return namespaces
}

func emptyIfNil(s *string) string {
	return defaultIfNil(s, "")
}

func defaultIfNil(s *string, d string) string {
	if s != nil {
		return *s
	}
	return d
}

func equal(have *string, expected string) bool {
	return have != nil && *have == expected
}

func nilOrEqual(have *string, expected string) bool {
	return have == nil || *have == expected
}

func StrPointer(s string) *string {
	return &s
}

func humanReadableJoin(ss []string) string {
	switch len(ss) {
	case 0:
		return ""
	case 1:
		return ss[0]
	case 2:
		return ss[0] + " and " + ss[1]
	default:
		return strings.Join(ss[:len(ss)-1], ", ") + ", and " + ss[len(ss)-1]
	}
}

// NamespaceNameLabel represents that label added automatically to namespaces is newer Kubernetes clusters
const NamespaceNameLabel = "kubernetes.io/metadata.name"

// toNamespaceSet converts a set of namespace labels to a Set that can be used to select against.
func toNamespaceSet(name string, labels map[string]string) klabels.Set {
	// If namespace label is not set, implicitly insert it to support older Kubernetes versions
	if labels[NamespaceNameLabel] == name {
		// Already set, avoid copies
		return labels
	}
	// First we need a copy to not modify the underlying object
	ret := make(map[string]string, len(labels)+1)
	for k, v := range labels {
		ret[k] = v
	}
	ret[NamespaceNameLabel] = name
	return ret
}
