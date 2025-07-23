package core

import (
	"errors"
	"fmt"
)

func NewError(code *ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewFormattedError(code *ErrorCode, format string, arguments ...any) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, arguments...),
	}
}

func NewErrorFrom(code int32, message string) *Error {
	err := &Error{Message: message}
	if ec, ok := errorCodeIndex[code]; ok {
		err.Code = ec
	} else {
		err.Code = &ErrorCode{Code: code}
	}
	return err
}

func NewFormattedErrorFrom(code int32, format string, arguments ...any) *Error {
	return NewErrorFrom(code, fmt.Sprintf(format, arguments...))
}

func (e *Error) Is(err error) bool {
	return IsError(err)
}

func IsError(err error) bool {
	return errors.Is(err, &Error{})
}

func AsError(err error) *Error {
	e := &Error{}
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if len(e.Message) == 0 {
		return e.Code.Name
	}
	return e.Message
}

func (e *Error) StatusCode() int {
	if e == nil {
		return 200
	}
	if e.Code.HttpStatusCode > 0 {
		return int(e.Code.HttpStatusCode)
	}
	return int(e.Code.Code)
}

func (e *Error) AddDetail(detail any) *Error {
	if e != nil {
		v, _ := NewValue(detail)
		e.Details = append(e.Details, v)
	}
	return e
}
