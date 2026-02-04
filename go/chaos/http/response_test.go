package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var goResp = &http.Response{
	Status:     "200 OK",
	ProtoMajor: 1,
	ProtoMinor: 1,
	Header: map[string][]string{
		"Content-Encoding": {"gzip, deflate"},
		"Content-Language": {"en-us"},
		"Foo":              {"Bar", "Ping"},
	},
}

var resp = &Response{
	Version: &Version{Minor: 1, Major: 1},
	Status: &Status{
		Code:   200,
		Reason: "OK",
	},
	Header: NewHeaderFrom(map[string][]string{
		"Content-Encoding": {"gzip, deflate"},
		"Content-Language": {"en-us"},
		"Foo":              {"Bar", "Ping"},
	}),
}

func TestResponse_SyncFrom(t *testing.T) {
	res := &Response{}
	err := res.SyncFrom(goResp)

	assert.NoError(t, err)
	assert.Equal(t, resp.Status.Format(), res.Status.Format())
	assert.Equal(t, resp.Version.Major, res.Version.Major)
	assert.Equal(t, resp.Version.Minor, res.Version.Minor)
	assert.Equal(t, resp.Header.Values("Content-Encoding"), res.Header.Values("Content-Encoding"))
	assert.Equal(t, resp.Header.Values("Content-Language"), res.Header.Values("Content-Language"))
	assert.Equal(t, resp.Header.Values("Foo"), res.Header.Values("Foo"))
}
