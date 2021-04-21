package log

// A global variable so that l functions can be directly accessed
var l Logger

// NewLogger returns an instance of logger
func NewLogger(logger Logger) {
	l = logger
}

func Printf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	l.Warn(args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args...)
}

func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	l.Panic(args...)
}

func Tracef(format string, args ...interface{}) {
	l.Tracef(format, args...)
}

func Trace(args ...interface{}) {
	l.Trace(args...)
}

func WithFields(keyValues Fields) Logger {
	return l.WithFields(keyValues)
}

func WithField(key string, value string) Logger {
	return l.WithField(key, value)
}

func WithTypeOf(obj interface{}) Logger {
	return l.WithTypeOf(obj)
}

func GetLogger() Logger {
	return l
}
