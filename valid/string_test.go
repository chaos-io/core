package valid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
)

func TestAlphanumeric(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrEmptyString},
		{"\n", valid2.ErrInvalidCharacters},
		{"\r", valid2.ErrInvalidCharacters},
		{"Ⅸ", valid2.ErrInvalidCharacters},
		{"   fooo   ", valid2.ErrInvalidCharacters},
		{"abc!!!", valid2.ErrInvalidCharacters},
		{"abc〩", valid2.ErrInvalidCharacters},
		{"소주", valid2.ErrInvalidCharacters},
		{"소aBC", valid2.ErrInvalidCharacters},
		{"소", valid2.ErrInvalidCharacters},
		{"달기&Co.", valid2.ErrInvalidCharacters},
		{"〩Hours", valid2.ErrInvalidCharacters},
		{"\ufff0", valid2.ErrInvalidCharacters},

		{"\u0026", valid2.ErrInvalidCharacters}, // UTF-8(ASCII): &
		{"-00123", valid2.ErrInvalidCharacters},
		{"-0", valid2.ErrInvalidCharacters},
		{"123.123", valid2.ErrInvalidCharacters},
		{" ", valid2.ErrInvalidCharacters},
		{".", valid2.ErrInvalidCharacters},
		{"-1¾", valid2.ErrInvalidCharacters},
		{"1¾", valid2.ErrInvalidCharacters},
		{"〥〩", valid2.ErrInvalidCharacters},
		{"모자", valid2.ErrInvalidCharacters},
		{"۳۵۶۰", valid2.ErrInvalidCharacters},
		{"1--", valid2.ErrInvalidCharacters},
		{"1-1", valid2.ErrInvalidCharacters},
		{"-", valid2.ErrInvalidCharacters},
		{"--", valid2.ErrInvalidCharacters},
		{"1++", valid2.ErrInvalidCharacters},
		{"1+1", valid2.ErrInvalidCharacters},
		{"+", valid2.ErrInvalidCharacters},
		{"++", valid2.ErrInvalidCharacters},
		{"+1", valid2.ErrInvalidCharacters},

		{"abc", nil},
		{"abc123", nil},
		{"ABC111", nil},
		{"abc1", nil},
		{"ABC", nil},
		{"FoObAr", nil},
		{"ix", nil},
		{"0", nil},
		{"\u0030", nil}, // UTF-8(ASCII): 0
		{"123", nil},
		{"0123", nil},
		{"\u0070", nil}, // UTF-8(ASCII): p
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.Alphanumeric(tc.param))
		})
	}
}

func TestAlpha(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrEmptyString},
		{"\n", valid2.ErrInvalidCharacters},
		{"\r", valid2.ErrInvalidCharacters},
		{"Ⅸ", valid2.ErrInvalidCharacters},
		{"   fooo   ", valid2.ErrInvalidCharacters},
		{"abc!!!", valid2.ErrInvalidCharacters},
		{"abc1", valid2.ErrInvalidCharacters},
		{"abc〩", valid2.ErrInvalidCharacters},
		{"소주", valid2.ErrInvalidCharacters},
		{"소aBC", valid2.ErrInvalidCharacters},
		{"소", valid2.ErrInvalidCharacters},
		{"달기&Co.", valid2.ErrInvalidCharacters},
		{"〩Hours", valid2.ErrInvalidCharacters},
		{"\ufff0", valid2.ErrInvalidCharacters},
		{"\u0026", valid2.ErrInvalidCharacters}, // UTF-8(ASCII): &
		{"\u0030", valid2.ErrInvalidCharacters}, // UTF-8(ASCII): 0
		{"123", valid2.ErrInvalidCharacters},
		{"0123", valid2.ErrInvalidCharacters},
		{"-00123", valid2.ErrInvalidCharacters},
		{"0", valid2.ErrInvalidCharacters},
		{"-0", valid2.ErrInvalidCharacters},
		{"123.123", valid2.ErrInvalidCharacters},
		{" ", valid2.ErrInvalidCharacters},
		{".", valid2.ErrInvalidCharacters},
		{"-1¾", valid2.ErrInvalidCharacters},
		{"1¾", valid2.ErrInvalidCharacters},
		{"〥〩", valid2.ErrInvalidCharacters},
		{"모자", valid2.ErrInvalidCharacters},
		{"۳۵۶۰", valid2.ErrInvalidCharacters},
		{"1--", valid2.ErrInvalidCharacters},
		{"1-1", valid2.ErrInvalidCharacters},
		{"-", valid2.ErrInvalidCharacters},
		{"--", valid2.ErrInvalidCharacters},
		{"1++", valid2.ErrInvalidCharacters},
		{"1+1", valid2.ErrInvalidCharacters},
		{"+", valid2.ErrInvalidCharacters},
		{"++", valid2.ErrInvalidCharacters},
		{"+1", valid2.ErrInvalidCharacters},

		{"ix", nil},
		{"\u0070", nil}, // UTF-8(ASCII): p
		{"ABC", nil},
		{"FoObAr", nil},
		{"abc", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.Alpha(tc.param))
		})
	}
}

