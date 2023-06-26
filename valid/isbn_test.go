package valid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
)

func TestISBN(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrStringTooShort},
		{"foo", valid2.ErrStringTooShort},
		{"978-4-87311-368-5890", valid2.ErrInvalidISBN},
		{"3836221193", valid2.ErrInvalidISBN},

		{"3836221195", nil},
		{"1-61729-085-8", nil},
		{"3 423 21412 0", nil},
		{"3 401 01319 X", nil},
		{"9784873113685", nil},
		{"978-4-87311-368-5", nil},
		{"978 3401013190", nil},
		{"978-3-8362-2119-1", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.ISBN(tc.param))
		})
	}
}

func TestISBN10(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrStringTooShort},
		{"foo", valid2.ErrStringTooShort},
		{"342321412100", valid2.ErrInvalidStringLength},
		{"3836221191", valid2.ErrInvalidChecksum},

		{"3836221195", nil},
		{"1-61729-085-8", nil},
		{"3 423 21412 0", nil},
		{"3 401 01319 X", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.ISBN10(tc.param))
		})
	}
}

func TestISBN13(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrStringTooShort},
		{"foo", valid2.ErrStringTooShort},
		{"3-8362-2119-5", valid2.ErrInvalidStringLength},
		{"01234567890ab", valid2.ErrInvalidChecksum},
		{"978 3 8362 2119 0", valid2.ErrInvalidChecksum},

		{"9784873113685", nil},
		{"978-4-87311-368-5", nil},
		{"978 3401013190", nil},
		{"978-3-8362-2119-1", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.ISBN13(tc.param))
		})
	}
}

func BenchmarkISBN10(b *testing.B) {
	benchCases := []string{
		"3423214121",
		"978-3836221191",
		"3-423-21412-1",
		"3 423 21412 1",
		"3836221195",
		"1-61729-085-8",
		"3 423 21412 0",
		"3 401 01319 X",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid2.ISBN10(benchCases[i%len(benchCases)])
	}
}
