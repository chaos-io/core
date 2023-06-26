package headers_test

import (
	"testing"

	headers2 "github.com/chaos-io/core/httputil/headers"
	// "github.com/google/go-cmp/cmp"
	// "github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// examples for tests taken from https://tools.ietf.org/html/rfc2616#section-14.3
func TestParseAcceptEncoding(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    headers2.AcceptableEncodings
		expectedErr error
	}{
		{
			"ietf_example_1",
			"compress, gzip",
			headers2.AcceptableEncodings{
				{Encoding: headers2.ContentEncoding("compress"), Weight: 1.0},
				{Encoding: headers2.ContentEncoding("gzip"), Weight: 1.0},
			},
			nil,
		},
		{
			"ietf_example_2",
			"",
			nil,
			nil,
		},
		{
			"ietf_example_3",
			"*",
			headers2.AcceptableEncodings{
				{Encoding: headers2.ContentEncoding("*"), Weight: 1.0},
			},
			nil,
		},
		{
			"ietf_example_4",
			"compress;q=0.5, gzip;q=1.0",
			headers2.AcceptableEncodings{
				{Encoding: headers2.ContentEncoding("gzip"), Weight: 1.0},
				{Encoding: headers2.ContentEncoding("compress"), Weight: 0.5},
			},
			nil,
		},
		{
			"ietf_example_5",
			"gzip;q=1.0, identity; q=0.5, *;q=0",
			headers2.AcceptableEncodings{
				{Encoding: headers2.ContentEncoding("gzip"), Weight: 1.0},
				{Encoding: headers2.ContentEncoding("identity"), Weight: 0.5},
				{Encoding: headers2.ContentEncoding("*"), Weight: 0},
			},
			nil,
		},
		{
			"solomon_headers",
			"zstd,lz4,gzip,deflate",
			headers2.AcceptableEncodings{
				{Encoding: headers2.ContentEncoding("zstd"), Weight: 1.0},
				{Encoding: headers2.ContentEncoding("lz4"), Weight: 1.0},
				{Encoding: headers2.ContentEncoding("gzip"), Weight: 1.0},
				{Encoding: headers2.ContentEncoding("deflate"), Weight: 1.0},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			acceptableEncodings, err := headers2.ParseAcceptEncoding(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			require.Len(t, acceptableEncodings, len(tc.expected))

			// ignore pos field
			// opt := cmpopts.IgnoreUnexported(headers.AcceptableEncoding{})
			// assert.True(t, cmp.Equal(tc.expected, acceptableEncodings, opt), cmp.Diff(tc.expected, acceptableEncodings, opt))

			var ae headers2.AcceptableEncodings
			for _, e := range acceptableEncodings {
				ae = append(ae, headers2.AcceptableEncoding{
					Encoding: e.Encoding,
					Weight:   e.Weight,
				})
			}
			assert.Equal(t, tc.expected, ae)
		})
	}
}

func TestParseAccept(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    headers2.AcceptableTypes
		expectedErr error
	}{
		{
			"empty_header",
			"",
			nil,
			nil,
		},
		{
			"accept_any",
			"*/*",
			headers2.AcceptableTypes{
				{Type: headers2.ContentTypeAny, Weight: 1.0},
			},
			nil,
		},
		{
			"accept_single",
			"application/json",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationJSON, Weight: 1.0},
			},
			nil,
		},
		{
			"accept_multiple",
			"application/json, application/protobuf",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationJSON, Weight: 1.0},
				{Type: headers2.TypeApplicationProtobuf, Weight: 1.0},
			},
			nil,
		},
		{
			"accept_multiple_weighted",
			"application/json;q=0.8, application/protobuf",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationProtobuf, Weight: 1.0},
				{Type: headers2.TypeApplicationJSON, Weight: 0.8},
			},
			nil,
		},
		{
			"accept_multiple_weighted_unsorted",
			"text/plain;q=0.5, application/protobuf, application/json;q=0.5",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationProtobuf, Weight: 1.0},
				{Type: headers2.TypeTextPlain, Weight: 0.5},
				{Type: headers2.TypeApplicationJSON, Weight: 0.5},
			},
			nil,
		},
		{
			"unknown_type",
			"custom/type, unknown/my_type;q=0.2",
			headers2.AcceptableTypes{
				{Type: headers2.ContentType("custom/type"), Weight: 1.0},
				{Type: headers2.ContentType("unknown/my_type"), Weight: 0.2},
			},
			nil,
		},
		{
			"yabro_19.6.0",
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			headers2.AcceptableTypes{
				{Type: headers2.ContentType("text/html"), Weight: 1.0},
				{Type: headers2.ContentType("application/xhtml+xml"), Weight: 1.0},
				{Type: headers2.ContentType("image/webp"), Weight: 1.0},
				{Type: headers2.ContentType("image/apng"), Weight: 1.0},
				{Type: headers2.ContentType("application/signed-exchange"), Weight: 1.0, Extension: map[string]string{"v": "b3"}},
				{Type: headers2.ContentType("application/xml"), Weight: 0.9},
				{Type: headers2.ContentType("*/*"), Weight: 0.8},
			},
			nil,
		},
		{
			"chrome_81.0.4044",
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			headers2.AcceptableTypes{
				{Type: headers2.ContentType("text/html"), Weight: 1.0},
				{Type: headers2.ContentType("application/xhtml+xml"), Weight: 1.0},
				{Type: headers2.ContentType("image/webp"), Weight: 1.0},
				{Type: headers2.ContentType("image/apng"), Weight: 1.0},
				{Type: headers2.ContentType("application/xml"), Weight: 0.9},
				{Type: headers2.ContentType("application/signed-exchange"), Weight: 0.9, Extension: map[string]string{"v": "b3"}},
				{Type: headers2.ContentType("*/*"), Weight: 0.8},
			},
			nil,
		},
		{
			"firefox_77.0b3",
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
			headers2.AcceptableTypes{
				{Type: headers2.ContentType("text/html"), Weight: 1.0},
				{Type: headers2.ContentType("application/xhtml+xml"), Weight: 1.0},
				{Type: headers2.ContentType("image/webp"), Weight: 1.0},
				{Type: headers2.ContentType("application/xml"), Weight: 0.9},
				{Type: headers2.ContentType("*/*"), Weight: 0.8},
			},
			nil,
		},
		{
			"sort_by_most_specific",
			"text/*, text/html, */*, text/html;level=1",
			headers2.AcceptableTypes{
				{Type: headers2.ContentType("text/html"), Weight: 1.0, Extension: map[string]string{"level": "1"}},
				{Type: headers2.ContentType("text/html"), Weight: 1.0},
				{Type: headers2.ContentType("text/*"), Weight: 1.0},
				{Type: headers2.ContentType("*/*"), Weight: 1.0},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			at, err := headers2.ParseAccept(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			require.Len(t, at, len(tc.expected))

			// opt := cmpopts.IgnoreUnexported(headers.AcceptableType{})
			// assert.True(t, cmp.Equal(tc.expected, at, opt), cmp.Diff(tc.expected, at, opt))
			var ats headers2.AcceptableTypes
			for _, a := range at {
				ats = append(ats, headers2.AcceptableType{
					Type:      a.Type,
					Weight:    a.Weight,
					Extension: a.Extension,
				})
			}
			assert.Equal(t, tc.expected, ats)
		})
	}
}

