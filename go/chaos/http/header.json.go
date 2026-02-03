package http

import (
	"unsafe"

	"github.com/chaos-io/core/go/chaos/core"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	core.RegisterJSONTypeDecoder(HeaderTypeFullName, &HeaderCodec{})
	core.RegisterJSONTypeEncoder(HeaderTypeFullName, &HeaderCodec{})
}

type HeaderCodec struct{}

func (codec *HeaderCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	header := (*Header)(ptr)
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		a.ToVal(&header.Vals)
	}
}

func (codec *HeaderCodec) IsEmpty(ptr unsafe.Pointer) bool {
	header := (*Header)(ptr)
	return header == nil || len(header.Vals) == 0
}

func (codec *HeaderCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	header := (*Header)(ptr)
	stream.WriteVal(header.Vals)
}
