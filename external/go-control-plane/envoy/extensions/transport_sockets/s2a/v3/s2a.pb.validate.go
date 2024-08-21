// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/s2a/v3/s2a.proto

package envoy_extensions_transport_sockets_s2a_v3

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

// Validate checks the field values on S2AConfiguration with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *S2AConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetS2AAddress()) < 1 {
		return S2AConfigurationValidationError{
			field:  "S2AAddress",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// S2AConfigurationValidationError is the validation error returned by
// S2AConfiguration.Validate if the designated constraints aren't met.
type S2AConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e S2AConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e S2AConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e S2AConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e S2AConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e S2AConfigurationValidationError) ErrorName() string { return "S2AConfigurationValidationError" }

// Error satisfies the builtin error interface
func (e S2AConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sS2AConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = S2AConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = S2AConfigurationValidationError{}
