// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1beta1/virtual_service.proto

// Configuration affecting traffic routing. Here are a few terms useful to define
// in the context of traffic routing.
//
// `Service` a unit of application behavior bound to a unique name in a
// service registry. Services consist of multiple network *endpoints*
// implemented by workload instances running on pods, containers, VMs etc.
//
// `Service versions (a.k.a. subsets)` - In a continuous deployment
// scenario, for a given service, there can be distinct subsets of
// instances running different variants of the application binary. These
// variants are not necessarily different API versions. They could be
// iterative changes to the same service, deployed in different
// environments (prod, staging, dev, etc.). Common scenarios where this
// occurs include A/B testing, canary rollouts, etc. The choice of a
// particular version can be decided based on various criterion (headers,
// url, etc.) and/or by weights assigned to each version. Each service has
// a default version consisting of all its instances.
//
// `Source` - A downstream client calling a service.
//
// `Host` - The address used by a client when attempting to connect to a
// service.
//
// `Access model` - Applications address only the destination service
// (Host) without knowledge of individual service versions (subsets). The
// actual choice of the version is determined by the proxy/sidecar, enabling the
// application code to decouple itself from the evolution of dependent
// services.
//
// A `VirtualService` defines a set of traffic routing rules to apply when a host is
// addressed. Each routing rule defines matching criteria for traffic of a specific
// protocol. If the traffic is matched, then it is sent to a named destination service
// (or subset/version of it) defined in the registry.
//
// The source of traffic can also be matched in a routing rule. This allows routing
// to be customized for specific client contexts.
//
// The following example on Kubernetes, routes all HTTP traffic by default to
// pods of the reviews service with label "version: v1". In addition,
// HTTP requests with path starting with /wpcatalog/ or /consumercatalog/ will
// be rewritten to /newcatalog and sent to pods with label "version: v2".
//
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: reviews-route
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   http:
//   - name: "reviews-v2-routes"
//     match:
//     - uri:
//         prefix: "/wpcatalog"
//     - uri:
//         prefix: "/consumercatalog"
//     rewrite:
//       uri: "/newcatalog"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v2
//   - name: "reviews-v1-route"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v1
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: reviews-route
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   http:
//   - name: "reviews-v2-routes"
//     match:
//     - uri:
//         prefix: "/wpcatalog"
//     - uri:
//         prefix: "/consumercatalog"
//     rewrite:
//       uri: "/newcatalog"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v2
//   - name: "reviews-v1-route"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v1
// ```
// {{</tab>}}
// {{</tabset>}}
//
// A subset/version of a route destination is identified with a reference
// to a named service subset which must be declared in a corresponding
// `DestinationRule`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: reviews-destination
// spec:
//   host: reviews.prod.svc.cluster.local
//   subsets:
//   - name: v1
//     labels:
//       version: v1
//   - name: v2
//     labels:
//       version: v2
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: reviews-destination
// spec:
//   host: reviews.prod.svc.cluster.local
//   subsets:
//   - name: v1
//     labels:
//       version: v1
//   - name: v2
//     labels:
//       version: v2
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using VirtualService within kubernetes types, where deepcopy-gen is used.
func (in *VirtualService) DeepCopyInto(out *VirtualService) {
	p := proto.Clone(in).(*VirtualService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService. Required by controller-gen.
func (in *VirtualService) DeepCopy() *VirtualService {
	if in == nil {
		return nil
	}
	out := new(VirtualService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService. Required by controller-gen.
func (in *VirtualService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Destination within kubernetes types, where deepcopy-gen is used.
func (in *Destination) DeepCopyInto(out *Destination) {
	p := proto.Clone(in).(*Destination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination. Required by controller-gen.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Destination. Required by controller-gen.
func (in *Destination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRoute within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRoute) DeepCopyInto(out *HTTPRoute) {
	p := proto.Clone(in).(*HTTPRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute. Required by controller-gen.
func (in *HTTPRoute) DeepCopy() *HTTPRoute {
	if in == nil {
		return nil
	}
	out := new(HTTPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute. Required by controller-gen.
func (in *HTTPRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Delegate within kubernetes types, where deepcopy-gen is used.
func (in *Delegate) DeepCopyInto(out *Delegate) {
	p := proto.Clone(in).(*Delegate)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delegate. Required by controller-gen.
func (in *Delegate) DeepCopy() *Delegate {
	if in == nil {
		return nil
	}
	out := new(Delegate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Delegate. Required by controller-gen.
func (in *Delegate) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Headers within kubernetes types, where deepcopy-gen is used.
func (in *Headers) DeepCopyInto(out *Headers) {
	p := proto.Clone(in).(*Headers)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Headers. Required by controller-gen.
func (in *Headers) DeepCopy() *Headers {
	if in == nil {
		return nil
	}
	out := new(Headers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Headers. Required by controller-gen.
func (in *Headers) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Headers_HeaderOperations within kubernetes types, where deepcopy-gen is used.
func (in *Headers_HeaderOperations) DeepCopyInto(out *Headers_HeaderOperations) {
	p := proto.Clone(in).(*Headers_HeaderOperations)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Headers_HeaderOperations. Required by controller-gen.
func (in *Headers_HeaderOperations) DeepCopy() *Headers_HeaderOperations {
	if in == nil {
		return nil
	}
	out := new(Headers_HeaderOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Headers_HeaderOperations. Required by controller-gen.
func (in *Headers_HeaderOperations) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TLSRoute within kubernetes types, where deepcopy-gen is used.
func (in *TLSRoute) DeepCopyInto(out *TLSRoute) {
	p := proto.Clone(in).(*TLSRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSRoute. Required by controller-gen.
func (in *TLSRoute) DeepCopy() *TLSRoute {
	if in == nil {
		return nil
	}
	out := new(TLSRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TLSRoute. Required by controller-gen.
func (in *TLSRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TCPRoute within kubernetes types, where deepcopy-gen is used.
func (in *TCPRoute) DeepCopyInto(out *TCPRoute) {
	p := proto.Clone(in).(*TCPRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPRoute. Required by controller-gen.
func (in *TCPRoute) DeepCopy() *TCPRoute {
	if in == nil {
		return nil
	}
	out := new(TCPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TCPRoute. Required by controller-gen.
func (in *TCPRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPMatchRequest within kubernetes types, where deepcopy-gen is used.
func (in *HTTPMatchRequest) DeepCopyInto(out *HTTPMatchRequest) {
	p := proto.Clone(in).(*HTTPMatchRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatchRequest. Required by controller-gen.
func (in *HTTPMatchRequest) DeepCopy() *HTTPMatchRequest {
	if in == nil {
		return nil
	}
	out := new(HTTPMatchRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatchRequest. Required by controller-gen.
func (in *HTTPMatchRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRouteDestination within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRouteDestination) DeepCopyInto(out *HTTPRouteDestination) {
	p := proto.Clone(in).(*HTTPRouteDestination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteDestination. Required by controller-gen.
func (in *HTTPRouteDestination) DeepCopy() *HTTPRouteDestination {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteDestination. Required by controller-gen.
func (in *HTTPRouteDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RouteDestination within kubernetes types, where deepcopy-gen is used.
func (in *RouteDestination) DeepCopyInto(out *RouteDestination) {
	p := proto.Clone(in).(*RouteDestination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteDestination. Required by controller-gen.
func (in *RouteDestination) DeepCopy() *RouteDestination {
	if in == nil {
		return nil
	}
	out := new(RouteDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RouteDestination. Required by controller-gen.
func (in *RouteDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using L4MatchAttributes within kubernetes types, where deepcopy-gen is used.
func (in *L4MatchAttributes) DeepCopyInto(out *L4MatchAttributes) {
	p := proto.Clone(in).(*L4MatchAttributes)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4MatchAttributes. Required by controller-gen.
func (in *L4MatchAttributes) DeepCopy() *L4MatchAttributes {
	if in == nil {
		return nil
	}
	out := new(L4MatchAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new L4MatchAttributes. Required by controller-gen.
func (in *L4MatchAttributes) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TLSMatchAttributes within kubernetes types, where deepcopy-gen is used.
func (in *TLSMatchAttributes) DeepCopyInto(out *TLSMatchAttributes) {
	p := proto.Clone(in).(*TLSMatchAttributes)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSMatchAttributes. Required by controller-gen.
func (in *TLSMatchAttributes) DeepCopy() *TLSMatchAttributes {
	if in == nil {
		return nil
	}
	out := new(TLSMatchAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TLSMatchAttributes. Required by controller-gen.
func (in *TLSMatchAttributes) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRedirect within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRedirect) DeepCopyInto(out *HTTPRedirect) {
	p := proto.Clone(in).(*HTTPRedirect)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirect. Required by controller-gen.
func (in *HTTPRedirect) DeepCopy() *HTTPRedirect {
	if in == nil {
		return nil
	}
	out := new(HTTPRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirect. Required by controller-gen.
func (in *HTTPRedirect) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPInternalActiveRedirect within kubernetes types, where deepcopy-gen is used.
func (in *HTTPInternalActiveRedirect) DeepCopyInto(out *HTTPInternalActiveRedirect) {
	p := proto.Clone(in).(*HTTPInternalActiveRedirect)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPInternalActiveRedirect. Required by controller-gen.
func (in *HTTPInternalActiveRedirect) DeepCopy() *HTTPInternalActiveRedirect {
	if in == nil {
		return nil
	}
	out := new(HTTPInternalActiveRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPInternalActiveRedirect. Required by controller-gen.
func (in *HTTPInternalActiveRedirect) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPInternalActiveRedirect_RedirectPolicy within kubernetes types, where deepcopy-gen is used.
func (in *HTTPInternalActiveRedirect_RedirectPolicy) DeepCopyInto(out *HTTPInternalActiveRedirect_RedirectPolicy) {
	p := proto.Clone(in).(*HTTPInternalActiveRedirect_RedirectPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPInternalActiveRedirect_RedirectPolicy. Required by controller-gen.
func (in *HTTPInternalActiveRedirect_RedirectPolicy) DeepCopy() *HTTPInternalActiveRedirect_RedirectPolicy {
	if in == nil {
		return nil
	}
	out := new(HTTPInternalActiveRedirect_RedirectPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPInternalActiveRedirect_RedirectPolicy. Required by controller-gen.
func (in *HTTPInternalActiveRedirect_RedirectPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPDirectResponse within kubernetes types, where deepcopy-gen is used.
func (in *HTTPDirectResponse) DeepCopyInto(out *HTTPDirectResponse) {
	p := proto.Clone(in).(*HTTPDirectResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPDirectResponse. Required by controller-gen.
func (in *HTTPDirectResponse) DeepCopy() *HTTPDirectResponse {
	if in == nil {
		return nil
	}
	out := new(HTTPDirectResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPDirectResponse. Required by controller-gen.
func (in *HTTPDirectResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRewrite within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRewrite) DeepCopyInto(out *HTTPRewrite) {
	p := proto.Clone(in).(*HTTPRewrite)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRewrite. Required by controller-gen.
func (in *HTTPRewrite) DeepCopy() *HTTPRewrite {
	if in == nil {
		return nil
	}
	out := new(HTTPRewrite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRewrite. Required by controller-gen.
func (in *HTTPRewrite) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RegexMatchAndSubstitute within kubernetes types, where deepcopy-gen is used.
func (in *RegexMatchAndSubstitute) DeepCopyInto(out *RegexMatchAndSubstitute) {
	p := proto.Clone(in).(*RegexMatchAndSubstitute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexMatchAndSubstitute. Required by controller-gen.
func (in *RegexMatchAndSubstitute) DeepCopy() *RegexMatchAndSubstitute {
	if in == nil {
		return nil
	}
	out := new(RegexMatchAndSubstitute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RegexMatchAndSubstitute. Required by controller-gen.
func (in *RegexMatchAndSubstitute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using StringMatch within kubernetes types, where deepcopy-gen is used.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	p := proto.Clone(in).(*StringMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRetry within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRetry) DeepCopyInto(out *HTTPRetry) {
	p := proto.Clone(in).(*HTTPRetry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry. Required by controller-gen.
func (in *HTTPRetry) DeepCopy() *HTTPRetry {
	if in == nil {
		return nil
	}
	out := new(HTTPRetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry. Required by controller-gen.
func (in *HTTPRetry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CorsPolicy within kubernetes types, where deepcopy-gen is used.
func (in *CorsPolicy) DeepCopyInto(out *CorsPolicy) {
	p := proto.Clone(in).(*CorsPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy. Required by controller-gen.
func (in *CorsPolicy) DeepCopy() *CorsPolicy {
	if in == nil {
		return nil
	}
	out := new(CorsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy. Required by controller-gen.
func (in *CorsPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection) DeepCopyInto(out *HTTPFaultInjection) {
	p := proto.Clone(in).(*HTTPFaultInjection)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection. Required by controller-gen.
func (in *HTTPFaultInjection) DeepCopy() *HTTPFaultInjection {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection. Required by controller-gen.
func (in *HTTPFaultInjection) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection_Delay within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection_Delay) DeepCopyInto(out *HTTPFaultInjection_Delay) {
	p := proto.Clone(in).(*HTTPFaultInjection_Delay)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Delay. Required by controller-gen.
func (in *HTTPFaultInjection_Delay) DeepCopy() *HTTPFaultInjection_Delay {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection_Delay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Delay. Required by controller-gen.
func (in *HTTPFaultInjection_Delay) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection_Abort within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection_Abort) DeepCopyInto(out *HTTPFaultInjection_Abort) {
	p := proto.Clone(in).(*HTTPFaultInjection_Abort)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Abort. Required by controller-gen.
func (in *HTTPFaultInjection_Abort) DeepCopy() *HTTPFaultInjection_Abort {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection_Abort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Abort. Required by controller-gen.
func (in *HTTPFaultInjection_Abort) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PortSelector within kubernetes types, where deepcopy-gen is used.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	p := proto.Clone(in).(*PortSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Percent within kubernetes types, where deepcopy-gen is used.
func (in *Percent) DeepCopyInto(out *Percent) {
	p := proto.Clone(in).(*Percent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Percent. Required by controller-gen.
func (in *Percent) DeepCopy() *Percent {
	if in == nil {
		return nil
	}
	out := new(Percent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Percent. Required by controller-gen.
func (in *Percent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFilter within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFilter) DeepCopyInto(out *HTTPFilter) {
	p := proto.Clone(in).(*HTTPFilter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFilter. Required by controller-gen.
func (in *HTTPFilter) DeepCopy() *HTTPFilter {
	if in == nil {
		return nil
	}
	out := new(HTTPFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFilter. Required by controller-gen.
func (in *HTTPFilter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using IPAccessControl within kubernetes types, where deepcopy-gen is used.
func (in *IPAccessControl) DeepCopyInto(out *IPAccessControl) {
	p := proto.Clone(in).(*IPAccessControl)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAccessControl. Required by controller-gen.
func (in *IPAccessControl) DeepCopy() *IPAccessControl {
	if in == nil {
		return nil
	}
	out := new(IPAccessControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new IPAccessControl. Required by controller-gen.
func (in *IPAccessControl) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalRateLimit within kubernetes types, where deepcopy-gen is used.
func (in *LocalRateLimit) DeepCopyInto(out *LocalRateLimit) {
	p := proto.Clone(in).(*LocalRateLimit)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit. Required by controller-gen.
func (in *LocalRateLimit) DeepCopy() *LocalRateLimit {
	if in == nil {
		return nil
	}
	out := new(LocalRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit. Required by controller-gen.
func (in *LocalRateLimit) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TokenBucket within kubernetes types, where deepcopy-gen is used.
func (in *TokenBucket) DeepCopyInto(out *TokenBucket) {
	p := proto.Clone(in).(*TokenBucket)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenBucket. Required by controller-gen.
func (in *TokenBucket) DeepCopy() *TokenBucket {
	if in == nil {
		return nil
	}
	out := new(TokenBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TokenBucket. Required by controller-gen.
func (in *TokenBucket) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
