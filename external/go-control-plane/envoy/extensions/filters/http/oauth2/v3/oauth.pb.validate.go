// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/oauth2/v3/oauth.proto

package envoy_extensions_filters_http_oauth2_v3

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

// Validate checks the field values on OAuth2Credentials with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OAuth2Credentials) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetClientId()) < 1 {
		return OAuth2CredentialsValidationError{
			field:  "ClientId",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetTokenSecret() == nil {
		return OAuth2CredentialsValidationError{
			field:  "TokenSecret",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTokenSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2CredentialsValidationError{
				field:  "TokenSecret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.TokenFormation.(type) {

	case *OAuth2Credentials_HmacSecret:

		if m.GetHmacSecret() == nil {
			return OAuth2CredentialsValidationError{
				field:  "HmacSecret",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetHmacSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OAuth2CredentialsValidationError{
					field:  "HmacSecret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return OAuth2CredentialsValidationError{
			field:  "TokenFormation",
			reason: "value is required",
		}

	}

	return nil
}

// OAuth2CredentialsValidationError is the validation error returned by
// OAuth2Credentials.Validate if the designated constraints aren't met.
type OAuth2CredentialsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2CredentialsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2CredentialsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2CredentialsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2CredentialsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2CredentialsValidationError) ErrorName() string {
	return "OAuth2CredentialsValidationError"
}

// Error satisfies the builtin error interface
func (e OAuth2CredentialsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Credentials.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2CredentialsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2CredentialsValidationError{}

// Validate checks the field values on OAuth2Config with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OAuth2Config) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTokenEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "TokenEndpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetAuthorizationEndpoint()) < 1 {
		return OAuth2ConfigValidationError{
			field:  "AuthorizationEndpoint",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetCredentials() == nil {
		return OAuth2ConfigValidationError{
			field:  "Credentials",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "Credentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetRedirectUri()) < 1 {
		return OAuth2ConfigValidationError{
			field:  "RedirectUri",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetRedirectPathMatcher() == nil {
		return OAuth2ConfigValidationError{
			field:  "RedirectPathMatcher",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRedirectPathMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "RedirectPathMatcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignoutPath() == nil {
		return OAuth2ConfigValidationError{
			field:  "SignoutPath",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSignoutPath()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "SignoutPath",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ForwardBearerToken

	for idx, item := range m.GetPassThroughMatcher() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OAuth2ConfigValidationError{
					field:  fmt.Sprintf("PassThroughMatcher[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OAuth2ConfigValidationError is the validation error returned by
// OAuth2Config.Validate if the designated constraints aren't met.
type OAuth2ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ConfigValidationError) ErrorName() string { return "OAuth2ConfigValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Config.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ConfigValidationError{}

// Validate checks the field values on OAuth2 with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OAuth2) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OAuth2ValidationError is the validation error returned by OAuth2.Validate if
// the designated constraints aren't met.
type OAuth2ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ValidationError) ErrorName() string { return "OAuth2ValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ValidationError{}
