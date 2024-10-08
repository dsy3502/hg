// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/accesslog/v3/accesslog.proto

package envoy_config_accesslog_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on AccessLog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AccessLog) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.ConfigType.(type) {

	case *AccessLog_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AccessLogValidationError is the validation error returned by
// AccessLog.Validate if the designated constraints aren't met.
type AccessLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessLogValidationError) ErrorName() string { return "AccessLogValidationError" }

// Error satisfies the builtin error interface
func (e AccessLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessLogValidationError{}

// Validate checks the field values on AccessLogFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AccessLogFilter) Validate() error {
	if m == nil {
		return nil
	}

	switch m.FilterSpecifier.(type) {

	case *AccessLogFilter_StatusCodeFilter:

		if v, ok := interface{}(m.GetStatusCodeFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "StatusCodeFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_DurationFilter:

		if v, ok := interface{}(m.GetDurationFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "DurationFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_NotHealthCheckFilter:

		if v, ok := interface{}(m.GetNotHealthCheckFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "NotHealthCheckFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_TraceableFilter:

		if v, ok := interface{}(m.GetTraceableFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "TraceableFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_RuntimeFilter:

		if v, ok := interface{}(m.GetRuntimeFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "RuntimeFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_AndFilter:

		if v, ok := interface{}(m.GetAndFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "AndFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_OrFilter:

		if v, ok := interface{}(m.GetOrFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "OrFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_HeaderFilter:

		if v, ok := interface{}(m.GetHeaderFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "HeaderFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_ResponseFlagFilter:

		if v, ok := interface{}(m.GetResponseFlagFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "ResponseFlagFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_GrpcStatusFilter:

		if v, ok := interface{}(m.GetGrpcStatusFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "GrpcStatusFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_ExtensionFilter:

		if v, ok := interface{}(m.GetExtensionFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "ExtensionFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AccessLogFilter_MetadataFilter:

		if v, ok := interface{}(m.GetMetadataFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					field:  "MetadataFilter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return AccessLogFilterValidationError{
			field:  "FilterSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// AccessLogFilterValidationError is the validation error returned by
// AccessLogFilter.Validate if the designated constraints aren't met.
type AccessLogFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessLogFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessLogFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessLogFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessLogFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessLogFilterValidationError) ErrorName() string { return "AccessLogFilterValidationError" }

// Error satisfies the builtin error interface
func (e AccessLogFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLogFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessLogFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessLogFilterValidationError{}

// Validate checks the field values on ComparisonFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ComparisonFilter) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := ComparisonFilter_Op_name[int32(m.GetOp())]; !ok {
		return ComparisonFilterValidationError{
			field:  "Op",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComparisonFilterValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ComparisonFilterValidationError is the validation error returned by
// ComparisonFilter.Validate if the designated constraints aren't met.
type ComparisonFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComparisonFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComparisonFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComparisonFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComparisonFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComparisonFilterValidationError) ErrorName() string { return "ComparisonFilterValidationError" }

// Error satisfies the builtin error interface
func (e ComparisonFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComparisonFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComparisonFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComparisonFilterValidationError{}

// Validate checks the field values on StatusCodeFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StatusCodeFilter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetComparison() == nil {
		return StatusCodeFilterValidationError{
			field:  "Comparison",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetComparison()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusCodeFilterValidationError{
				field:  "Comparison",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StatusCodeFilterValidationError is the validation error returned by
// StatusCodeFilter.Validate if the designated constraints aren't met.
type StatusCodeFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusCodeFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusCodeFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusCodeFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusCodeFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusCodeFilterValidationError) ErrorName() string { return "StatusCodeFilterValidationError" }

// Error satisfies the builtin error interface
func (e StatusCodeFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusCodeFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusCodeFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusCodeFilterValidationError{}

// Validate checks the field values on DurationFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DurationFilter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetComparison() == nil {
		return DurationFilterValidationError{
			field:  "Comparison",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetComparison()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DurationFilterValidationError{
				field:  "Comparison",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DurationFilterValidationError is the validation error returned by
// DurationFilter.Validate if the designated constraints aren't met.
type DurationFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DurationFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DurationFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DurationFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DurationFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DurationFilterValidationError) ErrorName() string { return "DurationFilterValidationError" }

// Error satisfies the builtin error interface
func (e DurationFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDurationFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DurationFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DurationFilterValidationError{}

// Validate checks the field values on NotHealthCheckFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NotHealthCheckFilter) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// NotHealthCheckFilterValidationError is the validation error returned by
// NotHealthCheckFilter.Validate if the designated constraints aren't met.
type NotHealthCheckFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotHealthCheckFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotHealthCheckFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotHealthCheckFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotHealthCheckFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotHealthCheckFilterValidationError) ErrorName() string {
	return "NotHealthCheckFilterValidationError"
}

// Error satisfies the builtin error interface
func (e NotHealthCheckFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotHealthCheckFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotHealthCheckFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotHealthCheckFilterValidationError{}

// Validate checks the field values on TraceableFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TraceableFilter) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// TraceableFilterValidationError is the validation error returned by
// TraceableFilter.Validate if the designated constraints aren't met.
type TraceableFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraceableFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraceableFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraceableFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraceableFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraceableFilterValidationError) ErrorName() string { return "TraceableFilterValidationError" }

// Error satisfies the builtin error interface
func (e TraceableFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraceableFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraceableFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraceableFilterValidationError{}

// Validate checks the field values on RuntimeFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RuntimeFilter) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetRuntimeKey()) < 1 {
		return RuntimeFilterValidationError{
			field:  "RuntimeKey",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetPercentSampled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuntimeFilterValidationError{
				field:  "PercentSampled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UseIndependentRandomness

	return nil
}

// RuntimeFilterValidationError is the validation error returned by
// RuntimeFilter.Validate if the designated constraints aren't met.
type RuntimeFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeFilterValidationError) ErrorName() string { return "RuntimeFilterValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeFilterValidationError{}

// Validate checks the field values on AndFilter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AndFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFilters()) < 2 {
		return AndFilterValidationError{
			field:  "Filters",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AndFilterValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AndFilterValidationError is the validation error returned by
// AndFilter.Validate if the designated constraints aren't met.
type AndFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AndFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AndFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AndFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AndFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AndFilterValidationError) ErrorName() string { return "AndFilterValidationError" }

// Error satisfies the builtin error interface
func (e AndFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAndFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AndFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AndFilterValidationError{}

// Validate checks the field values on OrFilter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OrFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFilters()) < 2 {
		return OrFilterValidationError{
			field:  "Filters",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrFilterValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OrFilterValidationError is the validation error returned by
// OrFilter.Validate if the designated constraints aren't met.
type OrFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrFilterValidationError) ErrorName() string { return "OrFilterValidationError" }

// Error satisfies the builtin error interface
func (e OrFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrFilterValidationError{}

// Validate checks the field values on HeaderFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HeaderFilter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHeader() == nil {
		return HeaderFilterValidationError{
			field:  "Header",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderFilterValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HeaderFilterValidationError is the validation error returned by
// HeaderFilter.Validate if the designated constraints aren't met.
type HeaderFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderFilterValidationError) ErrorName() string { return "HeaderFilterValidationError" }

// Error satisfies the builtin error interface
func (e HeaderFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderFilterValidationError{}

// Validate checks the field values on ResponseFlagFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResponseFlagFilter) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetFlags() {
		_, _ = idx, item

		if _, ok := _ResponseFlagFilter_Flags_InLookup[item]; !ok {
			return ResponseFlagFilterValidationError{
				field:  fmt.Sprintf("Flags[%v]", idx),
				reason: "value must be in list [LH UH UT LR UR UF UC UO NR DI FI RL UAEX RLSE DC URX SI IH DPE UMSDR RFCF NFCF DT UPE NC OM]",
			}
		}

	}

	return nil
}

// ResponseFlagFilterValidationError is the validation error returned by
// ResponseFlagFilter.Validate if the designated constraints aren't met.
type ResponseFlagFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseFlagFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseFlagFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseFlagFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseFlagFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseFlagFilterValidationError) ErrorName() string {
	return "ResponseFlagFilterValidationError"
}

// Error satisfies the builtin error interface
func (e ResponseFlagFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseFlagFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseFlagFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseFlagFilterValidationError{}

var _ResponseFlagFilter_Flags_InLookup = map[string]struct{}{
	"LH":    {},
	"UH":    {},
	"UT":    {},
	"LR":    {},
	"UR":    {},
	"UF":    {},
	"UC":    {},
	"UO":    {},
	"NR":    {},
	"DI":    {},
	"FI":    {},
	"RL":    {},
	"UAEX":  {},
	"RLSE":  {},
	"DC":    {},
	"URX":   {},
	"SI":    {},
	"IH":    {},
	"DPE":   {},
	"UMSDR": {},
	"RFCF":  {},
	"NFCF":  {},
	"DT":    {},
	"UPE":   {},
	"NC":    {},
	"OM":    {},
}

// Validate checks the field values on GrpcStatusFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GrpcStatusFilter) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStatuses() {
		_, _ = idx, item

		if _, ok := GrpcStatusFilter_Status_name[int32(item)]; !ok {
			return GrpcStatusFilterValidationError{
				field:  fmt.Sprintf("Statuses[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
		}

	}

	// no validation rules for Exclude

	return nil
}

// GrpcStatusFilterValidationError is the validation error returned by
// GrpcStatusFilter.Validate if the designated constraints aren't met.
type GrpcStatusFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcStatusFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcStatusFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcStatusFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcStatusFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcStatusFilterValidationError) ErrorName() string { return "GrpcStatusFilterValidationError" }

// Error satisfies the builtin error interface
func (e GrpcStatusFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcStatusFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcStatusFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcStatusFilterValidationError{}

// Validate checks the field values on MetadataFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MetadataFilter) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataFilterValidationError{
				field:  "Matcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMatchIfKeyNotFound()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataFilterValidationError{
				field:  "MatchIfKeyNotFound",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MetadataFilterValidationError is the validation error returned by
// MetadataFilter.Validate if the designated constraints aren't met.
type MetadataFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataFilterValidationError) ErrorName() string { return "MetadataFilterValidationError" }

// Error satisfies the builtin error interface
func (e MetadataFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadataFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataFilterValidationError{}

// Validate checks the field values on ExtensionFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExtensionFilter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	switch m.ConfigType.(type) {

	case *ExtensionFilter_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtensionFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ExtensionFilterValidationError is the validation error returned by
// ExtensionFilter.Validate if the designated constraints aren't met.
type ExtensionFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtensionFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtensionFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtensionFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtensionFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtensionFilterValidationError) ErrorName() string { return "ExtensionFilterValidationError" }

// Error satisfies the builtin error interface
func (e ExtensionFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtensionFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtensionFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtensionFilterValidationError{}
