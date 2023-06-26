package recovery

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chaos-io/core/log"
)

func TestWithLogger(t *testing.T) {
	var mw middleware

	logger := log.DefaultLog
	opt := WithLogger(logger)
	opt(&mw)

	assert.Same(t, logger, mw.l)
}

func TestWithCallBack(t *testing.T) {
	var mw middleware

	callback := func(_ http.ResponseWriter, _ *http.Request, _ error) {}

	opt := WithCallBack(callback)
	opt(&mw)

	assert.Equal(t, fmt.Sprintf("%p", callback), fmt.Sprintf("%p", mw.panicCallback))
}
