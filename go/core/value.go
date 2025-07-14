package core

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"unicode/utf8"

	"google.golang.org/protobuf/runtime/protoimpl"
)

const ValueTypeName = "Value"
const ValueTypeFullName = "core.Value"

// NewValue constructs a Value from a general-purpose Go interface.
//
//	╔═══════════════════════════════════════╤════════════════════════════════════════════╗
//	║ Go type                               │ Conversion                                 ║
//	╠═══════════════════════════════════════╪════════════════════════════════════════════╣
//	║ nil                                   │ stored as NullValue                        ║
//	║ bool                                  │ stored as BoolValue                        ║
//	║ int, int8, int16, int32, int64        │ stored as NumberValue                      ║
//	║ uint, uint8, uint16, uint32, uint64   │ stored as NumberValue                      ║
//	║ float32, float64                      │ stored as NumberValue                      ║
//	║ json.Number                           │ stored as NumberValue                      ║
//	║ string                                │ stored as StringValue; must be valid UTF-8 ║
//	║ []byte                                │ stored as StringValue; base64-encoded      ║
//	║ map[string]any                        │ stored as StructValue                      ║
//	║ []any                                 │ stored as ListValue                        ║
//	╚═══════════════════════════════════════╧════════════════════════════════════════════╝
//
// When converting an int64 or uint64 to a NumberValue, numeric precision loss
// is possible since they are stored as a float64.
func NewValue(v any) (*Value, error) {
	switch v := v.(type) {
	case nil:
		return NewNullValue(), nil
	case bool:
		return NewBoolValue(v), nil
	case int:
		return NewNumberValue(float64(v)), nil
	case int8:
		return NewNumberValue(float64(v)), nil
	case int16:
		return NewNumberValue(float64(v)), nil
	case int32:
		return NewNumberValue(float64(v)), nil
	case int64:
		return NewNumberValue(float64(v)), nil
	case uint:
		return NewNumberValue(float64(v)), nil
	case uint8:
		return NewNumberValue(float64(v)), nil
	case uint16:
		return NewNumberValue(float64(v)), nil
	case uint32:
		return NewNumberValue(float64(v)), nil
	case uint64:
		return NewNumberValue(float64(v)), nil
	case float32:
		return NewNumberValue(float64(v)), nil
	case float64:
		return NewNumberValue(v), nil
	case json.Number:
		n, err := v.Float64()
		if err != nil {
			return nil, protoimpl.X.NewError("invalid number format %q, expected a float64: %v", v, err)
		}
		return NewNumberValue(n), nil
	case string:
		if !utf8.ValidString(v) {
			return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", v)
		}
		return NewStringValue(v), nil
	case []byte:
		s := base64.StdEncoding.EncodeToString(v)
		return NewStringValue(s), nil
	case map[string]any:
		v2, err := NewObject(v)
		if err != nil {
			return nil, err
		}
		return NewObjectValue(v2), nil
	case []any:
		v2, err := NewValues(v)
		if err != nil {
			return nil, err
		}
		return NewValuesValue(v2), nil
	default:
		return nil, protoimpl.X.NewError("invalid type: %T", v)
	}
}

// NewNullValue constructs a new null Value.
func NewNullValue() *Value {
	return &Value{Kind: &Value_NullValue{NullValue: &Null{}}}
}

// NewBoolValue constructs a new boolean Value.
func NewBoolValue(v bool) *Value {
	return &Value{Kind: &Value_BoolValue{BoolValue: v}}
}

// NewNumberValue constructs a new number Value.
func NewNumberValue(v float64) *Value {
	return &Value{Kind: &Value_NumberValue{NumberValue: v}}
}

func NewIntValue(v int) *Value {
	return NewInt64Value(int64(v))
}

func NewInt8Value(v int8) *Value {
	return NewInt64Value(int64(v))
}

func NewInt16Value(v int16) *Value {
	return NewInt64Value(int64(v))
}

func NewInt32Value(v int32) *Value {
	return NewInt64Value(int64(v))
}

func NewInt64Value(v int64) *Value {
	if v >= 0 {
		return &Value{Kind: &Value_PositiveValue{PositiveValue: uint64(v)}}
	}

	var val uint64
	if v == math.MinInt64 {
		//val = uint64(math.MinInt64) + 1
		val = uint64(1) << 63
	} else {
		val = uint64(-v)
	}

	return &Value{Kind: &Value_PositiveValue{PositiveValue: val}}
}

func NewUintValue(v uint) *Value {
	return NewUint64Value(uint64(v))
}