func TestNumeric(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrEmptyString},
		{"\n", valid2.ErrInvalidCharacters},
		{"\r", valid2.ErrInvalidCharacters},
		{"Ⅸ", valid2.ErrInvalidCharacters},
		{"   fooo   ", valid2.ErrInvalidCharacters},
		{"abc!!!", valid2.ErrInvalidCharacters},
		{"abc1", valid2.ErrInvalidCharacters},
		{"abc〩", valid2.ErrInvalidCharacters},
		{"abc", valid2.ErrInvalidCharacters},
		{"소주", valid2.ErrInvalidCharacters},
		{"ABC", valid2.ErrInvalidCharacters},
		{"FoObAr", valid2.ErrInvalidCharacters},
		{"소aBC", valid2.ErrInvalidCharacters},
		{"소", valid2.ErrInvalidCharacters},
		{"달기&Co.", valid2.ErrInvalidCharacters},
		{"〩Hours", valid2.ErrInvalidCharacters},
		{"\ufff0", valid2.ErrInvalidCharacters},
		{"\u0070", valid2.ErrInvalidCharacters}, // UTF-8(ASCII): p
		{"\u0026", valid2.ErrInvalidCharacters}, // UTF-8(ASCII): &
		{"\u0030", nil},                         // UTF-8(ASCII): 0
		{"-00123", valid2.ErrInvalidCharacters},
		{"+00123", valid2.ErrInvalidCharacters},
		{"-0", valid2.ErrInvalidCharacters},
		{"123.123", valid2.ErrInvalidCharacters},
		{" ", valid2.ErrInvalidCharacters},
		{".", valid2.ErrInvalidCharacters},
		{"12𐅪3", valid2.ErrInvalidCharacters},
		{"-1¾", valid2.ErrInvalidCharacters},
		{"1¾", valid2.ErrInvalidCharacters},
		{"〥〩", valid2.ErrInvalidCharacters},
		{"모자", valid2.ErrInvalidCharacters},
		{"ix", valid2.ErrInvalidCharacters},
		{"۳۵۶۰", valid2.ErrInvalidCharacters},
		{"1--", valid2.ErrInvalidCharacters},
		{"1-1", valid2.ErrInvalidCharacters},
		{"-", valid2.ErrInvalidCharacters},
		{"--", valid2.ErrInvalidCharacters},
		{"1++", valid2.ErrInvalidCharacters},
		{"1+1", valid2.ErrInvalidCharacters},
		{"+", valid2.ErrInvalidCharacters},
		{"++", valid2.ErrInvalidCharacters},
		{"+1", valid2.ErrInvalidCharacters},

		{"0", nil},
		{"123", nil},
		{"0123", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.Numeric(tc.param))
		})
	}
}

func TestDouble(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrEmptyString},
		{".", valid2.ErrBadFormat},
		{"ololo", valid2.ErrInvalidCharacters},
		{".0", valid2.ErrBadFormat},
		{"01234.0", valid2.ErrBadFormat},

		{"0", nil},
		{"1234", nil},
		{"0.0001", nil},
		{"1234.00", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.Double(tc.param))
		})
	}
}

func TestHexColor(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrStringTooShort},
		{"#ff", valid2.ErrInvalidStringLength},
		{"fff0", valid2.ErrInvalidStringLength},
		{"#ff12FG", valid2.ErrInvalidCharacters},

		{"CCccCC", nil},
		{"fff", nil},
		{"#f00", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.HexColor(tc.param))
		})
	}
}

func TestStringLen(t *testing.T) {
	testCases := []struct {
		str       string
		min       int
		max       int
		expectErr error
	}{
		{"anything", 5, 2, valid2.ErrBadParams},
		{"", 2, 5, valid2.ErrStringTooShort},
		{"a", 2, 5, valid2.ErrStringTooShort},
		{"ab", 2, 5, nil},
		{"abc", 2, 5, nil},
		{"abcd", 2, 5, nil},
		{"abcde", 2, 5, nil},
		{"abcdef", 2, 5, valid2.ErrStringTooLong},
		{"abcdefg", 2, 5, valid2.ErrStringTooLong},

		{"just_min", 9, 0, valid2.ErrStringTooShort},
		{"just_min", 5, 0, nil},

		{"just_max", 0, 5, valid2.ErrStringTooLong},
		{"just_min", 0, 10, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.StringLen(tc.str, tc.min, tc.max))
		})
	}
}
