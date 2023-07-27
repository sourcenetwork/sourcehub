// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sourcenetwork/sourcehub/acp/tx.proto

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

// Validate checks the field values on MsgCreatePolicy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MsgCreatePolicy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgCreatePolicy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgCreatePolicyMultiError, or nil if none found.
func (m *MsgCreatePolicy) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgCreatePolicy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Creator

	// no validation rules for Policy

	// no validation rules for MarshalType

	if all {
		switch v := interface{}(m.GetCreationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgCreatePolicyValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgCreatePolicyValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgCreatePolicyValidationError{
				field:  "CreationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MsgCreatePolicyMultiError(errors)
	}

	return nil
}

// MsgCreatePolicyMultiError is an error wrapping multiple validation errors
// returned by MsgCreatePolicy.ValidateAll() if the designated constraints
// aren't met.
type MsgCreatePolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgCreatePolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgCreatePolicyMultiError) AllErrors() []error { return m }

// MsgCreatePolicyValidationError is the validation error returned by
// MsgCreatePolicy.Validate if the designated constraints aren't met.
type MsgCreatePolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgCreatePolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgCreatePolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgCreatePolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgCreatePolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgCreatePolicyValidationError) ErrorName() string { return "MsgCreatePolicyValidationError" }

// Error satisfies the builtin error interface
func (e MsgCreatePolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgCreatePolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgCreatePolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgCreatePolicyValidationError{}

// Validate checks the field values on MsgCreatePolicyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgCreatePolicyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgCreatePolicyResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgCreatePolicyResponseMultiError, or nil if none found.
func (m *MsgCreatePolicyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgCreatePolicyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return MsgCreatePolicyResponseMultiError(errors)
	}

	return nil
}

// MsgCreatePolicyResponseMultiError is an error wrapping multiple validation
// errors returned by MsgCreatePolicyResponse.ValidateAll() if the designated
// constraints aren't met.
type MsgCreatePolicyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgCreatePolicyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgCreatePolicyResponseMultiError) AllErrors() []error { return m }

// MsgCreatePolicyResponseValidationError is the validation error returned by
// MsgCreatePolicyResponse.Validate if the designated constraints aren't met.
type MsgCreatePolicyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgCreatePolicyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgCreatePolicyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgCreatePolicyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgCreatePolicyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgCreatePolicyResponseValidationError) ErrorName() string {
	return "MsgCreatePolicyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MsgCreatePolicyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgCreatePolicyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgCreatePolicyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgCreatePolicyResponseValidationError{}

// Validate checks the field values on MsgCreateRelationship with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgCreateRelationship) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgCreateRelationship with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgCreateRelationshipMultiError, or nil if none found.
func (m *MsgCreateRelationship) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgCreateRelationship) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Creator

	// no validation rules for CreatorDid

	// no validation rules for PolicyId

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgCreateRelationshipValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgCreateRelationshipValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgCreateRelationshipValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	if all {
		switch v := interface{}(m.GetActor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgCreateRelationshipValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgCreateRelationshipValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgCreateRelationshipValidationError{
				field:  "Actor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ActorRelation

	if len(errors) > 0 {
		return MsgCreateRelationshipMultiError(errors)
	}

	return nil
}

// MsgCreateRelationshipMultiError is an error wrapping multiple validation
// errors returned by MsgCreateRelationship.ValidateAll() if the designated
// constraints aren't met.
type MsgCreateRelationshipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgCreateRelationshipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgCreateRelationshipMultiError) AllErrors() []error { return m }

// MsgCreateRelationshipValidationError is the validation error returned by
// MsgCreateRelationship.Validate if the designated constraints aren't met.
type MsgCreateRelationshipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgCreateRelationshipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgCreateRelationshipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgCreateRelationshipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgCreateRelationshipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgCreateRelationshipValidationError) ErrorName() string {
	return "MsgCreateRelationshipValidationError"
}

// Error satisfies the builtin error interface
func (e MsgCreateRelationshipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgCreateRelationship.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgCreateRelationshipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgCreateRelationshipValidationError{}

// Validate checks the field values on MsgCreateRelationshipResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgCreateRelationshipResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgCreateRelationshipResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// MsgCreateRelationshipResponseMultiError, or nil if none found.
func (m *MsgCreateRelationshipResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgCreateRelationshipResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Created

	if len(errors) > 0 {
		return MsgCreateRelationshipResponseMultiError(errors)
	}

	return nil
}

// MsgCreateRelationshipResponseMultiError is an error wrapping multiple
// validation errors returned by MsgCreateRelationshipResponse.ValidateAll()
// if the designated constraints aren't met.
type MsgCreateRelationshipResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgCreateRelationshipResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgCreateRelationshipResponseMultiError) AllErrors() []error { return m }

