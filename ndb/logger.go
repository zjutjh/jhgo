package ndb

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/sirupsen/logrus"
)

type dbLogger struct {
	logger.Config
	l *logrus.Logger
}

func newDBLogger(l *logrus.Logger, config logger.Config) logger.Interface {
	return &dbLogger{
		Config: config,
		l:      l,
	}
}

// LogMode log mode
func (l *dbLogger) LogMode(level logger.LogLevel) logger.Interface {
	nl := *l
	nl.LogLevel = level
	return &nl
}

// Info print info
func (l *dbLogger) Info(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= logger.Info {
		l.l.WithContext(ctx).WithField("line", utils.FileWithLineNum()).Infof(msg, data...)
	}
}

// Warn print warn messages
func (l *dbLogger) Warn(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= logger.Warn {
		l.l.WithContext(ctx).WithField("line", utils.FileWithLineNum()).Warnf(msg, data...)
	}
}

// Error print error messages
func (l *dbLogger) Error(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= logger.Error {
		l.l.WithContext(ctx).WithField("line", utils.FileWithLineNum()).Errorf(msg, data...)
	}
}

// Trace print sql message
func (l *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.l.WithContext(ctx).WithError(err).WithFields(logrus.Fields{
			"line": utils.FileWithLineNum(),
			"cost": elapsed.String(),
			"rows": rows,
			"sql":  sql,
		}).Error("执行SQL错误")
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.l.WithContext(ctx).WithFields(logrus.Fields{
			"line":      utils.FileWithLineNum(),
			"cost":      elapsed.String(),
			"rows":      rows,
			"sql":       sql,
			"threshold": l.SlowThreshold.String(),
		}).Warn("执行SQL时长超过期望")
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		l.l.WithContext(ctx).WithFields(logrus.Fields{
			"line": utils.FileWithLineNum(),
			"cost": elapsed.String(),
			"rows": rows,
			"sql":  sql,
		}).Info("执行SQL记录")
	}
}

// ParamsFilter filter params
func (l *dbLogger) ParamsFilter(ctx context.Context, sql string, params ...any) (string, []any) {
	if l.Config.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}
