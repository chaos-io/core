package http

import (
	"testing"

	"github.com/chaos-io/core/go/chaos/core"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestHeaderCodec_Decode(t *testing.T) {
	s := `{"Content-Type":["application/json"]}`
	header := &Header{}
	err := jsoniter.UnmarshalFromString(s, header)
	assert.NoError(t, err)
	assert.EqualValues(t, &core.StringValues{Vals: []string{"application/json"}}, header.Vals["Content-Type"])
}

func TestHeaderCodec_Encode(t *testing.T) {
	header := &Header{Vals: map[string]*core.StringValues{"Content-Type": {Vals: []string{"application/json"}}}}
	str, err := jsoniter.MarshalToString(header)
	assert.NoError(t, err)
	assert.Equal(t, `{"Content-Type":["application/json"]}`, str)
}

func TestHeaderCodec_IsEmpty(t *testing.T) {
	header := &Header{}
	str, err := jsoniter.MarshalToString(header)
	assert.NoError(t, err)
	assert.EqualValues(t, "null", str)
}
