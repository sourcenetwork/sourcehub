// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sourcenetwork/sourcehub/acp/policy_appdata.proto

package types

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

// Validate checks the field values on PolicyData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PolicyData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PolicyData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PolicyDataMultiError, or
// nil if none found.
func (m *PolicyData) ValidateAll() error {
	return m.validate(true)
}

func (m *PolicyData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAcpPolicy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PolicyDataValidationError{
					field:  "AcpPolicy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PolicyDataValidationError{
					field:  "AcpPolicy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAcpPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyDataValidationError{
				field:  "AcpPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetManagementGraph()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PolicyDataValidationError{
					field:  "ManagementGraph",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PolicyDataValidationError{
					field:  "ManagementGraph",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetManagementGraph()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PolicyDataValidationError{
				field:  "ManagementGraph",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PolicyDataMultiError(errors)
	}

	return nil
}

// PolicyDataMultiError is an error wrapping multiple validation errors
// returned by PolicyData.ValidateAll() if the designated constraints aren't met.
type PolicyDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolicyDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolicyDataMultiError) AllErrors() []error { return m }

// PolicyDataValidationError is the validation error returned by
// PolicyData.Validate if the designated constraints aren't met.
type PolicyDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolicyDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolicyDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolicyDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolicyDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolicyDataValidationError) ErrorName() string { return "PolicyDataValidationError" }

// Error satisfies the builtin error interface
func (e PolicyDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolicyData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolicyDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolicyDataValidationError{}

// Validate checks the field values on ManagementGraph with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ManagementGraph) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ManagementGraph with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ManagementGraphMultiError, or nil if none found.
func (m *ManagementGraph) ValidateAll() error {
	return m.validate(true)
}

func (m *ManagementGraph) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetNodes()))
		i := 0
		for key := range m.GetNodes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetNodes()[key]
			_ = val

			// no validation rules for Nodes[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, ManagementGraphValidationError{
							field:  fmt.Sprintf("Nodes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, ManagementGraphValidationError{
							field:  fmt.Sprintf("Nodes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return ManagementGraphValidationError{
						field:  fmt.Sprintf("Nodes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetForwardEdges()))
		i := 0
		for key := range m.GetForwardEdges() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetForwardEdges()[key]
			_ = val

			// no validation rules for ForwardEdges[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, ManagementGraphValidationError{
							field:  fmt.Sprintf("ForwardEdges[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, ManagementGraphValidationError{
							field:  fmt.Sprintf("ForwardEdges[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return ManagementGraphValidationError{
						field:  fmt.Sprintf("ForwardEdges[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetBackwardEdges()))
		i := 0
		for key := range m.GetBackwardEdges() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetBackwardEdges()[key]
			_ = val

			// no validation rules for BackwardEdges[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, ManagementGraphValidationError{
							field:  fmt.Sprintf("BackwardEdges[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, ManagementGraphValidationError{
							field:  fmt.Sprintf("BackwardEdges[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return ManagementGraphValidationError{
						field:  fmt.Sprintf("BackwardEdges[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return ManagementGraphMultiError(errors)
	}

	return nil
}

// ManagementGraphMultiError is an error wrapping multiple validation errors
// returned by ManagementGraph.ValidateAll() if the designated constraints
// aren't met.
type ManagementGraphMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ManagementGraphMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ManagementGraphMultiError) AllErrors() []error { return m }

// ManagementGraphValidationError is the validation error returned by
// ManagementGraph.Validate if the designated constraints aren't met.
type ManagementGraphValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ManagementGraphValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ManagementGraphValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ManagementGraphValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ManagementGraphValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ManagementGraphValidationError) ErrorName() string { return "ManagementGraphValidationError" }

// Error satisfies the builtin error interface
func (e ManagementGraphValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sManagementGraph.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ManagementGraphValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ManagementGraphValidationError{}

// Validate checks the field values on ManagerNode with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ManagerNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ManagerNode with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ManagerNodeMultiError, or
// nil if none found.
func (m *ManagerNode) ValidateAll() error {
	return m.validate(true)
}

func (m *ManagerNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	if len(errors) > 0 {
		return ManagerNodeMultiError(errors)
	}

	return nil
}

// ManagerNodeMultiError is an error wrapping multiple validation errors
// returned by ManagerNode.ValidateAll() if the designated constraints aren't met.
type ManagerNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ManagerNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ManagerNodeMultiError) AllErrors() []error { return m }

// ManagerNodeValidationError is the validation error returned by
// ManagerNode.Validate if the designated constraints aren't met.
type ManagerNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ManagerNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ManagerNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ManagerNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ManagerNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ManagerNodeValidationError) ErrorName() string { return "ManagerNodeValidationError" }

// Error satisfies the builtin error interface
func (e ManagerNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sManagerNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ManagerNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ManagerNodeValidationError{}

// Validate checks the field values on ManagerEdges with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ManagerEdges) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ManagerEdges with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ManagerEdgesMultiError, or
// nil if none found.
func (m *ManagerEdges) ValidateAll() error {
	return m.validate(true)
}

func (m *ManagerEdges) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Edges

	if len(errors) > 0 {
		return ManagerEdgesMultiError(errors)
	}

	return nil
}

// ManagerEdgesMultiError is an error wrapping multiple validation errors
// returned by ManagerEdges.ValidateAll() if the designated constraints aren't met.
type ManagerEdgesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ManagerEdgesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ManagerEdgesMultiError) AllErrors() []error { return m }

// ManagerEdgesValidationError is the validation error returned by
// ManagerEdges.Validate if the designated constraints aren't met.
type ManagerEdgesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ManagerEdgesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ManagerEdgesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ManagerEdgesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ManagerEdgesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ManagerEdgesValidationError) ErrorName() string { return "ManagerEdgesValidationError" }

// Error satisfies the builtin error interface
func (e ManagerEdgesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sManagerEdges.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ManagerEdgesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ManagerEdgesValidationError{}