func TestAcceptableTypesString(t *testing.T) {
	testCases := []struct {
		name     string
		types    headers2.AcceptableTypes
		expected string
	}{
		{
			"empty",
			headers2.AcceptableTypes{},
			"",
		},
		{
			"single",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationJSON},
			},
			"application/json",
		},
		{
			"single_weighted",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationJSON, Weight: 0.8},
			},
			"application/json;q=0.8",
		},
		{
			"multiple",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationJSON},
				{Type: headers2.TypeApplicationProtobuf},
			},
			"application/json, application/protobuf",
		},
		{
			"multiple_weighted",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationProtobuf},
				{Type: headers2.TypeApplicationJSON, Weight: 0.8},
			},
			"application/protobuf, application/json;q=0.8",
		},
		{
			"multiple_weighted_with_extension",
			headers2.AcceptableTypes{
				{Type: headers2.TypeApplicationProtobuf},
				{Type: headers2.TypeApplicationJSON, Weight: 0.8},
				{Type: headers2.TypeApplicationXML, Weight: 0.5, Extension: map[string]string{"label": "1"}},
			},
			"application/protobuf, application/json;q=0.8, application/xml;q=0.5;label=1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.types.String())
		})
	}
}

func BenchmarkParseAccept(b *testing.B) {
	benchCases := []string{
		"",
		"*/*",
		"application/json",
		"application/json, application/protobuf",
		"application/json;q=0.8, application/protobuf",
		"text/plain;q=0.5, application/protobuf, application/json;q=0.5",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"text/*, text/html, */*, text/html;level=1",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = headers2.ParseAccept(benchCases[i%len(benchCases)])
	}
}
