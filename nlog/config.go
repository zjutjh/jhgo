package nlog

import "github.com/sirupsen/logrus"

var DefaultConfig = Config{
	Filename:   "./logs/app.log",
	MaxSize:    100,
	MaxAge:     7,
	MaxBackups: 14,
	LocalTime:  false,
	Compress:   false,

	Level: logrus.InfoLevel,

	FeishuHook: FeishuHookConfig{
		Feishu: "",
		NoticeLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
	},
}

type Config struct {
	Filename   string `mapstructure:"filename"`    // Filename 日志文件路径
	MaxSize    int    `mapstructure:"max_size"`    // MaxSize 触发日志切割大小 单位 MB
	MaxAge     int    `mapstructure:"max_age"`     // MaxAge 日志切割后文件保留天数
	MaxBackups int    `mapstructure:"max_backups"` // MaxBackups 日志切割后文件保留数量
	LocalTime  bool   `mapstructure:"local_time"`  // LocalTime 日志切割文件是否采用服务器本地时间
	Compress   bool   `mapstructure:"compress"`    // Compress 日志切割后是否对归档文件进行压缩

	Level logrus.Level `mapstructure:"level"` // Level 日志实例记录等级

	FeishuHook FeishuHookConfig `mapstructure:"feishu_hook"`
}

type FeishuHookConfig struct {
	Feishu       string         `mapstructure:"feishu"`
	NoticeLevels []logrus.Level `mapstructure:"notice_levels"`
}