// MsgCreateRelationshipResponseValidationError is the validation error
// returned by MsgCreateRelationshipResponse.Validate if the designated
// constraints aren't met.
type MsgCreateRelationshipResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgCreateRelationshipResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgCreateRelationshipResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgCreateRelationshipResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgCreateRelationshipResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgCreateRelationshipResponseValidationError) ErrorName() string {
	return "MsgCreateRelationshipResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MsgCreateRelationshipResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgCreateRelationshipResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgCreateRelationshipResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgCreateRelationshipResponseValidationError{}

// Validate checks the field values on MsgDeleteRelationship with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgDeleteRelationship) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgDeleteRelationship with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgDeleteRelationshipMultiError, or nil if none found.
func (m *MsgDeleteRelationship) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgDeleteRelationship) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Creator

	// no validation rules for CreatorDid

	// no validation rules for PolicyId

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgDeleteRelationshipValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgDeleteRelationshipValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgDeleteRelationshipValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	if all {
		switch v := interface{}(m.GetActor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgDeleteRelationshipValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgDeleteRelationshipValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgDeleteRelationshipValidationError{
				field:  "Actor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ActorRelation

	if len(errors) > 0 {
		return MsgDeleteRelationshipMultiError(errors)
	}

	return nil
}

// MsgDeleteRelationshipMultiError is an error wrapping multiple validation
// errors returned by MsgDeleteRelationship.ValidateAll() if the designated
// constraints aren't met.
type MsgDeleteRelationshipMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgDeleteRelationshipMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgDeleteRelationshipMultiError) AllErrors() []error { return m }

// MsgDeleteRelationshipValidationError is the validation error returned by
// MsgDeleteRelationship.Validate if the designated constraints aren't met.
type MsgDeleteRelationshipValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgDeleteRelationshipValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgDeleteRelationshipValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgDeleteRelationshipValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgDeleteRelationshipValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgDeleteRelationshipValidationError) ErrorName() string {
	return "MsgDeleteRelationshipValidationError"
}

// Error satisfies the builtin error interface
func (e MsgDeleteRelationshipValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgDeleteRelationship.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgDeleteRelationshipValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgDeleteRelationshipValidationError{}

// Validate checks the field values on MsgDeleteRelationshipResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgDeleteRelationshipResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgDeleteRelationshipResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// MsgDeleteRelationshipResponseMultiError, or nil if none found.
func (m *MsgDeleteRelationshipResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgDeleteRelationshipResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Found

	if len(errors) > 0 {
		return MsgDeleteRelationshipResponseMultiError(errors)
	}

	return nil
}

// MsgDeleteRelationshipResponseMultiError is an error wrapping multiple
// validation errors returned by MsgDeleteRelationshipResponse.ValidateAll()
// if the designated constraints aren't met.
type MsgDeleteRelationshipResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgDeleteRelationshipResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgDeleteRelationshipResponseMultiError) AllErrors() []error { return m }

// MsgDeleteRelationshipResponseValidationError is the validation error
// returned by MsgDeleteRelationshipResponse.Validate if the designated
// constraints aren't met.
type MsgDeleteRelationshipResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgDeleteRelationshipResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgDeleteRelationshipResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgDeleteRelationshipResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgDeleteRelationshipResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgDeleteRelationshipResponseValidationError) ErrorName() string {
	return "MsgDeleteRelationshipResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MsgDeleteRelationshipResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgDeleteRelationshipResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgDeleteRelationshipResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgDeleteRelationshipResponseValidationError{}

// Validate checks the field values on MsgRegisterObject with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MsgRegisterObject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgRegisterObject with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgRegisterObjectMultiError, or nil if none found.
func (m *MsgRegisterObject) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgRegisterObject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Creator

	// no validation rules for CreatorDid

	// no validation rules for PolicyId

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgRegisterObjectValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgRegisterObjectValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgRegisterObjectValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MsgRegisterObjectMultiError(errors)
	}

	return nil
}

