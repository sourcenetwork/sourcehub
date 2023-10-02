package zanzi

import (
	"github.com/cometbft/cometbft/libs/log"
)

// loggerWrapper wraps a cosmos SDK logger into a zanzi compatible logger
type loggerWrapper struct {
	logger log.Logger
}

func (l *loggerWrapper) Infof(fmt string, args ...any) {
	l.logger.Info(fmt, args...)
}

func (l *loggerWrapper) Debugf(fmt string, args ...any) {
	l.logger.Debug(fmt, args...)
}

func (l *loggerWrapper) Errorf(fmt string, args ...any) {
	l.logger.Error(fmt, args...)
}

func (l *loggerWrapper) Warnf(fmt string, args ...any) {
	l.logger.Info(fmt, args...)
}
