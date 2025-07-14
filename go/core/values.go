package core

import "google.golang.org/protobuf/encoding/protojson"

// NewValues constructs a ListValue from a general-purpose Go slice.
// The slice elements are converted using NewValue.
func NewValues(v []any) (*Values, error) {
	x := &Values{Values: make([]*Value, len(v))}
	for i, v := range v {
		var err error
		x.Values[i], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

// AsSlice converts x to a general-purpose Go slice.
// The slice elements are converted by calling Value.AsInterface.
func (x *Values) AsSlice() []any {
	vals := x.GetValues()
	vs := make([]any, len(vals))
	for i, v := range vals {
		vs[i] = v.AsInterface()
	}
	return vs
}

func (x *Values) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Values) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}
