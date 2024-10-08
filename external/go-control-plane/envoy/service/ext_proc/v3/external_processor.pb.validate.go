// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/ext_proc/v3/external_processor.proto

package envoy_service_ext_proc_v3

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

// Validate checks the field values on ProcessingRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProcessingRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AsyncMode

	switch m.Request.(type) {

	case *ProcessingRequest_RequestHeaders:

		if v, ok := interface{}(m.GetRequestHeaders()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingRequestValidationError{
					field:  "RequestHeaders",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingRequest_ResponseHeaders:

		if v, ok := interface{}(m.GetResponseHeaders()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingRequestValidationError{
					field:  "ResponseHeaders",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingRequest_RequestBody:

		if v, ok := interface{}(m.GetRequestBody()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingRequestValidationError{
					field:  "RequestBody",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingRequest_ResponseBody:

		if v, ok := interface{}(m.GetResponseBody()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingRequestValidationError{
					field:  "ResponseBody",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingRequest_RequestTrailers:

		if v, ok := interface{}(m.GetRequestTrailers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingRequestValidationError{
					field:  "RequestTrailers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingRequest_ResponseTrailers:

		if v, ok := interface{}(m.GetResponseTrailers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingRequestValidationError{
					field:  "ResponseTrailers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ProcessingRequestValidationError{
			field:  "Request",
			reason: "value is required",
		}

	}

	return nil
}

// ProcessingRequestValidationError is the validation error returned by
// ProcessingRequest.Validate if the designated constraints aren't met.
type ProcessingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessingRequestValidationError) ErrorName() string {
	return "ProcessingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProcessingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessingRequestValidationError{}

// Validate checks the field values on ProcessingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProcessingResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDynamicMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingResponseValidationError{
				field:  "DynamicMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetModeOverride()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingResponseValidationError{
				field:  "ModeOverride",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Response.(type) {

	case *ProcessingResponse_RequestHeaders:

		if v, ok := interface{}(m.GetRequestHeaders()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "RequestHeaders",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingResponse_ResponseHeaders:

		if v, ok := interface{}(m.GetResponseHeaders()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "ResponseHeaders",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingResponse_RequestBody:

		if v, ok := interface{}(m.GetRequestBody()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "RequestBody",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingResponse_ResponseBody:

		if v, ok := interface{}(m.GetResponseBody()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "ResponseBody",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingResponse_RequestTrailers:

		if v, ok := interface{}(m.GetRequestTrailers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "RequestTrailers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingResponse_ResponseTrailers:

		if v, ok := interface{}(m.GetResponseTrailers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "ResponseTrailers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProcessingResponse_ImmediateResponse:

		if v, ok := interface{}(m.GetImmediateResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProcessingResponseValidationError{
					field:  "ImmediateResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ProcessingResponseValidationError{
			field:  "Response",
			reason: "value is required",
		}

	}

	return nil
}

// ProcessingResponseValidationError is the validation error returned by
// ProcessingResponse.Validate if the designated constraints aren't met.
type ProcessingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessingResponseValidationError) ErrorName() string {
	return "ProcessingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProcessingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessingResponseValidationError{}

// Validate checks the field values on HttpHeaders with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpHeaders) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpHeadersValidationError{
				field:  "Headers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetAttributes() {
		_ = val

		// no validation rules for Attributes[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpHeadersValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for EndOfStream

	return nil
}

// HttpHeadersValidationError is the validation error returned by
// HttpHeaders.Validate if the designated constraints aren't met.
type HttpHeadersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpHeadersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpHeadersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpHeadersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpHeadersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpHeadersValidationError) ErrorName() string { return "HttpHeadersValidationError" }

// Error satisfies the builtin error interface
func (e HttpHeadersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpHeaders.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpHeadersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpHeadersValidationError{}

// Validate checks the field values on HttpBody with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpBody) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Body

	// no validation rules for EndOfStream

	return nil
}

// HttpBodyValidationError is the validation error returned by
// HttpBody.Validate if the designated constraints aren't met.
type HttpBodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpBodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpBodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpBodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpBodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpBodyValidationError) ErrorName() string { return "HttpBodyValidationError" }

// Error satisfies the builtin error interface
func (e HttpBodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpBody.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpBodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpBodyValidationError{}

// Validate checks the field values on HttpTrailers with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpTrailers) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTrailers()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpTrailersValidationError{
				field:  "Trailers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpTrailersValidationError is the validation error returned by
// HttpTrailers.Validate if the designated constraints aren't met.
type HttpTrailersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpTrailersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpTrailersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpTrailersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpTrailersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpTrailersValidationError) ErrorName() string { return "HttpTrailersValidationError" }

// Error satisfies the builtin error interface
func (e HttpTrailersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpTrailers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpTrailersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpTrailersValidationError{}

// Validate checks the field values on HeadersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeadersResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeadersResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HeadersResponseValidationError is the validation error returned by
// HeadersResponse.Validate if the designated constraints aren't met.
type HeadersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeadersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeadersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeadersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeadersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeadersResponseValidationError) ErrorName() string { return "HeadersResponseValidationError" }

// Error satisfies the builtin error interface
func (e HeadersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeadersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeadersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeadersResponseValidationError{}

// Validate checks the field values on TrailersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TrailersResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHeaderMutation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TrailersResponseValidationError{
				field:  "HeaderMutation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TrailersResponseValidationError is the validation error returned by
// TrailersResponse.Validate if the designated constraints aren't met.
type TrailersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrailersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrailersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrailersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrailersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrailersResponseValidationError) ErrorName() string { return "TrailersResponseValidationError" }

// Error satisfies the builtin error interface
func (e TrailersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrailersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrailersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrailersResponseValidationError{}

// Validate checks the field values on BodyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BodyResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BodyResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BodyResponseValidationError is the validation error returned by
// BodyResponse.Validate if the designated constraints aren't met.
type BodyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BodyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BodyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BodyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BodyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BodyResponseValidationError) ErrorName() string { return "BodyResponseValidationError" }

// Error satisfies the builtin error interface
func (e BodyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBodyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BodyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BodyResponseValidationError{}

// Validate checks the field values on CommonResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CommonResponse) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := CommonResponse_ResponseStatus_name[int32(m.GetStatus())]; !ok {
		return CommonResponseValidationError{
			field:  "Status",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetHeaderMutation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonResponseValidationError{
				field:  "HeaderMutation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBodyMutation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonResponseValidationError{
				field:  "BodyMutation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTrailers()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonResponseValidationError{
				field:  "Trailers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ClearRouteCache

	return nil
}

// CommonResponseValidationError is the validation error returned by
// CommonResponse.Validate if the designated constraints aren't met.
type CommonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonResponseValidationError) ErrorName() string { return "CommonResponseValidationError" }

// Error satisfies the builtin error interface
func (e CommonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonResponseValidationError{}

// Validate checks the field values on ImmediateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ImmediateResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStatus() == nil {
		return ImmediateResponseValidationError{
			field:  "Status",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImmediateResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImmediateResponseValidationError{
				field:  "Headers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Body

	if v, ok := interface{}(m.GetGrpcStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImmediateResponseValidationError{
				field:  "GrpcStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Details

	return nil
}

// ImmediateResponseValidationError is the validation error returned by
// ImmediateResponse.Validate if the designated constraints aren't met.
type ImmediateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImmediateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImmediateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImmediateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImmediateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImmediateResponseValidationError) ErrorName() string {
	return "ImmediateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ImmediateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImmediateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImmediateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImmediateResponseValidationError{}

// Validate checks the field values on GrpcStatus with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GrpcStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// GrpcStatusValidationError is the validation error returned by
// GrpcStatus.Validate if the designated constraints aren't met.
type GrpcStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcStatusValidationError) ErrorName() string { return "GrpcStatusValidationError" }

// Error satisfies the builtin error interface
func (e GrpcStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcStatusValidationError{}

// Validate checks the field values on HeaderMutation with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HeaderMutation) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HeaderMutationValidationError{
					field:  fmt.Sprintf("SetHeaders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HeaderMutationValidationError is the validation error returned by
// HeaderMutation.Validate if the designated constraints aren't met.
type HeaderMutationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderMutationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderMutationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderMutationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderMutationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderMutationValidationError) ErrorName() string { return "HeaderMutationValidationError" }

// Error satisfies the builtin error interface
func (e HeaderMutationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderMutation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderMutationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderMutationValidationError{}

// Validate checks the field values on BodyMutation with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BodyMutation) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Mutation.(type) {

	case *BodyMutation_Body:
		// no validation rules for Body

	case *BodyMutation_ClearBody:
		// no validation rules for ClearBody

	}

	return nil
}

// BodyMutationValidationError is the validation error returned by
// BodyMutation.Validate if the designated constraints aren't met.
type BodyMutationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BodyMutationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BodyMutationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BodyMutationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BodyMutationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BodyMutationValidationError) ErrorName() string { return "BodyMutationValidationError" }

// Error satisfies the builtin error interface
func (e BodyMutationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBodyMutation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BodyMutationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BodyMutationValidationError{}
