package log

import (
	"context"
	"io"
)

// Logger is our contract for the logger
type Logger interface {
	Printf(format string, args ...interface{})

	Tracef(format string, args ...interface{})

	Trace(args ...interface{})

	Debugf(format string, args ...interface{})

	Debug(args ...interface{})

	Infof(format string, args ...interface{})

	Info(args ...interface{})

	Warnf(format string, args ...interface{})

	Warn(args ...interface{})

	Errorf(format string, args ...interface{})

	Error(args ...interface{})

	Fatalf(format string, args ...interface{})

	Fatal(args ...interface{})

	Panicf(format string, args ...interface{})

	Panic(args ...interface{})

	WithFields(keyValues Fields) Logger

	WithField(key string, value interface{}) Logger

	WithTypeOf(obj interface{}) Logger

	ToContext(ctx context.Context) context.Context

	FromContext(ctx context.Context) Logger

	Output() io.Writer
}
