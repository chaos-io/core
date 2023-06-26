package valid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
)

func TestLuhn(t *testing.T) {
	var testCases = []struct {
		value     string
		expectErr error
	}{
		{"1111111111", valid2.ErrInvalidChecksum},
		{"7992739871", valid2.ErrInvalidChecksum},
		{"4222222222222222", valid2.ErrInvalidChecksum},
		{"49927398717", valid2.ErrInvalidChecksum},
		{"1234567812345678", valid2.ErrInvalidChecksum},

		{"4276380091945522", nil},
		{"356938035643809", nil},
		{"49927398716", nil},
		{"1111111116", nil},
		{"12345674", nil},
		{"5515805738324655", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.value, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.Luhn(tc.value))
		})
	}
}

func BenchmarkLuhn(b *testing.B) {
	benchCases := []string{
		"1111111111",
		"7992739871",
		"4222222222222222",
		"49927398717",
		"1234567812345678",
		"4276380091945522",
		"356938035643809",
		"49927398716",
		"1111111116",
		"12345674",
		"5515805738324655",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid2.Luhn(benchCases[i%len(benchCases)])
	}
}
