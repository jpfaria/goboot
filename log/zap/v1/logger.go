package gizap

import (
	"context"
	"io"
	"os"
	"reflect"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey string

const key ctxKey = "ctxfields"

func NewLogger() gilog.Logger {

	cores := []zapcore.Core{}
	var writers []io.Writer

	if giconfig.Bool(gilog.ConsoleEnabled) {
		level := getZapLevel(giconfig.String(gilog.ConsoleLevel))
		writer := zapcore.Lock(os.Stdout)
		coreconsole := zapcore.NewCore(getEncoder(giconfig.String(ConsoleFormatter)), writer, level)
		cores = append(cores, coreconsole)
		writers = append(writers, writer)
	}

	if giconfig.Bool(gilog.FileEnabled) {
		s := []string{giconfig.String(gilog.FilePath), "/", giconfig.String(gilog.FileName)}
		fileLocation := strings.Join(s, "")

		lumber := &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  giconfig.Int(gilog.FileMaxSize),
			Compress: giconfig.Bool(gilog.FileCompress),
			MaxAge:   giconfig.Int(gilog.FileMaxAge),
		}

		level := getZapLevel(giconfig.String(gilog.FileLevel))
		writer := zapcore.AddSync(lumber)
		corefile := zapcore.NewCore(getEncoder(giconfig.String(FileFormatter)), writer, level)
		cores = append(cores, corefile)
		writers = append(writers, lumber)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	zaplogger := newSugaredLogger(combinedCore)

	newlogger := &zapLogger{
		fields:        gilog.Fields{},
		sugaredLogger: zaplogger,
		writers:       writers,
		core:          combinedCore,
	}

	gilog.NewLogger(newlogger)
	return newlogger
}

func newSugaredLogger(core zapcore.Core) *zap.SugaredLogger {
	return zap.New(core,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()
}

func getEncoder(format string) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	switch format {

	case "JSON":

		return zapcore.NewJSONEncoder(encoderConfig)

	default:

		return zapcore.NewConsoleEncoder(encoderConfig)

	}

}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "TRACE":
		return zapcore.DebugLevel
	case "WARN":
		return zapcore.WarnLevel
	case "DEBUG":
		return zapcore.DebugLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
	fields        gilog.Fields
	writers       []io.Writer
	core          zapcore.Core
}

func (l *zapLogger) Tracef(format string, args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Trace(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Print(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) WithField(key string, value interface{}) gilog.Logger {
	newFields := gilog.Fields{}
	for k, v := range l.fields {
		newFields[k] = v
	}

	newFields[key] = value

	f := mapToSlice(newFields)
	newLogger := newSugaredLogger(l.core).With(f...)
	return &zapLogger{newLogger, newFields, l.writers, l.core}
}

func (l *zapLogger) Output() io.Writer {
	return io.MultiWriter(l.writers...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Printf(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) WithFields(fields gilog.Fields) gilog.Logger {
	newFields := gilog.Fields{}

	for k, v := range l.fields {
		newFields[k] = v
	}

	for k, v := range fields {
		newFields[k] = v
	}

	f := mapToSlice(newFields)
	newLogger := newSugaredLogger(l.core).With(f...)
	return &zapLogger{newLogger, newFields, l.writers, l.core}
}

func (l *zapLogger) WithTypeOf(obj interface{}) gilog.Logger {

	t := reflect.TypeOf(obj)

	return l.WithFields(gilog.Fields{
		"reflect.type.name":    t.Name(),
		"reflect.type.package": t.PkgPath(),
	})
}

func (l *zapLogger) GetFields() gilog.Fields {
	return l.fields
}

func (l *zapLogger) ToContext(ctx context.Context) context.Context {
	fields := l.GetFields()

	ctxFields := fieldsFromContext(ctx)

	if ctxFields == nil {
		ctxFields = map[string]interface{}{}
	}

	for k, v := range fields {
		ctxFields[k] = v
	}

	return context.WithValue(ctx, key, ctxFields)
}

func (l *zapLogger) FromContext(ctx context.Context) gilog.Logger {
	fields := fieldsFromContext(ctx)
	return l.WithFields(fields)
}

func fieldsFromContext(ctx context.Context) gilog.Fields {
	fields := make(gilog.Fields)

	if ctx == nil {
		return fields
	}

	if f, ok := ctx.Value(key).(gilog.Fields); ok && f != nil {
		for k, v := range f {
			fields[k] = v
		}
	}

	return fields
}

func mapToSlice(m gilog.Fields) []interface{} {
	var f = make([]interface{}, 0)
	for k, v := range m {
		f = append(f, k)
		f = append(f, v)
	}

	return f
}
