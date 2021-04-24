package logger

import (
	"github.com/b2wdigital/goignite/v2/contrib/go.uber.org/zap.v1"
	"github.com/b2wdigital/goignite/v2/contrib/rs/zerolog.v1"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
)

func New() {
	switch Impl() {
	case "ZEROLOG":
		zerolog.NewLogger()
	case "ZAP":
		zap.NewLogger()
	default:
		logrus.NewLogger()
	}
}
