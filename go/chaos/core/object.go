package core

import (
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const ObjectTypeName = "Object"
const ObjectTypeFullName = "core.Object"

// NewObject constructs a Struct from a general-purpose Go map.
// The map keys must be valid UTF-8.
// The map values are converted using NewValue.
func NewObject(v map[string]any) (*Object, error) {
	x := &Object{Fields: make(map[string]*Value, len(v))}
	for k, v := range v {
		if !utf8.ValidString(k) {
			return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", k)
		}
		var err error
		x.Fields[k], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

// AsMap converts x to a general-purpose Go map.
// The map values are converted by calling Value.AsInterface.
func (x *Object) AsMap() map[string]any {
	f := x.GetFields()
	vs := make(map[string]any, len(f))
	for k, v := range f {
		vs[k] = v.AsInterface()
	}
	return vs
}

func (x *Object) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Object) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}
