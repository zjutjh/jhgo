package ndb

import (
	"time"

	"gorm.io/gorm/logger"
)

var DefaultConfig = Config{
	Host:     "localhost",
	Port:     3306,
	Database: "",
	Username: "",
	Password: "",

	OpenLogger:                true,
	SlowThreshold:             1 * time.Second,
	Colorful:                  false,
	IgnoreRecordNotFoundError: true,
	ParameterizedQueries:      false,
	LogLevel:                  logger.Warn,

	MaxIdleConns:    100,
	MaxOpenConns:    200,
	ConnMaxLifetime: 5 * time.Minute,
	ConnMaxIdleTime: 1 * time.Minute,
}

type Config struct {
	// 基础系列
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`

	// gorm logger系列
	OpenLogger                bool            `mapstructure:"open_logger"`
	SlowThreshold             time.Duration   `mapstructure:"slow_threshold"`
	Colorful                  bool            `mapstructure:"colorful"`
	IgnoreRecordNotFoundError bool            `mapstructure:"ignore_record_not_found_error"`
	ParameterizedQueries      bool            `mapstructure:"parameterized_queries"`
	LogLevel                  logger.LogLevel `mapstructure:"log_level"`

	// sql系列
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}
