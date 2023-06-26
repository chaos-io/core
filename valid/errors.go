package valid

import (
	"fmt"
	"io"
	"strings"

	xerrors2 "github.com/chaos-io/core/xerrors"
)

var (
	ErrValidation = xerrors2.NewSentinel("validation error")

	ErrEmptyString         = xerrors2.NewSentinel("empty string given")
	ErrStringTooShort      = xerrors2.NewSentinel("given string too short")
	ErrStringTooLong       = xerrors2.NewSentinel("given string too long")
	ErrInvalidStringLength = xerrors2.NewSentinel("invalid string length")
	ErrBadFormat           = xerrors2.NewSentinel("bad string format")

	ErrInvalidPrefix        = xerrors2.NewSentinel("invalid prefix")
	ErrInvalidCharsSequence = xerrors2.NewSentinel("invalid characters sequence")
	ErrInvalidCharacters    = xerrors2.NewSentinel("invalid characters detected")
	ErrInvalidChecksum      = xerrors2.NewSentinel("invalid checksum")

	ErrBadParams      = xerrors2.NewSentinel("bad validation params")
	ErrStructExpected = xerrors2.NewSentinel("param expected to be struct")
	ErrInvalidType    = xerrors2.NewSentinel("one or more arguments have invalid type")

	ErrNotImplemented = xerrors2.NewSentinel("not implemented")
)

type Errors []error

// Error implements error type
func (es Errors) Error() string {
	return es.join("; ")
}

// String implements Stringer interface
func (es Errors) String() string {
	return es.join("\n")
}

// Format implements Formatter interface
func (es Errors) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			for _, e := range es {
				_, _ = io.WriteString(s, fmt.Sprintf("%+v", e))
			}
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, es.Error())
	}
}

// Has checks if Errors hold the specified error
func (es Errors) Has(err error) bool {
	for _, e := range es {
		if xerrors2.Is(e, err) {
			return true
		}
	}
	return false
}

// joins errors into single string with given glue
func (es Errors) join(glue string) string {
	if len(es) == 0 {
		return ""
	}

	var b strings.Builder
	for i, e := range es {
		b.WriteString(e.Error())
		if i < len(es)-1 {
			b.WriteString(glue)
		}
	}
	return b.String()
}

// FieldError holds additional information about struct field validation error
type FieldError struct {
	field string
	path  string
	err   error
}

// Error implements error type
func (e FieldError) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

// Implements xerrors.Is interface
func (e FieldError) Is(target error) bool {
	return xerrors2.Is(e.err, target)
}

// Implements xerrors.As interface
func (e FieldError) As(target interface{}) bool {
	return xerrors2.As(e.err, target)
}

// Path returns path to invalid struct field starting from top struct.
func (e FieldError) Path() string {
	return e.path
}

// Field returns invalid struct field name.
func (e FieldError) Field() string {
	return e.field
}

// Format implements Formatter interface
func (e FieldError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%s.%s: %s", e.path, e.field, e.Error())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, e.Error())
	}
}
