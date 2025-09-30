//go:build windows

package httpserver

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// initGinLoggerWriter 初始化gin access logger writer error logger writer
func initGinLoggerWriter(conf Config) (io.Writer, io.Writer, error) {
	// access logger
	var aw io.Writer
	if conf.Log.AccessFilename == "/dev/stdout" {
		aw = os.Stdout
	} else {
		_, err := os.OpenFile(conf.Log.AccessFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, nil, err
		}
		aw = &lumberjack.Logger{
			Filename:   conf.Log.AccessFilename,
			MaxSize:    conf.Log.MaxSize,
			MaxAge:     conf.Log.MaxAge,
			MaxBackups: conf.Log.MaxBackups,
			LocalTime:  conf.Log.LocalTime,
			Compress:   conf.Log.Compress,
		}
	}

	// error logger
	var ew io.Writer
	if conf.Log.ErrorFilename == "/dev/stderr" {
		ew = os.Stderr
	} else {
		_, err := os.OpenFile(conf.Log.ErrorFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, nil, err
		}
		ew = &lumberjack.Logger{
			Filename:   conf.Log.ErrorFilename,
			MaxSize:    conf.Log.MaxSize,
			MaxAge:     conf.Log.MaxAge,
			MaxBackups: conf.Log.MaxBackups,
			LocalTime:  conf.Log.LocalTime,
			Compress:   conf.Log.Compress,
		}
	}

	return aw, ew, nil
}
