package xerrors

import (
	"fmt"

	"github.com/chaos-io/core/x/xruntime"
	modes2 "github.com/chaos-io/core/xerrors/internal/modes"
)

func DefaultStackTraceMode() {
	modes2.DefaultStackTraceMode()
}

func EnableFrames() {
	modes2.SetStackTraceMode(modes2.StackTraceModeFrames)
}

func EnableStacks() {
	modes2.SetStackTraceMode(modes2.StackTraceModeStacks)
}

func EnableStackThenFrames() {
	modes2.SetStackTraceMode(modes2.StackTraceModeStackThenFrames)
}

func EnableStackThenNothing() {
	modes2.SetStackTraceMode(modes2.StackTraceModeStackThenNothing)
}

func DisableStackTraces() {
	modes2.SetStackTraceMode(modes2.StackTraceModeNothing)
}

// newStackTrace returns stacktrace based on current mode and frames count
func newStackTrace(skip int, err error) *xruntime.StackTrace {
	skip++
	m := modes2.GetStackTraceMode()
	switch m {
	case modes2.StackTraceModeFrames:
		return xruntime.NewFrame(skip)
	case modes2.StackTraceModeStackThenFrames:
		if err != nil && StackTraceOfEffect(err) != nil {
			return xruntime.NewFrame(skip)
		}

		return _newStackTrace(skip)
	case modes2.StackTraceModeStackThenNothing:
		if err != nil && StackTraceOfEffect(err) != nil {
			return nil
		}

		return _newStackTrace(skip)
	case modes2.StackTraceModeStacks:
		return _newStackTrace(skip)
	case modes2.StackTraceModeNothing:
		return nil
	}

	panic(fmt.Sprintf("unknown stack trace mode %d", m))
}

func MaxStackFrames16() {
	modes2.SetStackFramesCountMax(modes2.StackFramesCount16)
}

func MaxStackFrames32() {
	modes2.SetStackFramesCountMax(modes2.StackFramesCount32)
}

func MaxStackFrames64() {
	modes2.SetStackFramesCountMax(modes2.StackFramesCount64)
}

func MaxStackFrames128() {
	modes2.SetStackFramesCountMax(modes2.StackFramesCount128)
}

func _newStackTrace(skip int) *xruntime.StackTrace {
	skip++
	count := modes2.GetStackFramesCountMax()
	switch count {
	case 16:
		return xruntime.NewStackTrace16(skip)
	case 32:
		return xruntime.NewStackTrace32(skip)
	case 64:
		return xruntime.NewStackTrace64(skip)
	case 128:
		return xruntime.NewStackTrace128(skip)
	}

	panic(fmt.Sprintf("unknown stack frames count %d", count))
}