// MsgRegisterObjectMultiError is an error wrapping multiple validation errors
// returned by MsgRegisterObject.ValidateAll() if the designated constraints
// aren't met.
type MsgRegisterObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgRegisterObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgRegisterObjectMultiError) AllErrors() []error { return m }

// MsgRegisterObjectValidationError is the validation error returned by
// MsgRegisterObject.Validate if the designated constraints aren't met.
type MsgRegisterObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgRegisterObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgRegisterObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgRegisterObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgRegisterObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgRegisterObjectValidationError) ErrorName() string {
	return "MsgRegisterObjectValidationError"
}

// Error satisfies the builtin error interface
func (e MsgRegisterObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgRegisterObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgRegisterObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgRegisterObjectValidationError{}

// Validate checks the field values on MsgRegisterObjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgRegisterObjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgRegisterObjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgRegisterObjectResponseMultiError, or nil if none found.
func (m *MsgRegisterObjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgRegisterObjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgRegisterObjectResponseMultiError(errors)
	}

	return nil
}

// MsgRegisterObjectResponseMultiError is an error wrapping multiple validation
// errors returned by MsgRegisterObjectResponse.ValidateAll() if the
// designated constraints aren't met.
type MsgRegisterObjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgRegisterObjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgRegisterObjectResponseMultiError) AllErrors() []error { return m }

// MsgRegisterObjectResponseValidationError is the validation error returned by
// MsgRegisterObjectResponse.Validate if the designated constraints aren't met.
type MsgRegisterObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgRegisterObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgRegisterObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgRegisterObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgRegisterObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgRegisterObjectResponseValidationError) ErrorName() string {
	return "MsgRegisterObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MsgRegisterObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgRegisterObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgRegisterObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgRegisterObjectResponseValidationError{}

// Validate checks the field values on MsgUnregisterObject with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgUnregisterObject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgUnregisterObject with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgUnregisterObjectMultiError, or nil if none found.
func (m *MsgUnregisterObject) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgUnregisterObject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Creator

	// no validation rules for CreatorDid

	// no validation rules for PolicyId

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgUnregisterObjectValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgUnregisterObjectValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgUnregisterObjectValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MsgUnregisterObjectMultiError(errors)
	}

	return nil
}

// MsgUnregisterObjectMultiError is an error wrapping multiple validation
// errors returned by MsgUnregisterObject.ValidateAll() if the designated
// constraints aren't met.
type MsgUnregisterObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgUnregisterObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgUnregisterObjectMultiError) AllErrors() []error { return m }

// MsgUnregisterObjectValidationError is the validation error returned by
// MsgUnregisterObject.Validate if the designated constraints aren't met.
type MsgUnregisterObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgUnregisterObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgUnregisterObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgUnregisterObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgUnregisterObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgUnregisterObjectValidationError) ErrorName() string {
	return "MsgUnregisterObjectValidationError"
}

// Error satisfies the builtin error interface
func (e MsgUnregisterObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgUnregisterObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgUnregisterObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgUnregisterObjectValidationError{}

// Validate checks the field values on MsgUnregisterObjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgUnregisterObjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgUnregisterObjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgUnregisterObjectResponseMultiError, or nil if none found.
func (m *MsgUnregisterObjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgUnregisterObjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Found

	if len(errors) > 0 {
		return MsgUnregisterObjectResponseMultiError(errors)
	}

	return nil
}

// MsgUnregisterObjectResponseMultiError is an error wrapping multiple
// validation errors returned by MsgUnregisterObjectResponse.ValidateAll() if
// the designated constraints aren't met.
type MsgUnregisterObjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgUnregisterObjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgUnregisterObjectResponseMultiError) AllErrors() []error { return m }

// MsgUnregisterObjectResponseValidationError is the validation error returned
// by MsgUnregisterObjectResponse.Validate if the designated constraints
// aren't met.
type MsgUnregisterObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgUnregisterObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgUnregisterObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgUnregisterObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgUnregisterObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgUnregisterObjectResponseValidationError) ErrorName() string {
	return "MsgUnregisterObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MsgUnregisterObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgUnregisterObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgUnregisterObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgUnregisterObjectResponseValidationError{}
