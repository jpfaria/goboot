package newrelic

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Logger struct {
}

func NewLogger() newrelic.Logger {
	return &Logger{}
}

func (l *Logger) Error(msg string, context map[string]interface{}) {
	log.WithFields(context).Error(msg)
}

func (l *Logger) Warn(msg string, context map[string]interface{}) {
	log.WithFields(context).Warn(msg)
}

func (l *Logger) Info(msg string, context map[string]interface{}) {
	log.WithFields(context).Info(msg)
}

func (l *Logger) Debug(msg string, context map[string]interface{}) {
	log.WithFields(context).Debug(msg)
}

func (l *Logger) DebugEnabled() bool {
	return config.String(log.ConsoleLevel) == "DEBUG" || config.String(log.FileLevel) == "DEBUG" ||
		config.String(log.ConsoleLevel) == "TRACE" || config.String(log.FileLevel) == "TRACE"
}
