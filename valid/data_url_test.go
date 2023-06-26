package valid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
)

func TestDataURI(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrEmptyString},
		{"data:", valid2.ErrStringTooShort},
		{"img:iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrInvalidPrefix},
		{"data:iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrTooFewDataParts},
		{"data:image/png,iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrTooFewDataParts},
		{"data:image/png,base64,iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrTooFewDataParts},
		{"data:image/png;iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrInvalidCharsSequence},

		{"data:,", nil},
		{"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==", nil},
		{"data:text/plain;charset=UTF-8;page=21,the%20data:1234,5678", nil},
		{"data:text/vnd-example+xyz;foo=bar;base64,R0lGODdh", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.DataURL(tc.param))
		})
	}
}

func BenchmarkDataURI(b *testing.B) {
	benchCases := []string{
		"",
		"data:",
		"img:iVBORw0KGgoAAAANSUhEUgAA==",
		"data:iVBORw0KGgoAAAANSUhEUgAA==",
		"data:image/png,iVBORw0KGgoAAAANSUhEUgAA==",
		"data:image/png,base64,iVBORw0KGgoAAAANSUhEUgAA==",
		"data:image/png;iVBORw0KGgoAAAANSUhEUgAA==",
		"data:image/png;;base64,iVBORw0KGgoAAAANSUhEUgAA==",
		"data:,",
		"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==",
		"data:text/plain;charset=UTF-8;page=21,the%20data:1234,5678",
		"data:text/vnd-example+xyz;foo=bar;base64,R0lGODdh",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid2.DataURL(benchCases[i%len(benchCases)])
	}
}

func TestBase64DataURI(t *testing.T) {
	testCases := []struct {
		param     string
		expectErr error
	}{
		{"", valid2.ErrEmptyString},
		{"data:", valid2.ErrStringTooShort},
		{"data:,", valid2.ErrStringTooShort},
		{"img:iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrInvalidPrefix},
		{"data:iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrTooFewDataParts},
		{"data:image/png,iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrTooFewDataParts},
		{"data:image/png,base64,iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrTooFewDataParts},
		{"data:image/png;iVBORw0KGgoAAAANSUhEUgAA==", valid2.ErrInvalidCharsSequence},
		{"data:text/plain;charset=UTF-8;page=21,the%20data:1234,5678", valid2.ErrInvalidCharsSequence},
		{"data:image/png;base64,iVBORw0KGgoAAAANHxglj-&==", valid2.ErrInvalidCharacters},

		{"data:;base64,", nil},
		{"data:text/plain;charset=UTF-8;base64,iVBORw0KGgoAAAANSUhEUgAA==", nil},
		{"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			assert.Equal(t, tc.expectErr, valid2.Base64DataURL(tc.param))
		})
	}
}

func BenchmarkBase64DataURI(b *testing.B) {
	benchCases := []string{
		"",
		"data:",
		"data:,",
		"img:iVBORw0KGgoAAAANSUhEUgAA==",
		"data:iVBORw0KGgoAAAANSUhEUgAA==",
		"data:image/png",
		"data:image/png;iVBORw0KGgoAAAANSUhEUgAA==",
		"data:text/plain;charset=UTF-8;page=21",
		"data:;base64",
		"data:text/plain;charset=UTF-8;base64",
		"data:image/png;base64",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid2.Base64DataURL(benchCases[i%len(benchCases)])
	}
}
