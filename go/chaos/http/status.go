package http

import (
	"strconv"
	"strings"
)

func NewStatus(statusCode int, status string) *Status {
	s := &Status{}

	if statusCode > 0 {
		s.Code = int32(statusCode)
		s.Reason = strings.TrimSpace(strings.TrimPrefix(status, strconv.Itoa(statusCode)))
	} else {
		fields := strings.Fields(status)
		if len(fields) > 0 {
			code, err := strconv.ParseInt(fields[0], 10, 32)
			if err == nil {
				s.Code = int32(code)
				s.Reason = strings.Join(fields[1:], " ")
			}
		}
	}

	return s
}
