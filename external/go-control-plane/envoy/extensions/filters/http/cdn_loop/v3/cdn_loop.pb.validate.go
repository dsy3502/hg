// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/cdn_loop/v3/cdn_loop.proto

package envoy_extensions_filters_http_cdn_loop_v3

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

// Validate checks the field values on CdnLoopConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CdnLoopConfig) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetCdnId()) < 1 {
		return CdnLoopConfigValidationError{
			field:  "CdnId",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for MaxAllowedOccurrences

	return nil
}

// CdnLoopConfigValidationError is the validation error returned by
// CdnLoopConfig.Validate if the designated constraints aren't met.
type CdnLoopConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CdnLoopConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CdnLoopConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CdnLoopConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CdnLoopConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CdnLoopConfigValidationError) ErrorName() string { return "CdnLoopConfigValidationError" }

// Error satisfies the builtin error interface
func (e CdnLoopConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCdnLoopConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CdnLoopConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CdnLoopConfigValidationError{}
