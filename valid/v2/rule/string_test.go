package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inspection2 "github.com/chaos-io/core/valid/v2/inspection"
)

func TestIsAlphanumeric(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", ErrEmptyString},
		{"\n", ErrInvalidCharacters},
		{"\r", ErrInvalidCharacters},
		{"Ⅸ", ErrInvalidCharacters},
		{"   fooo   ", ErrInvalidCharacters},
		{"abc!!!", ErrInvalidCharacters},
		{"abc〩", ErrInvalidCharacters},
		{"소주", ErrInvalidCharacters},
		{"소aBC", ErrInvalidCharacters},
		{"소", ErrInvalidCharacters},
		{"달기&Co.", ErrInvalidCharacters},
		{"〩Hours", ErrInvalidCharacters},
		{"\ufff0", ErrInvalidCharacters},

		{"\u0026", ErrInvalidCharacters}, // UTF-8(ASCII): &
		{"-00123", ErrInvalidCharacters},
		{"-0", ErrInvalidCharacters},
		{"123.123", ErrInvalidCharacters},
		{" ", ErrInvalidCharacters},
		{".", ErrInvalidCharacters},
		{"-1¾", ErrInvalidCharacters},
		{"1¾", ErrInvalidCharacters},
		{"〥〩", ErrInvalidCharacters},
		{"모자", ErrInvalidCharacters},
		{"۳۵۶۰", ErrInvalidCharacters},
		{"1--", ErrInvalidCharacters},
		{"1-1", ErrInvalidCharacters},
		{"-", ErrInvalidCharacters},
		{"--", ErrInvalidCharacters},
		{"1++", ErrInvalidCharacters},
		{"1+1", ErrInvalidCharacters},
		{"+", ErrInvalidCharacters},
		{"++", ErrInvalidCharacters},
		{"+1", ErrInvalidCharacters},

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
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, IsAlphanumeric(v))
		})
	}
}

func TestIsAlpha(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", ErrEmptyString},
		{"\n", ErrInvalidCharacters},
		{"\r", ErrInvalidCharacters},
		{"Ⅸ", ErrInvalidCharacters},
		{"   fooo   ", ErrInvalidCharacters},
		{"abc!!!", ErrInvalidCharacters},
		{"abc1", ErrInvalidCharacters},
		{"abc〩", ErrInvalidCharacters},
		{"소주", ErrInvalidCharacters},
		{"소aBC", ErrInvalidCharacters},
		{"소", ErrInvalidCharacters},
		{"달기&Co.", ErrInvalidCharacters},
		{"〩Hours", ErrInvalidCharacters},
		{"\ufff0", ErrInvalidCharacters},
		{"\u0026", ErrInvalidCharacters}, // UTF-8(ASCII): &
		{"\u0030", ErrInvalidCharacters}, // UTF-8(ASCII): 0
		{"123", ErrInvalidCharacters},
		{"0123", ErrInvalidCharacters},
		{"-00123", ErrInvalidCharacters},
		{"0", ErrInvalidCharacters},
		{"-0", ErrInvalidCharacters},
		{"123.123", ErrInvalidCharacters},
		{" ", ErrInvalidCharacters},
		{".", ErrInvalidCharacters},
		{"-1¾", ErrInvalidCharacters},
		{"1¾", ErrInvalidCharacters},
		{"〥〩", ErrInvalidCharacters},
		{"모자", ErrInvalidCharacters},
		{"۳۵۶۰", ErrInvalidCharacters},
		{"1--", ErrInvalidCharacters},
		{"1-1", ErrInvalidCharacters},
		{"-", ErrInvalidCharacters},
		{"--", ErrInvalidCharacters},
		{"1++", ErrInvalidCharacters},
		{"1+1", ErrInvalidCharacters},
		{"+", ErrInvalidCharacters},
		{"++", ErrInvalidCharacters},
		{"+1", ErrInvalidCharacters},

		{"ix", nil},
		{"\u0070", nil}, // UTF-8(ASCII): p
		{"ABC", nil},
		{"FoObAr", nil},
		{"abc", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, IsAlpha(v))
		})
	}
}

