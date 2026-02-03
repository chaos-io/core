package core

import jsoniter "github.com/json-iterator/go"

func (v *BoolValue) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *BoolValue) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (vs *BoolValues) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *BoolValues) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (v *Int32Value) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *Int32Value) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *Int64Value) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *Int64Value) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *Uint32Value) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *Uint32Value) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *Uint64Value) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *Uint64Value) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *Float32Value) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *Float32Value) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *Float64Value) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *Float64Value) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *StringValue) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *StringValue) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (v *BytesValue) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(v.Val)
}

func (v *BytesValue) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &v.Val)
}

func (vs *Int32Values) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *Int32Values) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (vs *Int64Values) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *Int64Values) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (vs *Uint32Values) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *Uint32Values) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (vs *Uint64Values) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *Uint64Values) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (vs *Float32Values) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *Float32Values) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (vs *Float64Values) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *Float64Values) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (vs *StringValues) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(vs.Vals)
}

func (vs *StringValues) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &vs.Vals)
}

func (m *StringMap) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(m.Vals)
}

func (m *StringMap) UnmarshalJSON(data []byte) error {
	return jsoniter.Unmarshal(data, &m.Vals)
}

func (m *StringsMap) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}

	tmp := make(map[string]*StringValues, len(m.Vals))
	for k, v := range m.Vals {
		if v != nil {
			tmp[k] = v
		}
	}

	return jsoniter.Marshal(tmp)
}

func (m *StringsMap) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	tmp := make(map[string]*StringValues)
	if err := jsoniter.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if m.Vals == nil {
		m.Vals = make(map[string]*StringValues, len(tmp))
	}

	for k, v := range tmp {
		m.Vals[k] = v
	}

	return nil
}
