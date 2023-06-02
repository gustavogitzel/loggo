package loggologrus

import (
	"github.com/gustavogitzel/loggo/pkg/loggo"
	"github.com/sirupsen/logrus"
)

type LoggoLogrusBackend struct {
	logger *logrus.Logger
}

func NewLoggoLogrusBackend(logger *logrus.Logger) loggo.Backend {
	return &LoggoLogrusBackend{logger}
}

func (b *LoggoLogrusBackend) Log(fields loggo.Fields, level string, args ...interface{}) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		b.logger.Warn("Failed to parse log level: ", level, ". Falling back to INFO.")
		lvl = logrus.InfoLevel
	}
	b.logger.WithFields(fields).Log(lvl, args...)
}
