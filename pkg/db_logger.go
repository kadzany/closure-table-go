package pkg

import (
	"context"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sirupsen/logrus"
)

type logrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter(logger *logrus.Logger) sqldblogger.Logger {
	return &logrusAdapter{logger: logger}
}

func (l *logrusAdapter) Log(ctx context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {
	entry := l.logger.WithContext(ctx).WithFields(data)

	switch level {
	case sqldblogger.LevelError:
		entry.Error(msg)
	case sqldblogger.LevelInfo:
		entry.Info(msg)
	case sqldblogger.LevelDebug:
		entry.Debug(msg)
	case sqldblogger.LevelTrace:
		entry.Trace(msg)
	default:
		entry.Debug(msg)
	}
}
