package loggologrus

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggoLogFormatter struct{}

func (f *LoggoLogFormatter) Format(e *logrus.Entry) ([]byte, error) {
	m := e.Data
	m["timestamp"] = e.Time.UTC().Format(time.RFC3339Nano)
	m["level"] = strings.ToUpper(e.Level.String())
	m["message"] = e.Message

	b, err := json.Marshal(m)
	if err != nil {
		return b, err
	}

	b = append(b, '\n')
	return b, nil
}
