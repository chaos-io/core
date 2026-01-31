package core

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestObjectCodec_Decode(t *testing.T) {
	val := "{\"k\":\"v\"}"
	obj := &Object{}
	err := jsoniter.Unmarshal([]byte(val), obj)
	assert.NoError(t, err)
	assert.Equal(t, "v", obj.GetVals())
}

func TestObjectCodec_Encode(t *testing.T) {
	type args struct {
		ptr    unsafe.Pointer
		stream *jsoniter.Stream
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := &ObjectCodec{}
			codec.Encode(tt.args.ptr, tt.args.stream)
		})
	}
}

func TestObjectCodec_IsEmpty(t *testing.T) {
	type args struct {
		ptr unsafe.Pointer
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := &ObjectCodec{}
			assert.Equalf(t, tt.want, codec.IsEmpty(tt.args.ptr), "IsEmpty(%v)", tt.args.ptr)
		})
	}
}
