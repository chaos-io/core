package testdata

import (
	"testing"

	"github.com/chaos-io/core/go/logs"
)

func TestLogs(t *testing.T) {
	logs.Debugw("TestLogs", "123", 2222)
	logs.Infow("TestLogs", "123", 2222)

	logs.Logger().Debugw("logger TestLogs", "123", 2222)
}
