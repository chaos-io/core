package core

import (
	"strconv"
	"strings"
)

const (
	DoubleQuote = `"`
	SingleQuote = `'`
)

func IsQuotedString(str, quote string) bool {
	return strings.HasPrefix(str, quote) && strings.HasSuffix(str, quote)
}

func Quote(str string) string {
	return strconv.Quote(str)
}

func Unquote(str string) (string, error) {
	return strconv.Unquote(str)
}
