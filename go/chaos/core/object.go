package core

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/chaos-io/core/go/chaos/core/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"google.golang.org/protobuf/encoding/protojson"
)

const ObjectTypeName = "Object"
const ObjectTypeFullName = "chaos.core.Object"

func NewObject() *Object {
	return &Object{Vals: make(map[string]*Value)}
}

// NewObjectFromMap constructs a Struct from a general-purpose Go map.
// The map keys must be valid UTF-8.
// The map values are converted using NewValue.
func NewObjectFromMap(m map[string]any) (*Object, error) {
	x := &Object{Vals: make(map[string]*Value, len(m))}
	for k, v := range m {
		if !utf8.ValidString(k) {
			return nil, fmt.Errorf("invalid UTF-8 in object key: %q", k)
		}
		var err error
		x.Vals[k], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

func NewObjectFromKeyVals(kvs ...any) (*Object, error) {
	if len(kvs)%2 != 0 {
		return nil, fmt.Errorf("invalid number of key/value pairs: %d", len(kvs))
	}

	m := make(map[string]any)
	for i := 0; i < len(kvs); i += 2 {
		k, ok := kvs[i].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key/value pair at index %d", i)
		}
		m[k] = kvs[i+1]
	}

	return NewObjectFromMap(m)
}

func NewObjectFrom(val any) (*Object, error) {
	obj := NewObject()
	return obj, obj.From(val)
}

func NewObjectFromValues(v map[string]*Value) *Object {
	return &Object{Vals: v}
}

func MergeObjects(objs ...*Object) *Object {
	obj := NewObject()
	for _, o := range objs {
		obj.Merge(o)
	}
	return obj
}

// AsMap converts Object into a map[string]any using Value.AsInterface.
// The result is intended for JSON serialization or dynamic inspection,
// not for round-trip binary fidelity.
func (x *Object) AsMap() map[string]any {
	if x == nil {
		return nil
	}

	f := x.GetVals()
	vs := make(map[string]any, len(f))
	for k, v := range f {
		if v != nil {
			vs[k] = v.AsInterface()
		} else {
			vs[k] = nil
		}
	}

	return vs
}

func (x *Object) To(val any) error {
	if x != nil {
		marshal, err := jsoniter.ConfigFastest.Marshal(x)
		if err != nil {
			return err
		}
		return jsoniter.ConfigFastest.Unmarshal(marshal, val)
	}
	return nil
}

func (x *Object) From(val any) error {
	if x == nil {
		return nil
	}
	if val == nil || reflect.ValueOf(val).IsZero() {
		return nil
	}

	if x.Vals == nil {
		x.Vals = make(map[string]*Value)
	}

	switch v := val.(type) {
	case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, []byte:
		return fmt.Errorf("invalid basic object format: %T", v)
	case []any, *Values:
		return fmt.Errorf("invalid array object format: %T", v)
	case *Value:
		if obj := v.GetObject(); obj != nil {
			x.Vals = obj.Vals
		} else {
			return fmt.Errorf("invalid value object format: %T", v)
		}
	case *Object:
		x.Vals = v.Vals
	case map[string]*Value:
		x.Vals = v
	default:
		valueOf := reflect.ValueOf(v)
		typ := reflect.Indirect(valueOf).Type()
		if valueOf.Kind() == reflect.Ptr {
			valueOf = valueOf.Elem()
			typ = reflect.Indirect(valueOf).Type()
		}
		if typ.Kind() != reflect.Struct {
			return fmt.Errorf("invalid default object format: %T", v)
		}

		if _, ok := registerJSONEncoderTypes[typ.String()]; ok {
			if marshal, err := jsoniter.ConfigFastest.Marshal(v); err != nil {
				return err
			} else if err := jsoniter.ConfigFastest.Unmarshal(marshal, &x.Vals); err != nil {
				return err
			}
		} else {
			for i := 0; i < typ.NumField(); i++ {
				key := typ.Field(i).Name
				if key[0] >= 'a' && key[0] <= 'z' {
					continue
				}
				if valueOf.Field(i).IsZero() {
					continue
				}

				_val := &Value{}
				if encoder, ok := registerJSONEncoderTypeFields[typ.String()+"."+key]; ok {
					_typ := reflect2.TypeOf(valueOf)
					if _typ.Kind() == reflect.Ptr {
						_typ = _typ.(*reflect2.UnsafePtrType).Elem()
					}
					if obj, ok := _typ.(*reflect2.UnsafeStructType); ok {
						feild := obj.FieldByName(key)
						f := feild.UnsafeGet(reflect2.PtrOf(val))

						buf := &strings.Builder{}
						stream := jsoniter.NewStream(jsoniter.ConfigFastest, buf, 1024)
						encoder.Encode(f, stream)
						if err := jsoniter.ConfigFastest.Unmarshal(stream.Buffer(), &_val); err != nil {
							return err
						}
					}
				} else {
					var err error
					if _val, err = NewValue(valueOf.Field(i).Interface()); err != nil {
						return err
					}
				}
				x.Vals[strcase.ToLowerCamel(key)] = _val
			}
		}
	}
	return nil
}

func (x *Object) IsEmpty() bool {
	return x == nil || len(x.GetVals()) == 0
}

func (x *Object) init() {
	if x != nil && x.Vals == nil {
		x.Vals = make(map[string]*Value)
	}
}

func (x *Object) SetValue(key string, val *Value) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = val
	}
	return x
}

