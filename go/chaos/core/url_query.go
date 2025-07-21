package core

import (
	"net/url"
	"strconv"
	"strings"
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
