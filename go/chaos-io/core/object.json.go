package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder(ObjectTypeFullName, &ObjectCodec{})
	RegisterJSONTypeEncoder(ObjectTypeFullName, &ObjectCodec{})
}

type ObjectCodec struct{}

func (codec *ObjectCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	obj := (*Object)(ptr)
	if a.ValueType() == jsoniter.ObjectValue && a.Size() > 0 {
		obj.Fields = make(map[string]*Value)
		for _, k := range a.Keys() {
			val, err := NewValueCodec().DecodeAny(a.Get(k))
			if err != nil {
				iter.ReportError("ObjectCodec.Decode", err.Error())
			}
			obj.Fields[k] = val
		}
	}
}

func (codec *ObjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
	obj := (*Object)(ptr)
	return obj == nil || len(obj.Fields) == 0
}

func (codec *ObjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*Object)(ptr)
	stream.WriteVal(obj.Fields)
}