func (x *Object) SetBool(key string, val bool) *Object {
	x.SetValue(key, NewBoolValue(val))
	return x
}

func (x *Object) SetByte(key string, val []byte) *Object {
	x.SetValue(key, NewBytesValue(val))
	return x
}

func (x *Object) SetInt(key string, val int) *Object {
	x.SetValue(key, NewIntValue(val))
	return x
}

func (x *Object) SetInt32(key string, val int32) *Object {
	x.SetValue(key, NewInt32Value(val))
	return x
}

func (x *Object) SetInt64(key string, val int64) *Object {
	x.SetValue(key, NewInt64Value(val))
	return x
}

func (x *Object) SetUint(key string, val uint) *Object {
	x.SetValue(key, NewUintValue(val))
	return x
}

func (x *Object) SetUint32(key string, val uint32) *Object {
	x.SetValue(key, NewUint32Value(val))
	return x
}

func (x *Object) SetUint64(key string, val uint64) *Object {
	x.SetValue(key, NewUint64Value(val))
	return x
}

func (x *Object) SetFloat32(key string, val float32) *Object {
	x.SetValue(key, NewFloat32Value(val))
	return x
}

func (x *Object) SetFloat64(key string, val float64) *Object {
	x.SetValue(key, NewFloat64Value(val))
	return x
}

func (x *Object) SetString(key string, val string) *Object {
	x.SetValue(key, NewStringValue(val))
	return x
}

func (x *Object) SetObject(key string, val *Object) *Object {
	x.SetValue(key, NewObjectValue(val))
	return x
}

func (x *Object) SetIntArray(key string, vals ...int) *Object {
	x.SetValue(key, NewIntArrayValue(vals...))
	return x
}

func (x *Object) SetInt32Array(key string, vals ...int32) *Object {
	x.SetValue(key, NewInt32ArrayValue(vals...))
	return x
}

func (x *Object) SetInt64Array(key string, vals ...int64) *Object {
	x.SetValue(key, NewInt64ArrayValue(vals...))
	return x
}

func (x *Object) SetUintArray(key string, vals ...uint) *Object {
	x.SetValue(key, NewUintArrayValue(vals...))
	return x
}

func (x *Object) SetUint32Array(key string, vals ...uint32) *Object {
	x.SetValue(key, NewUint32ArrayValue(vals...))
	return x
}

func (x *Object) SetUint64Array(key string, vals ...uint64) *Object {
	x.SetValue(key, NewUint64ArrayValue(vals...))
	return x
}

func (x *Object) SetFloat32Array(key string, vals ...float32) *Object {
	x.SetValue(key, NewFloat32ArrayValue(vals...))
	return x
}

func (x *Object) SetFloat64Array(key string, vals ...float64) *Object {
	x.SetValue(key, NewFloat64ArrayValue(vals...))
	return x
}

func (x *Object) SetStringArray(key string, vals ...string) *Object {
	x.SetValue(key, NewStringArrayValue(vals...))
	return x
}

func (x *Object) SetObjectArray(key string, vals ...*Object) *Object {
	x.SetValue(key, NewObjectArrayValue(vals...))
	return x
}

func (x *Object) Merge(o *Object) *Object {
	if x != nil {
		for k, v := range o.Vals {
			x.Vals[k] = v
		}
	}
	return x
}

func (x *Object) Clone() *Object {
	if x != nil {
		obj := NewObject()
		for k, v := range x.Vals {
			obj.SetValue(k, v)
		}
		return obj
	}
	return x
}

func (x *Object) Delete(key string) *Object {
	if x != nil && x.Vals != nil {
		delete(x.Vals, key)
	}
	return x
}

func (x *Object) ToLowerCamelKeys() *Object {
	obj := NewObject()
	for k, v := range x.GetVals() {
		obj.SetValue(strcase.ToLowerCamel(k), v)
	}
	return obj
}

func (x *Object) ToSnakeKeys() *Object {
	obj := NewObject()
	for k, v := range x.GetVals() {
		obj.SetValue(strcase.ToSnake(k), v)
	}
	return obj
}

// MarshalJSON implements protobuf JSON encoding, not encoding/json semantics.
func (x *Object) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Object) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}
