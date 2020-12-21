package gizerolog

import (
	"bytes"
	"context"
	"io"
	"os"
	"reflect"
	"strings"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() gilog.Logger {
	fileEnabled := giconfig.Bool(gilog.FileEnabled)
	consoleEnabled := giconfig.Bool(gilog.ConsoleEnabled)

	format := giconfig.String(Formatter)
	writer := getWriter(format, fileEnabled, consoleEnabled)
	if writer == nil {
		zerologger := zerolog.Nop()
		logger := &logger{
			logger: zerologger,
		}

		gilog.NewLogger(logger)
		return logger
	}

	zerolog.MessageFieldName = "log_message"
	zerolog.LevelFieldName = "log_level"

	zerologger := zerolog.New(writer).With().Timestamp().Logger()
	level := getLogLevel(giconfig.String(gilog.ConsoleLevel))
	zerologger = zerologger.Level(level)

	logger := &logger{
		logger: zerologger,
		writer: writer,
	}

	gilog.NewLogger(logger)
	return logger
}

type logger struct {
	logger zerolog.Logger
	writer io.Writer
}

func getLogLevel(level string) zerolog.Level {
	switch level {
	case "DEBUG":
		return zerolog.DebugLevel
	case "WARN":
		return zerolog.WarnLevel
	case "FATAL":
		return zerolog.FatalLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "TRACE":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func getWriter(format string, fileEnabled bool, consoleEnabled bool) io.Writer {
	var writer io.Writer
	switch format {
	case "TEXT":
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	default:
		writer = os.Stdout
	}

	if fileEnabled {
		s := []string{giconfig.String(gilog.FilePath), "/", giconfig.String(gilog.FileName)}
		fileLocation := strings.Join(s, "")

		fileHandler := &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  giconfig.Int(gilog.FileMaxSize),
			Compress: giconfig.Bool(gilog.FileCompress),
			MaxAge:   giconfig.Int(gilog.FileMaxAge),
		}

		if consoleEnabled {
			return io.MultiWriter(writer, fileHandler)
		} else {
			return fileHandler
		}
	} else if consoleEnabled {
		return writer
	}

	return nil
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.logger.Trace().Msgf(format, args...)
}

func (l *logger) Trace(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Trace().Msgf(format.String(), args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Debug().Msgf(format.String(), args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *logger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Info().Msgf(format.String(), args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Warn().Msgf(format.String(), args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Error().Msgf(format.String(), args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Fatal().Msgf(format.String(), args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

func (l *logger) Panic(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Panic().Msgf(format.String(), args...)
}

func (l *logger) WithField(key string, value interface{}) gilog.Logger {
	newField := gilog.Fields{}
	newField[key] = value

	newLogger := l.logger.With().Fields(newField).Logger()
	return &logger{newLogger, l.writer}
}

func (l *logger) WithFields(fields gilog.Fields) gilog.Logger {
	newLogger := l.logger.With().Fields(fields).Logger()
	return &logger{newLogger, l.writer}
}

func (l *logger) WithTypeOf(obj interface{}) gilog.Logger {

	t := reflect.TypeOf(obj)

	return l.WithFields(gilog.Fields{
		"reflect.type.name":    t.Name(),
		"reflect.type.package": t.PkgPath(),
	})
}

func (l *logger) Output() io.Writer {
	return l.writer
}

func (l *logger) ToContext(ctx context.Context) context.Context {
	logger := l.logger
	return logger.WithContext(ctx)
}

func (l *logger) FromContext(ctx context.Context) gilog.Logger {
	zerologger := zerolog.Ctx(ctx)
	if zerologger.GetLevel() == zerolog.Disabled {
		return l
	}

	return &logger{*zerologger, l.writer}
}
