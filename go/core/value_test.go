package core

import (
	"reflect"
	"testing"
)

func TestNewInt64Value(t *testing.T) {
	tests := []struct {
		name string
		v    int64
		want *Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInt64Value(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt64Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
