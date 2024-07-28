// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/dataVisualization/v1/dataVisualization.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on DeviceDataGetReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeviceDataGetReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeviceDataGetReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeviceDataGetReqMultiError, or nil if none found.
func (m *DeviceDataGetReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeviceDataGetReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Code

	// no validation rules for PageSize

	// no validation rules for StartTime

	// no validation rules for EndTime

	// no validation rules for CurrPage

	// no validation rules for UserId

	if len(errors) > 0 {
		return DeviceDataGetReqMultiError(errors)
	}

	return nil
}

// DeviceDataGetReqMultiError is an error wrapping multiple validation errors
// returned by DeviceDataGetReq.ValidateAll() if the designated constraints
// aren't met.
type DeviceDataGetReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeviceDataGetReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeviceDataGetReqMultiError) AllErrors() []error { return m }

// DeviceDataGetReqValidationError is the validation error returned by
// DeviceDataGetReq.Validate if the designated constraints aren't met.
type DeviceDataGetReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceDataGetReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceDataGetReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceDataGetReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceDataGetReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceDataGetReqValidationError) ErrorName() string { return "DeviceDataGetReqValidationError" }

// Error satisfies the builtin error interface
func (e DeviceDataGetReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceDataGetReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceDataGetReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceDataGetReqValidationError{}

// Validate checks the field values on DataDetail with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataDetailMultiError, or
// nil if none found.
func (m *DataDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *DataDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	// no validation rules for Unit

	if len(errors) > 0 {
		return DataDetailMultiError(errors)
	}

	return nil
}

// DataDetailMultiError is an error wrapping multiple validation errors
// returned by DataDetail.ValidateAll() if the designated constraints aren't met.
type DataDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataDetailMultiError) AllErrors() []error { return m }

// DataDetailValidationError is the validation error returned by
// DataDetail.Validate if the designated constraints aren't met.
type DataDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataDetailValidationError) ErrorName() string { return "DataDetailValidationError" }

// Error satisfies the builtin error interface
func (e DataDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataDetailValidationError{}

// Validate checks the field values on DeviceData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeviceData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeviceData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeviceDataMultiError, or
// nil if none found.
func (m *DeviceData) ValidateAll() error {
	return m.validate(true)
}

func (m *DeviceData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeviceDataValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeviceDataValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeviceDataValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CreateTime

	if len(errors) > 0 {
		return DeviceDataMultiError(errors)
	}

	return nil
}

// DeviceDataMultiError is an error wrapping multiple validation errors
// returned by DeviceData.ValidateAll() if the designated constraints aren't met.
type DeviceDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeviceDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeviceDataMultiError) AllErrors() []error { return m }

// DeviceDataValidationError is the validation error returned by
// DeviceData.Validate if the designated constraints aren't met.
type DeviceDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceDataValidationError) ErrorName() string { return "DeviceDataValidationError" }

// Error satisfies the builtin error interface
func (e DeviceDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceDataValidationError{}

// Validate checks the field values on DeviceDataGetRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeviceDataGetRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeviceDataGetRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeviceDataGetResMultiError, or nil if none found.
func (m *DeviceDataGetRes) ValidateAll() error {
	return m.validate(true)
}

func (m *DeviceDataGetRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Code

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeviceDataGetResValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeviceDataGetResValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeviceDataGetResValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Count

	if len(errors) > 0 {
		return DeviceDataGetResMultiError(errors)
	}

	return nil
}

// DeviceDataGetResMultiError is an error wrapping multiple validation errors
// returned by DeviceDataGetRes.ValidateAll() if the designated constraints
// aren't met.
type DeviceDataGetResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeviceDataGetResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeviceDataGetResMultiError) AllErrors() []error { return m }

// DeviceDataGetResValidationError is the validation error returned by
// DeviceDataGetRes.Validate if the designated constraints aren't met.
type DeviceDataGetResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceDataGetResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceDataGetResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceDataGetResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceDataGetResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceDataGetResValidationError) ErrorName() string { return "DeviceDataGetResValidationError" }

