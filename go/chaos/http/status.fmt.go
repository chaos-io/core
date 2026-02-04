package http

import (
	"bytes"
	"strconv"
	"strings"
)

func ParseStatus(status string) (*Status, error) {
	s := &Status{}
	if err := s.Parse(status); err != nil {
		return nil, err
	}
	return s, nil
}

func (x *Status) Parse(status string) error {
	if x != nil && status != "" {
		fields := strings.Fields(status)
		if len(fields) > 0 {
			code, err := strconv.ParseInt(fields[0], 10, 32)
			if err == nil {
				x.Code = int32(code)
				x.Reason = strings.Join(fields[1:], " ")
			} else {
				return err
			}
		}
	}
	return nil
}

func (x *Status) Format() string {
	if x != nil {
		buf := bytes.Buffer{}
		buf.WriteString(strconv.FormatInt(int64(x.Code), 10))
		if x.Reason != "" {
			buf.WriteString(" ")
			buf.WriteString(x.Reason)
		}
		return buf.String()
	}
	return ""
}
