package core

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestNewUrlQueryFrom(t *testing.T) {
	tests := []struct {
		name   string
		values url.Values
		want   *Query
	}{
		{name: "empty", values: url.Values{}, want: &Query{Values: make(map[string]*StringValues)}},
		{name: "v1", values: url.Values{"key1": {"v1"}}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "v2", values: url.Values{"key1": {"v1"}, "key2": {"v2"}}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{"v2"}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUrlQueryFrom(tt.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrlQueryFrom() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

type MyToStringStruct struct{}

func (x *MyToStringStruct) ToString() string {
	return "this is my custom struct"
}

type MyFormatterStruct struct{}

func (x *MyFormatterStruct) Format() string {
	return "this is a formatter"
}

func TestQuery_Add(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		Values map[string]*StringValues
		args   args
		want   *Query
	}{
		{name: "empty", Values: map[string]*StringValues{}, args: args{}, want: &Query{Values: map[string]*StringValues{}}},
		{name: "add-nil", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "add-empty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "add-nonempty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1", v: "v11"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1", "v11"}}}}},
		{name: "add-other-empty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {}}}},
		{name: "add-other-nonempty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2", v: "v2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{"v2"}}}}},
		{name: "add-key-bool-true", Values: map[string]*StringValues{}, args: args{k: "key2", v: true}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"true"}}}}},
		{name: "add-key-bool-false", Values: map[string]*StringValues{}, args: args{k: "key2", v: false}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"false"}}}}},
		{name: "add-key-int8", Values: map[string]*StringValues{}, args: args{k: "key2", v: int8(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-int16", Values: map[string]*StringValues{}, args: args{k: "key2", v: int16(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-int32", Values: map[string]*StringValues{}, args: args{k: "key2", v: int32(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-int64", Values: map[string]*StringValues{}, args: args{k: "key2", v: int64(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-int", Values: map[string]*StringValues{}, args: args{k: "key2", v: int(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-uint8", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint8(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-uint16", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint16(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-uint32", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint32(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-uint64", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint64(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-uint", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-float32", Values: map[string]*StringValues{}, args: args{k: "key2", v: float32(13.0)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-float64", Values: map[string]*StringValues{}, args: args{k: "key2", v: float64(13.0)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "add-key-string-array", Values: map[string]*StringValues{}, args: args{k: "key2", v: []string{"v1", "v2"}}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"v1", "v2"}}}}},
		{name: "add-key-ToStringConverter", Values: map[string]*StringValues{}, args: args{k: "key2", v: &MyToStringStruct{}}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"this is my custom struct"}}}}},
		{name: "add-key-Formatter", Values: map[string]*StringValues{}, args: args{k: "key2", v: &MyFormatterStruct{}}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"this is a formatter"}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Query{
				Values: tt.Values,
			}
			if got := x.Add(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Set(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		Values map[string]*StringValues
		args   args
		want   *Query
	}{
		{name: "empty", Values: map[string]*StringValues{}, args: args{}, want: &Query{Values: map[string]*StringValues{}}},
		{name: "set-nil", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "set-empty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{}}}}},
		{name: "set-nonempty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1", v: "v11"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v11"}}}}},
		{name: "set-other-empty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{}}}}},
		{name: "set-other-nonempty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2", v: "v2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{"v2"}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Query{
				Values: tt.Values,
			}
			if got := x.Set(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Unmarshal_StringSlice(t *testing.T) {
	tests := []struct {
		name    string
		query   *Query
		k       string
		typ     reflect.Type
		want    any
		wantErr bool
	}{
		{name: "string-slice-1", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{"bar", "bba"}}}}, k: "foo", typ: reflect.TypeOf([]string{}), want: []string{"bar", "bba"}, wantErr: false},
		{name: "string-slice-2", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{"bar,bba"}}}}, k: "foo", typ: reflect.TypeOf([]string{}), want: []string{"bar", "bba"}, wantErr: false},
		{name: "string-slice-3", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{`"bar", "bba"`}}}}, typ: reflect.TypeOf([]string{}), k: "foo", want: []string{"bar", "bba"}, wantErr: false},
		{name: "string-slice-4", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{`"bar,bba`}}}}, typ: reflect.TypeOf([]string{}), k: "foo", want: []string{"\"bar", "bba"}, wantErr: false},
		{name: "string-slice-5", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{`["bar", "bba"]`}}}}, typ: reflect.TypeOf([]string{}), k: "foo", want: []string{"bar", "bba"}, wantErr: false},
		{name: "int-slice-1", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{"123", "234"}}}}, k: "foo", typ: reflect.TypeOf([]int32{}), want: []int32{123, 234}, wantErr: false},
		{name: "int-slice-2", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{"123,234"}}}}, k: "foo", typ: reflect.TypeOf([]int32{}), want: []int32{123, 234}, wantErr: false},
		//{name: "int32Value-slice-1", query: &Query{Values: map[string]*StringValues{"foo": {Values: []string{"123,234"}}}}, k: "foo", typ: reflect.TypeOf([]*Int32Value{}), want: []*Int32Value{{Value: 123}, {Value: 234}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ptr := reflect.New(tt.typ)
			if err := tt.query.Unmarshal(tt.k, ptr.Interface()); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			vals := ptr.Elem().Interface()
			if !reflect.DeepEqual(vals, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", vals, tt.want)
			}
		})
	}
}

func Test_splitQuoteString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{name: "empty", s: "", want: []string{""}},
		{name: "1 value", s: `"a"`, want: []string{`"a"`}},
		{name: "1 values-2", s: `"a,b"`, want: []string{`"a,b"`}},
		{name: "2 values", s: `"a","b"`, want: []string{`"a"`, `"b"`}},
		{name: "2 values-2", s: `"a,c","b"`, want: []string{`"a,c"`, `"b"`}},
		{name: "2 values-3", s: `"a,c" , "b"`, want: []string{`"a,c"`, `"b"`}},
		{name: "3 values", s: "a,b,c", want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitQuoteString(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitQuoteString() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_isStringSlice(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want bool
	}{
		{name: "strings", v: []string{}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringSlice(tt.v); got != tt.want {
				t.Errorf("isStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unmarshal(t *testing.T) {
	arrayStr := `["bar", "bba"]`
	var v []string
	err := jsoniter.ConfigFastest.UnmarshalFromString(arrayStr, &v)
	fmt.Println(err)
	fmt.Println(v)
}
