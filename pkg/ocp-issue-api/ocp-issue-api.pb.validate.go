// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-issue-api/ocp-issue-api.proto

package ocp_issue_api

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Issue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Issue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for ClassroomId

	// no validation rules for TaskId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetDeadline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssueValidationError{
				field:  "Deadline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// IssueValidationError is the validation error returned by Issue.Validate if
// the designated constraints aren't met.
type IssueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssueValidationError) ErrorName() string { return "IssueValidationError" }

// Error satisfies the builtin error interface
func (e IssueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssueValidationError{}

// Validate checks the field values on ListIssuesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListIssuesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListIssuesV1RequestValidationError is the validation error returned by
// ListIssuesV1Request.Validate if the designated constraints aren't met.
type ListIssuesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListIssuesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListIssuesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListIssuesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListIssuesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListIssuesV1RequestValidationError) ErrorName() string {
	return "ListIssuesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListIssuesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListIssuesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListIssuesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListIssuesV1RequestValidationError{}

// Validate checks the field values on ListIssuesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListIssuesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetIssues() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListIssuesV1ResponseValidationError{
					field:  fmt.Sprintf("Issues[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListIssuesV1ResponseValidationError is the validation error returned by
// ListIssuesV1Response.Validate if the designated constraints aren't met.
type ListIssuesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListIssuesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListIssuesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListIssuesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListIssuesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListIssuesV1ResponseValidationError) ErrorName() string {
	return "ListIssuesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListIssuesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListIssuesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListIssuesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListIssuesV1ResponseValidationError{}

// Validate checks the field values on CreateIssueV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateIssueV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ClassroomId

	// no validation rules for TaskId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetDeadline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateIssueV1RequestValidationError{
				field:  "Deadline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateIssueV1RequestValidationError is the validation error returned by
// CreateIssueV1Request.Validate if the designated constraints aren't met.
type CreateIssueV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIssueV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIssueV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIssueV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIssueV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIssueV1RequestValidationError) ErrorName() string {
	return "CreateIssueV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIssueV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIssueV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIssueV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIssueV1RequestValidationError{}

// Validate checks the field values on CreateIssueV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateIssueV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IssueId

	return nil
}

// CreateIssueV1ResponseValidationError is the validation error returned by
// CreateIssueV1Response.Validate if the designated constraints aren't met.
type CreateIssueV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIssueV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIssueV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIssueV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIssueV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIssueV1ResponseValidationError) ErrorName() string {
	return "CreateIssueV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIssueV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIssueV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIssueV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIssueV1ResponseValidationError{}

// Validate checks the field values on DescribeIssueV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeIssueV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IssueId

	return nil
}

// DescribeIssueV1RequestValidationError is the validation error returned by
// DescribeIssueV1Request.Validate if the designated constraints aren't met.
type DescribeIssueV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeIssueV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeIssueV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeIssueV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeIssueV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeIssueV1RequestValidationError) ErrorName() string {
	return "DescribeIssueV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeIssueV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeIssueV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeIssueV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeIssueV1RequestValidationError{}

// Validate checks the field values on DescribeIssueV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeIssueV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIssue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeIssueV1ResponseValidationError{
				field:  "Issue",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeIssueV1ResponseValidationError is the validation error returned by
// DescribeIssueV1Response.Validate if the designated constraints aren't met.
type DescribeIssueV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeIssueV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeIssueV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeIssueV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeIssueV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeIssueV1ResponseValidationError) ErrorName() string {
	return "DescribeIssueV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeIssueV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeIssueV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeIssueV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeIssueV1ResponseValidationError{}

// Validate checks the field values on UpdateIssueV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateIssueV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IssueId

	// no validation rules for ClassroomId

	// no validation rules for TaskId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetDeadline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateIssueV1RequestValidationError{
				field:  "Deadline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateIssueV1RequestValidationError is the validation error returned by
// UpdateIssueV1Request.Validate if the designated constraints aren't met.
type UpdateIssueV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateIssueV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateIssueV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateIssueV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateIssueV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateIssueV1RequestValidationError) ErrorName() string {
	return "UpdateIssueV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateIssueV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateIssueV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateIssueV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateIssueV1RequestValidationError{}

// Validate checks the field values on UpdateIssueV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateIssueV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// UpdateIssueV1ResponseValidationError is the validation error returned by
// UpdateIssueV1Response.Validate if the designated constraints aren't met.
type UpdateIssueV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateIssueV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateIssueV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateIssueV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateIssueV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateIssueV1ResponseValidationError) ErrorName() string {
	return "UpdateIssueV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateIssueV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateIssueV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateIssueV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateIssueV1ResponseValidationError{}

// Validate checks the field values on RemoveIssueV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveIssueV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IssueId

	return nil
}

// RemoveIssueV1RequestValidationError is the validation error returned by
// RemoveIssueV1Request.Validate if the designated constraints aren't met.
type RemoveIssueV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveIssueV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveIssueV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveIssueV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveIssueV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveIssueV1RequestValidationError) ErrorName() string {
	return "RemoveIssueV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveIssueV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveIssueV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveIssueV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveIssueV1RequestValidationError{}

// Validate checks the field values on RemoveIssueV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveIssueV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveIssueV1ResponseValidationError is the validation error returned by
// RemoveIssueV1Response.Validate if the designated constraints aren't met.
type RemoveIssueV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveIssueV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveIssueV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveIssueV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveIssueV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveIssueV1ResponseValidationError) ErrorName() string {
	return "RemoveIssueV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveIssueV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveIssueV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveIssueV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveIssueV1ResponseValidationError{}