func TestIsNumeric(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", ErrEmptyString},
		{"\n", ErrInvalidCharacters},
		{"\r", ErrInvalidCharacters},
		{"Ⅸ", ErrInvalidCharacters},
		{"   fooo   ", ErrInvalidCharacters},
		{"abc!!!", ErrInvalidCharacters},
		{"abc1", ErrInvalidCharacters},
		{"abc〩", ErrInvalidCharacters},
		{"abc", ErrInvalidCharacters},
		{"소주", ErrInvalidCharacters},
		{"ABC", ErrInvalidCharacters},
		{"FoObAr", ErrInvalidCharacters},
		{"소aBC", ErrInvalidCharacters},
		{"소", ErrInvalidCharacters},
		{"달기&Co.", ErrInvalidCharacters},
		{"〩Hours", ErrInvalidCharacters},
		{"\ufff0", ErrInvalidCharacters},
		{"\u0070", ErrInvalidCharacters}, // UTF-8(ASCII): p
		{"\u0026", ErrInvalidCharacters}, // UTF-8(ASCII): &
		{"\u0030", nil},                  // UTF-8(ASCII): 0
		{"-00123", ErrInvalidCharacters},
		{"+00123", ErrInvalidCharacters},
		{"-0", ErrInvalidCharacters},
		{"123.123", ErrInvalidCharacters},
		{" ", ErrInvalidCharacters},
		{".", ErrInvalidCharacters},
		{"12𐅪3", ErrInvalidCharacters},
		{"-1¾", ErrInvalidCharacters},
		{"1¾", ErrInvalidCharacters},
		{"〥〩", ErrInvalidCharacters},
		{"모자", ErrInvalidCharacters},
		{"ix", ErrInvalidCharacters},
		{"۳۵۶۰", ErrInvalidCharacters},
		{"1--", ErrInvalidCharacters},
		{"1-1", ErrInvalidCharacters},
		{"-", ErrInvalidCharacters},
		{"--", ErrInvalidCharacters},
		{"1++", ErrInvalidCharacters},
		{"1+1", ErrInvalidCharacters},
		{"+", ErrInvalidCharacters},
		{"++", ErrInvalidCharacters},
		{"+1", ErrInvalidCharacters},

		{"0", nil},
		{"123", nil},
		{"0123", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, IsNumeric(v))
		})
	}
}

func TestIsASCII(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", ErrEmptyString},
		{"\n", ErrInvalidCharacters},
		{"\r", ErrInvalidCharacters},
		{"Ⅸ", ErrInvalidCharacters},
		{"   fooo   ", nil},
		{"abc!!!", nil},
		{"abc1", nil},
		{"abc〩", ErrInvalidCharacters},
		{"소주", ErrInvalidCharacters},
		{"소aBC", ErrInvalidCharacters},
		{"소", ErrInvalidCharacters},
		{"달기&Co.", ErrInvalidCharacters},
		{"〩Hours", ErrInvalidCharacters},
		{"\ufff0", ErrInvalidCharacters},
		{"\u0026", nil}, // UTF-8(ASCII): &
		{"\u0030", nil}, // UTF-8(ASCII): 0
		{"123", nil},
		{"0123", nil},
		{"-00123", nil},
		{"0", nil},
		{"-0", nil},
		{"123.123", nil},
		{" ", nil},
		{".", nil},
		{"-1¾", ErrInvalidCharacters},
		{"1¾", ErrInvalidCharacters},
		{"〥〩", ErrInvalidCharacters},
		{"모자", ErrInvalidCharacters},
		{"۳۵۶۰", ErrInvalidCharacters},
		{"1--", nil},
		{"1-1", nil},
		{"-", nil},
		{"--", nil},
		{"1++", nil},
		{"1+1", nil},
		{"+", nil},
		{"++", nil},
		{"+1", nil},

		{"ix", nil},
		{"\u0070", nil}, // UTF-8(ASCII): p
		{"ABC", nil},
		{"FoObAr", nil},
		{"abc", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, IsASCII(v))
		})
	}
}

func TestHexColor(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", ErrStringTooShort},
		{"#ff", ErrInvalidStringLength},
		{"fff0", ErrInvalidStringLength},
		{"#ff12FG", ErrInvalidCharacters},

		{"CCccCC", nil},
		{"fff", nil},
		{"#f00", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, IsHexColor(v))
		})
	}
}

func TestHasPrefix(t *testing.T) {
	testCases := []struct {
		prefix    string
		param     string
		expectErr error
	}{
		{"test", "", ErrPatternMismatch},
		{"0", "fff0", ErrPatternMismatch},

		{"", "#ff", nil},
		{"CC", "CCccCC", nil},
		{"f", "fff", nil},
		{"#f00", "#f00", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, HasPrefix(tc.prefix)(v))
		})
	}
}

func TestHasSuffix(t *testing.T) {
	testCases := []struct {
		suffix    string
		param     string
		expectErr error
	}{
		{"test", "", ErrPatternMismatch},
		{"f", "fff0", ErrPatternMismatch},

		{"", "#ff", nil},
		{"CC", "CCccCC", nil},
		{"f", "fff", nil},
		{"#f00", "#f00", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, HasSuffix(tc.suffix)(v))
		})
	}
}