// Error satisfies the builtin error interface
func (e DeviceDataGetResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceDataGetRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceDataGetResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceDataGetResValidationError{}

// Validate checks the field values on DeviceDataCurveReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeviceDataCurveReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeviceDataCurveReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeviceDataCurveReqMultiError, or nil if none found.
func (m *DeviceDataCurveReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeviceDataCurveReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for DeviceCode

	// no validation rules for Interval

	if len(errors) > 0 {
		return DeviceDataCurveReqMultiError(errors)
	}

	return nil
}

// DeviceDataCurveReqMultiError is an error wrapping multiple validation errors
// returned by DeviceDataCurveReq.ValidateAll() if the designated constraints
// aren't met.
type DeviceDataCurveReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeviceDataCurveReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeviceDataCurveReqMultiError) AllErrors() []error { return m }

// DeviceDataCurveReqValidationError is the validation error returned by
// DeviceDataCurveReq.Validate if the designated constraints aren't met.
type DeviceDataCurveReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceDataCurveReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceDataCurveReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceDataCurveReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceDataCurveReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceDataCurveReqValidationError) ErrorName() string {
	return "DeviceDataCurveReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeviceDataCurveReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceDataCurveReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceDataCurveReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceDataCurveReqValidationError{}

// Validate checks the field values on DataLine with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataLine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataLine with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataLineMultiError, or nil
// if none found.
func (m *DataLine) ValidateAll() error {
	return m.validate(true)
}

func (m *DataLine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DataLineMultiError(errors)
	}

	return nil
}

// DataLineMultiError is an error wrapping multiple validation errors returned
// by DataLine.ValidateAll() if the designated constraints aren't met.
type DataLineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataLineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataLineMultiError) AllErrors() []error { return m }

// DataLineValidationError is the validation error returned by
// DataLine.Validate if the designated constraints aren't met.
type DataLineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataLineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataLineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataLineValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataLineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataLineValidationError) ErrorName() string { return "DataLineValidationError" }

// Error satisfies the builtin error interface
func (e DataLineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataLine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataLineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataLineValidationError{}

// Validate checks the field values on DeviceDataCurveRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeviceDataCurveRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeviceDataCurveRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeviceDataCurveResMultiError, or nil if none found.
func (m *DeviceDataCurveRes) ValidateAll() error {
	return m.validate(true)
}

func (m *DeviceDataCurveRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMonth()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Month",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Month",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMonth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceDataCurveResValidationError{
				field:  "Month",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPip()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Pip",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Pip",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPip()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceDataCurveResValidationError{
				field:  "Pip",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLine()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Line",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Line",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLine()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceDataCurveResValidationError{
				field:  "Line",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetElect()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Elect",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Elect",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetElect()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceDataCurveResValidationError{
				field:  "Elect",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetVolt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Volt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeviceDataCurveResValidationError{
					field:  "Volt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVolt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceDataCurveResValidationError{
				field:  "Volt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeviceDataCurveResMultiError(errors)
	}

	return nil
}

// DeviceDataCurveResMultiError is an error wrapping multiple validation errors
// returned by DeviceDataCurveRes.ValidateAll() if the designated constraints
// aren't met.
type DeviceDataCurveResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeviceDataCurveResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeviceDataCurveResMultiError) AllErrors() []error { return m }

// DeviceDataCurveResValidationError is the validation error returned by
// DeviceDataCurveRes.Validate if the designated constraints aren't met.
type DeviceDataCurveResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceDataCurveResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceDataCurveResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceDataCurveResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceDataCurveResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceDataCurveResValidationError) ErrorName() string {
	return "DeviceDataCurveResValidationError"
}

// Error satisfies the builtin error interface
func (e DeviceDataCurveResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceDataCurveRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceDataCurveResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceDataCurveResValidationError{}