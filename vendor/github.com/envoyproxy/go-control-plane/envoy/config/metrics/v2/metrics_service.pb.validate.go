// Code generated by protoc-gen-validate
// source: envoy/config/metrics/v2/metrics_service.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on MetricsServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MetricsServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return MetricsServiceConfigValidationError{
			Field:  "GrpcService",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGrpcService()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return MetricsServiceConfigValidationError{
				Field:  "GrpcService",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// MetricsServiceConfigValidationError is the validation error returned by
// MetricsServiceConfig.Validate if the designated constraints aren't met.
type MetricsServiceConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e MetricsServiceConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsServiceConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = MetricsServiceConfigValidationError{}
