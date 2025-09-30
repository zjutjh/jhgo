package nesty

import "time"

var DefaultConfig = Config{
	Logger:         "",
	InfoRecordTime: 100 * time.Millisecond,
	WarnRecordTime: 200 * time.Millisecond,

	Timeout: 5 * time.Second,

	MaxIdleConns:        0,
	MaxIdleConnsPerHost: 200,
	MaxConnsPerHost:     500,
	IdleConnTimeout:     30 * time.Second,

	RetryCount:       0,
	RetryWaitTime:    0,
	RetryMaxWaitTime: 0,
}

type Config struct {
	Logger         string        `mapstructure:"logger"`
	InfoRecordTime time.Duration `mapstructure:"info_record_time"`
	WarnRecordTime time.Duration `mapstructure:"warn_record_time"`

	Timeout time.Duration `mapstructure:"timeout"`

	// HTTP Client Transport配置
	MaxIdleConns        int           `mapstructure:"max_idle_conns"`
	MaxIdleConnsPerHost int           `mapstructure:"max_idle_conns_per_host"`
	MaxConnsPerHost     int           `mapstructure:"max_conns_per_host"`
	IdleConnTimeout     time.Duration `mapstructure:"idle_conn_timeout"`

	// resty Retry配置
	RetryCount       int           `mapstructure:"retry_count"`
	RetryWaitTime    time.Duration `mapstructure:"retry_wait_time"`
	RetryMaxWaitTime time.Duration `mapstructure:"retry_max_wait_time"`
}
