package crontab

import "time"

var DefaultConfig = Config{
	ShutdownWaitTimeout: 10 * time.Second,

	Log: LogConfig{
		ErrorFilename: "./logs/cron.log",
		MaxSize:       100,
		MaxAge:        7,
		MaxBackups:    14,
		LocalTime:     false,
		Compress:      false,
	},
}

type Config struct {
	ShutdownWaitTimeout time.Duration `mapstructure:"shutdown_wait_timeout"`

	Log LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	ErrorFilename string `mapstructure:"error_filename"` // ErrorFilename 日志文件名
	MaxSize       int    `mapstructure:"max_size"`       // MaxSize 触发日志切割大小 单位 MB
	MaxAge        int    `mapstructure:"max_age"`        // MaxAge 日志切割后文件保留天数
	MaxBackups    int    `mapstructure:"max_backups"`    // MaxBackups 日志切割后文件保留数量
	LocalTime     bool   `mapstructure:"local_time"`     // LocalTime 日志切割文件是否采用服务器本地时间
	Compress      bool   `mapstructure:"compress"`       // Compress 日志切割后是否对归档文件进行压缩
}
