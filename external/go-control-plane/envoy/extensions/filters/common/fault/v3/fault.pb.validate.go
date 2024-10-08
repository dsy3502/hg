// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/common/fault/v3/fault.proto

package envoy_extensions_filters_common_fault_v3

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

// Validate checks the field values on FaultDelay with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FaultDelay) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultDelayValidationError{
				field:  "Percentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.FaultDelaySecifier.(type) {

	case *FaultDelay_FixedDelay:

		if d := m.GetFixedDelay(); d != nil {
			dur, err := d.AsDuration(), d.CheckValid()
			if err != nil {
				return FaultDelayValidationError{
					field:  "FixedDelay",
					reason: "value is not a valid duration",
					cause:  err,
				}
			}

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				return FaultDelayValidationError{
					field:  "FixedDelay",
					reason: "value must be greater than 0s",
				}
			}

		}

	case *FaultDelay_HeaderDelay_:

		if v, ok := interface{}(m.GetHeaderDelay()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultDelayValidationError{
					field:  "HeaderDelay",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return FaultDelayValidationError{
			field:  "FaultDelaySecifier",
			reason: "value is required",
		}

	}

	return nil
}

// FaultDelayValidationError is the validation error returned by
// FaultDelay.Validate if the designated constraints aren't met.
type FaultDelayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultDelayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultDelayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultDelayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultDelayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultDelayValidationError) ErrorName() string { return "FaultDelayValidationError" }

// Error satisfies the builtin error interface
func (e FaultDelayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultDelay.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultDelayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultDelayValidationError{}

// Validate checks the field values on FaultRateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FaultRateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultRateLimitValidationError{
				field:  "Percentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.LimitType.(type) {

	case *FaultRateLimit_FixedLimit_:

		if v, ok := interface{}(m.GetFixedLimit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultRateLimitValidationError{
					field:  "FixedLimit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FaultRateLimit_HeaderLimit_:

		if v, ok := interface{}(m.GetHeaderLimit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultRateLimitValidationError{
					field:  "HeaderLimit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return FaultRateLimitValidationError{
			field:  "LimitType",
			reason: "value is required",
		}

	}

	return nil
}

// FaultRateLimitValidationError is the validation error returned by
// FaultRateLimit.Validate if the designated constraints aren't met.
type FaultRateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultRateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultRateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultRateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultRateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultRateLimitValidationError) ErrorName() string { return "FaultRateLimitValidationError" }

// Error satisfies the builtin error interface
func (e FaultRateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultRateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultRateLimitValidationError{}

// Validate checks the field values on FaultDelay_HeaderDelay with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultDelay_HeaderDelay) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FaultDelay_HeaderDelayValidationError is the validation error returned by
// FaultDelay_HeaderDelay.Validate if the designated constraints aren't met.
type FaultDelay_HeaderDelayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultDelay_HeaderDelayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultDelay_HeaderDelayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultDelay_HeaderDelayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultDelay_HeaderDelayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultDelay_HeaderDelayValidationError) ErrorName() string {
	return "FaultDelay_HeaderDelayValidationError"
}

// Error satisfies the builtin error interface
func (e FaultDelay_HeaderDelayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultDelay_HeaderDelay.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultDelay_HeaderDelayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultDelay_HeaderDelayValidationError{}

// Validate checks the field values on FaultRateLimit_FixedLimit with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultRateLimit_FixedLimit) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimitKbps() < 1 {
		return FaultRateLimit_FixedLimitValidationError{
			field:  "LimitKbps",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// FaultRateLimit_FixedLimitValidationError is the validation error returned by
// FaultRateLimit_FixedLimit.Validate if the designated constraints aren't met.
type FaultRateLimit_FixedLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultRateLimit_FixedLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultRateLimit_FixedLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultRateLimit_FixedLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultRateLimit_FixedLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultRateLimit_FixedLimitValidationError) ErrorName() string {
	return "FaultRateLimit_FixedLimitValidationError"
}

// Error satisfies the builtin error interface
func (e FaultRateLimit_FixedLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultRateLimit_FixedLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultRateLimit_FixedLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultRateLimit_FixedLimitValidationError{}

// Validate checks the field values on FaultRateLimit_HeaderLimit with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultRateLimit_HeaderLimit) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FaultRateLimit_HeaderLimitValidationError is the validation error returned
// by FaultRateLimit_HeaderLimit.Validate if the designated constraints aren't met.
type FaultRateLimit_HeaderLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultRateLimit_HeaderLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultRateLimit_HeaderLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultRateLimit_HeaderLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultRateLimit_HeaderLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultRateLimit_HeaderLimitValidationError) ErrorName() string {
	return "FaultRateLimit_HeaderLimitValidationError"
}

// Error satisfies the builtin error interface
func (e FaultRateLimit_HeaderLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultRateLimit_HeaderLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultRateLimit_HeaderLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultRateLimit_HeaderLimitValidationError{}
