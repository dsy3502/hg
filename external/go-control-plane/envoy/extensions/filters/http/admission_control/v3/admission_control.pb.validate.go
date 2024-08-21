// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/admission_control/v3/admission_control.proto

package envoy_extensions_filters_http_admission_control_v3

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

// Validate checks the field values on AdmissionControl with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AdmissionControl) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControlValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSamplingWindow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControlValidationError{
				field:  "SamplingWindow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAggression()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControlValidationError{
				field:  "Aggression",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSrThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControlValidationError{
				field:  "SrThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRpsThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControlValidationError{
				field:  "RpsThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxRejectionProbability()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControlValidationError{
				field:  "MaxRejectionProbability",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.EvaluationCriteria.(type) {

	case *AdmissionControl_SuccessCriteria_:

		if v, ok := interface{}(m.GetSuccessCriteria()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdmissionControlValidationError{
					field:  "SuccessCriteria",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return AdmissionControlValidationError{
			field:  "EvaluationCriteria",
			reason: "value is required",
		}

	}

	return nil
}

// AdmissionControlValidationError is the validation error returned by
// AdmissionControl.Validate if the designated constraints aren't met.
type AdmissionControlValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdmissionControlValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdmissionControlValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdmissionControlValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdmissionControlValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdmissionControlValidationError) ErrorName() string { return "AdmissionControlValidationError" }

// Error satisfies the builtin error interface
func (e AdmissionControlValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmissionControl.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdmissionControlValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdmissionControlValidationError{}

// Validate checks the field values on AdmissionControl_SuccessCriteria with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *AdmissionControl_SuccessCriteria) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttpCriteria()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControl_SuccessCriteriaValidationError{
				field:  "HttpCriteria",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetGrpcCriteria()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdmissionControl_SuccessCriteriaValidationError{
				field:  "GrpcCriteria",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AdmissionControl_SuccessCriteriaValidationError is the validation error
// returned by AdmissionControl_SuccessCriteria.Validate if the designated
// constraints aren't met.
type AdmissionControl_SuccessCriteriaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdmissionControl_SuccessCriteriaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdmissionControl_SuccessCriteriaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdmissionControl_SuccessCriteriaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdmissionControl_SuccessCriteriaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdmissionControl_SuccessCriteriaValidationError) ErrorName() string {
	return "AdmissionControl_SuccessCriteriaValidationError"
}

// Error satisfies the builtin error interface
func (e AdmissionControl_SuccessCriteriaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmissionControl_SuccessCriteria.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdmissionControl_SuccessCriteriaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdmissionControl_SuccessCriteriaValidationError{}

// Validate checks the field values on
// AdmissionControl_SuccessCriteria_HttpCriteria with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AdmissionControl_SuccessCriteria_HttpCriteria) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetHttpSuccessStatus()) < 1 {
		return AdmissionControl_SuccessCriteria_HttpCriteriaValidationError{
			field:  "HttpSuccessStatus",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetHttpSuccessStatus() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdmissionControl_SuccessCriteria_HttpCriteriaValidationError{
					field:  fmt.Sprintf("HttpSuccessStatus[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AdmissionControl_SuccessCriteria_HttpCriteriaValidationError is the
// validation error returned by
// AdmissionControl_SuccessCriteria_HttpCriteria.Validate if the designated
// constraints aren't met.
type AdmissionControl_SuccessCriteria_HttpCriteriaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdmissionControl_SuccessCriteria_HttpCriteriaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdmissionControl_SuccessCriteria_HttpCriteriaValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e AdmissionControl_SuccessCriteria_HttpCriteriaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdmissionControl_SuccessCriteria_HttpCriteriaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdmissionControl_SuccessCriteria_HttpCriteriaValidationError) ErrorName() string {
	return "AdmissionControl_SuccessCriteria_HttpCriteriaValidationError"
}

// Error satisfies the builtin error interface
func (e AdmissionControl_SuccessCriteria_HttpCriteriaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmissionControl_SuccessCriteria_HttpCriteria.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdmissionControl_SuccessCriteria_HttpCriteriaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdmissionControl_SuccessCriteria_HttpCriteriaValidationError{}

// Validate checks the field values on
// AdmissionControl_SuccessCriteria_GrpcCriteria with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AdmissionControl_SuccessCriteria_GrpcCriteria) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetGrpcSuccessStatus()) < 1 {
		return AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError{
			field:  "GrpcSuccessStatus",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError is the
// validation error returned by
// AdmissionControl_SuccessCriteria_GrpcCriteria.Validate if the designated
// constraints aren't met.
type AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError) ErrorName() string {
	return "AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError"
}

// Error satisfies the builtin error interface
func (e AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmissionControl_SuccessCriteria_GrpcCriteria.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdmissionControl_SuccessCriteria_GrpcCriteriaValidationError{}
