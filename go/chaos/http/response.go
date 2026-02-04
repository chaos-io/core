package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/chaos-io/core/go/chaos/core"
)

func NewResponseFrom(response *http.Response) (*Response, error) {
	resp := &Response{}
	err := resp.SyncFrom(response)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (x *Response) SyncFrom(response *http.Response) error {
	if x == nil {
		return fmt.Errorf("nil response")
	}

	if response == nil {
		return fmt.Errorf("nil http.Response")
	}

	if response.Body != nil {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		defer func() {
			_ = response.Body.Close()
		}()
		x.Body = core.NewBytesValue(body)
	} else {
		x.Body = nil
	}

	x.Version = &Version{Major: int32(response.ProtoMajor), Minor: int32(response.ProtoMinor)}

	x.Status = NewStatusFrom(response.StatusCode, response.Status)

	if len(response.Header) > 0 {
		x.Header = NewHeaderFrom(response.Header)
	} else {
		x.Header = nil
	}

	return nil
}
