package rule

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	inspection2 "github.com/chaos-io/core/valid/v2/inspection"
)

func TestMustMatch(t *testing.T) {
	testCases := []struct {
		name     string
		regex    *regexp.Regexp
		value    *inspection2.Inspected
		expected error
	}{
		{
			name:     "match",
			regex:    regexp.MustCompile("(shi|boo)mba"),
			value:    inspection2.Inspect("looken boomba tooken"),
			expected: nil,
		},
		{
			name:     "not_match",
			regex:    regexp.MustCompile("(shi|boo)mba"),
			value:    inspection2.Inspect("looken tooken chiken cooken"),
			expected: ErrPatternMismatch,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := MustMatch(tc.regex)
			assert.Equal(t, tc.expected, r(tc.value))
		})
	}
}
