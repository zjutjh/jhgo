package nesty

import "github.com/sirupsen/logrus"

func newLogger(l *logrus.Logger) *httpLogger {
	return &httpLogger{logger: l}
}

type httpLogger struct {
	logger *logrus.Logger
}

func (l *httpLogger) Errorf(format string, v ...any) {
	if l.logger != nil {
		l.logger.Errorf(format, v...)
	}
}

func (l *httpLogger) Warnf(format string, v ...any) {
	if l.logger != nil {
		l.logger.Warnf(format, v...)
	}
}

func (l *httpLogger) Debugf(format string, v ...any) {
	if l.logger != nil {
		l.logger.Infof(format, v...)
	}
}
