package core

import (
	"net/url"
	"reflect"
	"testing"
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
		{name: "v1-add-nil", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "v1-add-empty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "v1-add-nonempty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1", v: "v11"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1", "v11"}}}}},
		{name: "v1-add-other-empty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {}}}},
		{name: "v1-add-other-nonempty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2", v: "v2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{"v2"}}}}},
		{name: "v1-add-key-bool-true", Values: map[string]*StringValues{}, args: args{k: "key2", v: true}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"true"}}}}},
		{name: "v1-add-key-bool-false", Values: map[string]*StringValues{}, args: args{k: "key2", v: false}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"false"}}}}},
		{name: "v1-add-key-int8", Values: map[string]*StringValues{}, args: args{k: "key2", v: int8(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-int16", Values: map[string]*StringValues{}, args: args{k: "key2", v: int16(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-int32", Values: map[string]*StringValues{}, args: args{k: "key2", v: int32(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-int64", Values: map[string]*StringValues{}, args: args{k: "key2", v: int64(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-int", Values: map[string]*StringValues{}, args: args{k: "key2", v: int(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-uint8", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint8(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-uint16", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint16(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-uint32", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint32(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-uint64", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint64(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-uint", Values: map[string]*StringValues{}, args: args{k: "key2", v: uint(13)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-float32", Values: map[string]*StringValues{}, args: args{k: "key2", v: float32(13.0)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-float64", Values: map[string]*StringValues{}, args: args{k: "key2", v: float64(13.0)}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"13"}}}}},
		{name: "v1-add-key-string-array", Values: map[string]*StringValues{}, args: args{k: "key2", v: []string{"v1", "v2"}}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"v1", "v2"}}}}},
		{name: "v1-add-key-ToStringConverter", Values: map[string]*StringValues{}, args: args{k: "key2", v: &MyToStringStruct{}}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"this is my custom struct"}}}}},
		{name: "v1-add-key-Formatter", Values: map[string]*StringValues{}, args: args{k: "key2", v: &MyFormatterStruct{}}, want: &Query{Values: map[string]*StringValues{"key2": {Values: []string{"this is a formatter"}}}}},
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
		{name: "v1-set-nil", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}}},
		{name: "v1-set-empty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{}}}}},
		{name: "v1-set-nonempty-value", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key1", v: "v11"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v11"}}}}},
		{name: "v1-set-other-empty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{}}}}},
		{name: "v1-set-other-nonempty-key", Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}}, args: args{k: "key2", v: "v2"}, want: &Query{Values: map[string]*StringValues{"key1": {Values: []string{"v1"}}, "key2": {Values: []string{"v2"}}}}},
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