func NewUint8Value(v uint8) *Value {
	return NewUint64Value(uint64(v))
}

func NewUint16Value(v uint16) *Value {
	return NewUint64Value(uint64(v))
}

func NewUint32Value(v uint32) *Value {
	return NewUint64Value(uint64(v))
}

func NewUint64Value(v uint64) *Value {
	return &Value{Kind: &Value_PositiveValue{PositiveValue: v}}
}

func NewFloat32Value(v float32) *Value {
	return NewFloat64Value(float64(v))
}

func NewFloat64Value(v float64) *Value {
	return &Value{Kind: &Value_NumberValue{NumberValue: v}}
}

// NewStringValue constructs a new string Value.
func NewStringValue(v string) *Value {
	return &Value{Kind: &Value_StringValue{StringValue: v}}
}

func NewBytesValue(v []byte) *Value {
	return &Value{Kind: &Value_BytesValue{BytesValue: v}}
}

func NewMapValue(v map[string]*Value) *Value {
	return &Value{Kind: &Value_ObjectValue{ObjectValue: &Object{Fields: v}}}
}

// NewObjectValue constructs a new struct Value.
func NewObjectValue(v *Object) *Value {
	return &Value{Kind: &Value_ObjectValue{ObjectValue: v}}
}

// NewValuesValue constructs a new list Value.
func NewValuesValue(v *Values) *Value {
	return &Value{Kind: &Value_ValuesValue{ValuesValue: v}}
}

func NewValueArray(values ...*Value) *Value {
	return &Value{Kind: &Value_ValuesValue{ValuesValue: &Values{Values: values}}}
}

// AsInterface converts x to a general-purpose Go interface.
//
// Calling Value.MarshalJSON and "encoding/json".Marshal on this output produce
// semantically equivalent JSON (assuming no errors occur).
//
// Floating-point values (i.e., "NaN", "Infinity", and "-Infinity") are
// converted as strings to remain compatible with MarshalJSON.
func (x *Value) AsInterface() any {
	switch v := x.GetKind().(type) {
	case *Value_NumberValue:
		if v != nil {
			switch {
			case math.IsNaN(v.NumberValue):
				return "NaN"
			case math.IsInf(v.NumberValue, +1):
				return "Infinity"
			case math.IsInf(v.NumberValue, -1):
				return "-Infinity"
			default:
				return v.NumberValue
			}
		}
	case *Value_StringValue:
		if v != nil {
			return v.StringValue
		}
	case *Value_BoolValue:
		if v != nil {
			return v.BoolValue
		}
	case *Value_ObjectValue:
		if v != nil {
			return v.ObjectValue.AsMap()
		}
	case *Value_ValuesValue:
		if v != nil {
			return v.ValuesValue.AsSlice()
		}
	}
	return nil
}

func (x *Value) GetBool() bool {
	return x.GetBoolValue()
}

func (x *Value) GetInt() int {
	return int(x.GetInt64())
}

func (x *Value) GetInt32() int32 {
	return int32(x.GetInt64())
}

func (x *Value) GetInt64() int64 {
	if negative := x.GetNegativeValue(); negative > 0 {
		if negative == uint64(math.MaxUint64)+1 {
			return math.MaxInt64
		}
		return -int64(negative)
	}
	return int64(x.GetPositiveValue())
}

func (x *Value) GetUint() uint {
	return uint(x.GetPositiveValue())
}

func (x *Value) GetUint32() uint32 {
	return uint32(x.GetPositiveValue())
}

func (x *Value) GetUint64() uint64 {
	return x.GetPositiveValue()
}

func (x *Value) GetFloat32() float32 {
	return float32(x.GetFloat64())
}

func (x *Value) GetFloat64() float64 {
	return x.GetNumberValue()
}

func (x *Value) GetDouble() float64 {
	return x.GetNumberValue()
}

func (x *Value) GetString() string {
	return x.GetStringValue()
}

func (x *Value) GetBytes() []byte {
	return x.GetBytesValue()
}

func (x *Value) GetObject() *Object {
	return x.GetObjectValue()
}

func (x *Value) GetValues() *Values {
	return x.GetValuesValue()
}

func (x *Value) GetValueArray() []*Value {
	if values := x.GetValuesValue(); values != nil {
		return values.Values
	}
	return nil
}

func (x *Value) GetObjectArray() []*Object {
	if values := x.GetValueArray(); len(values) > 0 {
		objs := make([]*Object, len(values))
		for _, v := range values {
			objs = append(objs, v.GetObject())
		}
		return objs
	}
	return nil
}
