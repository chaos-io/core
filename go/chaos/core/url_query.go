package core

import (
	"encoding/csv"
	"io"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func NewUrlQuery(kvs ...any) *Query {
	query := &Query{
		Values: make(map[string]*StringValues),
	}

	for i := 0; i < len(kvs)-1; i += 2 {
		key := kvs[i].(string)
		query.Add(key, kvs[i+1])
	}
	return query
}

func NewUrlQueryFrom(values url.Values) *Query {
	query := &Query{
		Values: make(map[string]*StringValues),
	}
	return query.FromUrlValues(values)
}

func (x *Query) FromUrlValues(values url.Values) *Query {
	if x != nil {
		if x.Values == nil {
			x.Values = make(map[string]*StringValues)
		}

		for k, v := range values {
			x.Values[k] = &StringValues{Values: v}
		}
	}
	return x
}

func (x *Query) Has(k string) bool {
	if x != nil {
		_, ok := x.Values[k]
		return ok
	}
	return false
}

func (x *Query) Add(k string, v any) *Query {
	k = strings.TrimSpace(k)
	if x != nil && k != "" {
		if x.Values[k] == nil {
			x.Values[k] = &StringValues{}
		}
		x.Values[k].Values = append(x.Values[k].Values, queryValueFormat(v)...)
	}
	return x
}

func (x *Query) Set(k string, v any) *Query {
	k = strings.TrimSpace(k)
	if x != nil && k != "" {
		x.Values[k] = &StringValues{
			Values: queryValueFormat(v),
		}
	}
	return x
}

func (x *Query) Del(k string) *Query {
	k = strings.TrimSpace(k)
	if x != nil && k != "" {
		delete(x.Values, k)
	}
	return x
}

func queryValueFormat(val any) []string {
	switch v := val.(type) {
	case bool:
		return []string{strconv.FormatBool(v)}
	case int8:
		return []string{strconv.FormatInt(int64(v), 10)}
	case int16:
		return []string{strconv.FormatInt(int64(v), 10)}
	case int32:
		return []string{strconv.FormatInt(int64(v), 10)}
	case int64:
		return []string{strconv.FormatInt(v, 10)}
	case int:
		return []string{strconv.FormatInt(int64(v), 10)}
	case uint8:
		return []string{strconv.FormatUint(uint64(v), 10)}
	case uint16:
		return []string{strconv.FormatUint(uint64(v), 10)}
	case uint32:
		return []string{strconv.FormatUint(uint64(v), 10)}
	case uint64:
		return []string{strconv.FormatUint(v, 10)}
	case uint:
		return []string{strconv.FormatUint(uint64(v), 10)}
	case float32:
		return []string{strconv.FormatFloat(float64(v), 'g', -1, 32)}
	case float64:
		return []string{strconv.FormatFloat(v, 'g', -1, 64)}
	case string:
		return []string{v}
	case []string:
		return v
	}

	if c, ok := val.(ToStringConverter); ok {
		return []string{c.ToString()}
	}

	if f, ok := val.(Formatter); ok {
		return []string{f.Format()}
	}

	return []string{}
}

// Unmarshal
// normal: name=Tom
// list:
//
//	ids=1&ids=2
//	ids[0]=6&ids[1]=7
//	ids=1,2,3
//	ids=[1,2,3]
//
// map:
//
//	user[name]=Tom
//
// object:
//
//	filter={"page":1}
func (x *Query) Unmarshal(k string, v any) error {
	if x == nil || x.Values == nil {
		return nil
	}

	param, ok := x.Values[k]
	if !ok || param == nil || len(param.Values) == 0 {
		return nil
	}

	kind := reflect.Indirect(reflect.ValueOf(v)).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return decodeSlice(param.Values, v)
	}

	return decodeValue(param.Values[0], v)
}

func UnmarshalParam(str string, v any) error {
	if str == "" {
		return nil
	}

	kind := reflect.Indirect(reflect.ValueOf(v)).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return decodeSlice(strings.Split(str, ","), v)
	}

	return decodeValue(str, v)
}

func decodeSlice(values []string, v any) error {
	if len(values) == 1 {
		valStr := strings.TrimSpace(values[0])
		if strings.HasPrefix(valStr, "[") {
			return jsoniter.ConfigFastest.UnmarshalFromString(valStr, &v)
		} else {
			if isStringSlice(v) {
				values = splitQuoteString(valStr)
			} else {
				valStr = "[" + valStr + "]"
				return jsoniter.ConfigFastest.UnmarshalFromString(valStr, &v)
			}
		}
	}

	if isStringSlice(v) {
		for i, val := range values {
			val = strings.TrimSpace(val)
			if !IsQuotedString(val, DoubleQuote) {
				values[i] = strconv.Quote(val)
			}
		}
	}

	arrayStr := "[" + strings.Join(values, ",") + "]"
	return jsoniter.ConfigFastest.UnmarshalFromString(arrayStr, &v)
}

func decodeValue(value string, v any) error {
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() == reflect.String {
		val.SetString(value)
		return nil
	}
	return jsoniter.ConfigFastest.UnmarshalFromString(value, v)
}

func isStringSlice(v any) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
		if t.Kind() == reflect.String {
			return true
		}
	}
	return false
}

var separator = regexp.MustCompile(`" *, *"`)

func splitQuoteString(s string) []string {
	if !IsQuotedString(s, DoubleQuote) {
		return strings.Split(s, `,`)
	}

	vals := separator.Split(s, -1)
	if len(vals) > 1 {
		vals[0] += `"`
		vals[len(vals)-1] = `"` + vals[len(vals)-1]

		for i := 1; i < len(vals)-1; i++ {
			vals[i] = strconv.Quote(strings.TrimSpace(vals[i]))
		}
	}

	return vals
}

func splitQuoteString2(s string) []string {
	r := csv.NewReader(strings.NewReader(s))
	r.TrimLeadingSpace = true

	record, err := r.Read()
	if err != nil && err != io.EOF {
		parts := strings.Split(s, ",")
		for i, part := range parts {
			parts[i] = strconv.Quote(strings.TrimSpace(part))
		}
		return parts
	}

	for i, val := range record {
		record[i] = `"` + val + `"`
		//record[i] = strconv.Quote(val)
	}
	return record
}
