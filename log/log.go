package log

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	_ "unsafe"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultLog *ZapLogger

type ZapLogger = zap.SugaredLogger

var defaultLevel zap.AtomicLevel

type Level = zapcore.Level

func init() {
	cfg := &Config{
		Level:  "debug",
		Encode: "console",
		Output: "console",
	}
	DefaultLog = newZap(cfg)
}

func Logger() *ZapLogger {
	return DefaultLog
}

func LevelEnabled(level Level) bool {
	return defaultLevel.Enabled(level)
}

func With(args ...interface{}) *ZapLogger {
	return DefaultLog.With(args...)
}

// level string, encode string, port int, pattern string, initFields map[string]interface{}
func newZap(cfg *Config) *ZapLogger {
	var opts []zap.Option
	opts = append(opts, zap.Development())
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))
	opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))

	defaultLevel = zap.NewAtomicLevel()
	SetLevel(cfg.Level)

	cores := make([]zapcore.Core, 0)
	output := strings.ToLower(cfg.Output)
	if strings.Contains(output, "console") {
		cores = append(cores, initWithConsole(cfg.Encode))
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core)

	initFields := cfg.InitFields
	if len(initFields) > 0 {
		initFieldList := make([]zap.Field, 0)
		for k, v := range initFields {
			initFieldList = append(initFieldList, zap.Any(k, v))
		}
		logger = logger.With(initFieldList...)
	}

	logger = logger.WithOptions(opts...)
	return logger.Sugar()
}

func SetLevel(level string) {
	l := strings.ToLower(level)
	if l == "info" {
		defaultLevel.SetLevel(zap.InfoLevel)
	} else if l == "debug" {
		defaultLevel.SetLevel(zap.DebugLevel)
	} else if l == "error" {
		defaultLevel.SetLevel(zap.ErrorLevel)
	} else if l == "warn" {
		defaultLevel.SetLevel(zap.WarnLevel)
	} else if l == "panic" {
		defaultLevel.SetLevel(zap.PanicLevel)
	} else if l == "fatal" {
		defaultLevel.SetLevel(zap.FatalLevel)
	} else {
		defaultLevel.SetLevel(zap.InfoLevel)
	}
}

func initWithConsole(encode string) zapcore.Core {
	formatEncoder := standardEncode(encode)

	consoleDebugging := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(formatEncoder, consoleDebugging, defaultLevel)
}

func standardEncode(encode string) zapcore.Encoder {
	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var formatEncoder zapcore.Encoder
	enc := strings.ToLower(encode)
	if enc == "json" {
		formatEncoder = zapcore.NewJSONEncoder(encodeConfig)
	} else {
		formatEncoder = zapcore.NewConsoleEncoder(encodeConfig)
	}

	return formatEncoder
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(args ...interface{}) {
	Logger().Debug(args...)
}

// Info logs a message at InfoLevel.
func Info(args ...interface{}) {
	Logger().Info(args...)
}

// Warn logs a message at WarnLevel.
func Warn(args ...interface{}) {
	Logger().Warn(args...)
}

// Error logs a message at ErrorLevel.
func Error(args ...interface{}) {
	Logger().Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	Logger().Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	Logger().Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	Logger().Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	Logger().Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	Logger().Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger().Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Logger().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Logger().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Logger().Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger().Fatalw(msg, keysAndValues...)
}

func NewError(args ...interface{}) error {
	Logger().Error(args...)
	return errors.New(fmt.Sprint(args...))
}

func NewErrorf(template string, args ...interface{}) error {
	Logger().Errorf(template, args...)
	return fmt.Errorf(template, args...)
}

func NewErrorw(msg string, keysAndValues ...interface{}) error {
	Logger().Errorw(msg, keysAndValues...)

	buffer := bytes.NewBufferString(msg)
	buffer.WriteString(" ")
	for i := 0; i < len(keysAndValues)-1; i += 2 {
		if i > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprintf("%v: %v", keysAndValues[i], keysAndValues[i+1]))
	}
	return errors.New(buffer.String())
}
