package valid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
)

func TestErrorsError(t *testing.T) {
	testCases := []struct {
		name     string
		errs     valid2.Errors
		expected string
	}{
		{"no_errors", valid2.Errors(nil), ""},
		{"one_error", valid2.Errors{valid2.ErrEmptyString}, "empty string given"},
		{"multiple_errors", valid2.Errors{valid2.ErrEmptyString, valid2.ErrInvalidPrefix}, "empty string given; invalid prefix"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.errs.Error())
		})
	}
}

func TestErrorsString(t *testing.T) {
	testCases := []struct {
		name     string
		errs     valid2.Errors
		expected string
	}{
		{"no_errors", valid2.Errors(nil), ""},
		{"one_error", valid2.Errors{valid2.ErrEmptyString}, "empty string given"},
		{"multiple_errors", valid2.Errors{valid2.ErrEmptyString, valid2.ErrInvalidPrefix}, "empty string given\ninvalid prefix"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.errs.String())
		})
	}
}

func TestErrorsHas(t *testing.T) {
	testCases := []struct {
		name     string
		errs     valid2.Errors
		err      error
		expected bool
	}{
		{"empty_errors", valid2.Errors(nil), valid2.ErrEmptyString, false},
		{"has_ErrEmptyString", valid2.Errors{valid2.ErrEmptyString}, valid2.ErrEmptyString, true},
		{"has_wrapped_ErrEmptyString", valid2.Errors{valid2.ErrValidation.Wrap(valid2.ErrEmptyString)}, valid2.ErrEmptyString, true},
		{"has_not_ErrEmptyString", valid2.Errors{valid2.ErrInvalidPrefix}, valid2.ErrEmptyString, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.errs.Has(tc.err))
		})
	}
}

func BenchmarkErrorsString(b *testing.B) {
	benchCases := []valid2.Errors{
		valid2.Errors(nil),
		{valid2.ErrEmptyString},
		{valid2.ErrEmptyString, valid2.ErrInvalidPrefix},
		{valid2.ErrEmptyString, valid2.ErrInvalidPrefix, valid2.ErrBadParams},
		{valid2.ErrEmptyString, valid2.ErrInvalidPrefix, valid2.ErrBadParams, valid2.ErrEmptyDataPart},
		{valid2.ErrEmptyString, valid2.ErrInvalidPrefix, valid2.ErrBadParams, valid2.ErrEmptyDataPart, valid2.ErrInvalidISBN},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = benchCases[i%len(benchCases)].String()
	}
}
