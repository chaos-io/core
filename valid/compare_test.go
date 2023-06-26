package valid_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
)

type someCustomType struct{}

func TestEqual(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		value       interface{}
		param       string
		expectedErr error
	}{
		{int(10), "10", nil},
		{int8(78), "78", nil},
		{int16(52), "52", nil},
		{now, now.Format(time.RFC3339Nano), nil},
		{true, "true", nil},
		{"shimba", "shimba", nil},
		{14.42, "55", valid2.ErrNotEqual},
		{"ololo", "55", valid2.ErrNotEqual},
		{someCustomType{}, "42", valid2.ErrInvalidType},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := valid2.Equal(reflect.ValueOf(tc.value), tc.param)

			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		value       interface{}
		param       string
		expectedErr error
	}{
		{int(10), "11", nil},
		{int8(78), "79", nil},
		{int16(54), "52", nil},
		{now, now.Add(1 * time.Minute).Format(time.RFC3339Nano), nil},
		{true, "false", nil},
		{"ololo", "trololo", nil},
		{14.42, "14.42", valid2.ErrEqual},
		{"ololo", "ololo", valid2.ErrEqual},
		{someCustomType{}, "10", valid2.ErrInvalidType},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := valid2.NotEqual(reflect.ValueOf(tc.value), tc.param)

			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestMin(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		value       interface{}
		param       string
		expectedErr error
	}{
		{int(14), "10", nil},
		{int8(78), "42", nil},
		{int16(52), "52", nil},
		{now, now.Format(time.RFC3339Nano), nil},
		{now.Add(1 * time.Minute), now.Format(time.RFC3339Nano), nil},
		{14.42, "55.34", valid2.ErrLesserValue},
		{someCustomType{}, "10", valid2.ErrInvalidType},
		{"ololo", "olol", valid2.ErrUnsupportedComparator},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := valid2.Min(reflect.ValueOf(tc.value), tc.param)

			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestMax(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		value       interface{}
		param       string
		expectedErr error
	}{
		{10, "14", nil},
		{int8(42), "78", nil},
		{int16(52), "52", nil},
		{now, now.Format(time.RFC3339Nano), nil},
		{now, now.Add(1 * time.Minute).Format(time.RFC3339Nano), nil},
		{55.14, "14.22", valid2.ErrGreaterValue},
		{someCustomType{}, "10", valid2.ErrInvalidType},
		{"ololo", "olol", valid2.ErrUnsupportedComparator},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := valid2.Max(reflect.ValueOf(tc.value), tc.param)

			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestLesser(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		value       interface{}
		param       string
		expectedErr error
	}{
		{10, "14", nil},
		{int8(42), "78", nil},
		{now, now.Add(1 * time.Minute).Format(time.RFC3339Nano), nil},
		{now, now.Format(time.RFC3339Nano), valid2.ErrGreaterValue},
		{int16(52), "52", valid2.ErrGreaterValue},
		{55.17, "14.22", valid2.ErrGreaterValue},
		{someCustomType{}, "10", valid2.ErrInvalidType},
		{"ololo", "olol", valid2.ErrUnsupportedComparator},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := valid2.Lesser(reflect.ValueOf(tc.value), tc.param)

			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}

func TestGreater(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		value       interface{}
		param       string
		expectedErr error
	}{
		{int(14), "10", nil},
		{int8(78), "42", nil},
		{now.Add(1 * time.Minute), now.Format(time.RFC3339Nano), nil},
		{now, now.Format(time.RFC3339Nano), valid2.ErrLesserValue},
		{int16(52), "52", valid2.ErrLesserValue},
		{14.42, "55.17", valid2.ErrLesserValue},
		{someCustomType{}, "10", valid2.ErrInvalidType},
		{"ololo", "olol", valid2.ErrUnsupportedComparator},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			err := valid2.Greater(reflect.ValueOf(tc.value), tc.param)

			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}
