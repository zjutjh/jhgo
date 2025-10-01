//go:build !windows

package nlog

import (
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/sirupsen/logrus"

	"github.com/zjutjh/mygo/config"
)

func New(conf Config) *logrus.Logger {
	logger := logrus.New()

	// 设置Output
	os.OpenFile(conf.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxAge:     conf.MaxAge,
		MaxBackups: conf.MaxBackups,
		LocalTime:  conf.LocalTime,
		Compress:   conf.Compress,
	})

	// 设置Formatter
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "ts",
		},
	})

	// 设置Level
	logger.SetLevel(conf.Level)

	// 设置Hook
	logger.AddHook(&hookField{
		app: config.AppName(),
	})

	return logger
}