func TestIs2DMeasurements(t *testing.T) {
	testCases := []struct {
		separator string
		param     string
		expectErr error
	}{
		{"x", "200*300", ErrPatternMismatch},
		{"x", "200300", ErrPatternMismatch},
		{"x", "200*", ErrPatternMismatch},
		{"x", "*300", ErrPatternMismatch},
		{"x", "SHIMBAxBOOMBA", ErrPatternMismatch},
		{"x", "SHIMBAx400", ErrPatternMismatch},

		{"x", "200x300", nil},
		{"*", "200*300", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			v := inspection2.Inspect(tc.param)
			assert.Equal(t, tc.expectErr, Is2DMeasurements(tc.separator)(v))
		})
	}
}

func BenchmarkAlnum(b *testing.B) {
	testCases := []*inspection2.Inspected{
		inspection2.Inspect(""),
		inspection2.Inspect("\n"),
		inspection2.Inspect("\r"),
		inspection2.Inspect("Ⅸ"),
		inspection2.Inspect("   fooo   "),
		inspection2.Inspect("abc!!!"),
		inspection2.Inspect("abc〩"),
		inspection2.Inspect("소주"),
		inspection2.Inspect("소aBC"),
		inspection2.Inspect("소"),
		inspection2.Inspect("달기&Co."),
		inspection2.Inspect("〩Hours"),
		inspection2.Inspect("\ufff0"),
		inspection2.Inspect("\u0026"),
		inspection2.Inspect("-00123"),
		inspection2.Inspect("-0"),
		inspection2.Inspect("123.123"),
		inspection2.Inspect(" "),
		inspection2.Inspect("."),
		inspection2.Inspect("-1¾"),
		inspection2.Inspect("1¾"),
		inspection2.Inspect("〥〩"),
		inspection2.Inspect("모자"),
		inspection2.Inspect("۳۵۶۰"),
		inspection2.Inspect("1--"),
		inspection2.Inspect("1-1"),
		inspection2.Inspect("-"),
		inspection2.Inspect("--"),
		inspection2.Inspect("1++"),
		inspection2.Inspect("1+1"),
		inspection2.Inspect("+"),
		inspection2.Inspect("++"),
		inspection2.Inspect("+1"),
		inspection2.Inspect("abc"),
		inspection2.Inspect("abc123"),
		inspection2.Inspect("ABC111"),
		inspection2.Inspect("abc1"),
		inspection2.Inspect("ABC"),
		inspection2.Inspect("FoObAr"),
		inspection2.Inspect("ix"),
		inspection2.Inspect("0"),
		inspection2.Inspect("\u0030"),
		inspection2.Inspect("123"),
		inspection2.Inspect("0123"),
		inspection2.Inspect("\u0070"),
	}

	rules := map[string]Rule{
		"IsAlphanumeric": IsAlphanumeric,
		"IsAlpha":        IsAlpha,
		"IsNumeric":      IsNumeric,
		"IsASCII":        IsASCII,
	}

	for name, r := range rules {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = r(testCases[i%len(testCases)])
			}
		})
	}
}

func BenchmarkIsHexColor(b *testing.B) {
	testCases := []*inspection2.Inspected{
		inspection2.Inspect(""),
		inspection2.Inspect("#ff"),
		inspection2.Inspect("fff0"),
		inspection2.Inspect("#ff12FG"),
		inspection2.Inspect("CCcCCc"),
		inspection2.Inspect("fff"),
		inspection2.Inspect("#f00"),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsHexColor(testCases[i%len(testCases)])
	}
}

func BenchmarkHasPrefix(b *testing.B) {
	testCases := []*inspection2.Inspected{
		inspection2.Inspect(""),
		inspection2.Inspect("#ff"),
		inspection2.Inspect("fff0"),
		inspection2.Inspect("#ff12FG"),
		inspection2.Inspect("CCcCCc"),
		inspection2.Inspect("fff"),
		inspection2.Inspect("#f00"),
	}

	r := HasPrefix("f")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r(testCases[i%len(testCases)])
	}
}

func BenchmarkHasSuffix(b *testing.B) {
	testCases := []*inspection2.Inspected{
		inspection2.Inspect(""),
		inspection2.Inspect("#ff"),
		inspection2.Inspect("fff0"),
		inspection2.Inspect("#ff12FG"),
		inspection2.Inspect("CCcCCc"),
		inspection2.Inspect("fff"),
		inspection2.Inspect("#f00"),
	}

	r := HasSuffix("0")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r(testCases[i%len(testCases)])
	}
}

func BenchmarkIs2DMeasurements(b *testing.B) {
	testCases := []*inspection2.Inspected{
		inspection2.Inspect("200x200"),
		inspection2.Inspect("x200"),
		inspection2.Inspect("300x"),
		inspection2.Inspect("424242"),
		inspection2.Inspect("300*200"),
		inspection2.Inspect("SHIMBAxBOOMBA"),
		inspection2.Inspect("400xBOOMBA"),
	}

	r := Is2DMeasurements("x")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r(testCases[i%len(testCases)])
	}
}
