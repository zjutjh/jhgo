package feishu

import (
	"time"
)

var DefaultConfig = Config{
	Enable:        false,
	NoticeWebhook: "",
	NoticeSecret:  "",

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
	Enable        bool   `mapstructure:"enable"`
	NoticeWebhook string `mapstructure:"notice_webhook"`
	NoticeSecret  string `mapstructure:"notice_secret"`

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
