package nesty

import "time"

var DefaultConfig = Config{
	Log:            "",
	InfoRecordTime: 100 * time.Millisecond,
	WarnRecordTime: 200 * time.Millisecond,

	Timeout: 5 * time.Second,

	TLSHandshakeTimeout:    0,
	DisableKeepAlives:      false,
	DisableCompression:     false,
	MaxIdleConns:           0,
	MaxIdleConnsPerHost:    200,
	MaxConnsPerHost:        500,
	IdleConnTimeout:        30 * time.Second,
	ResponseHeaderTimeout:  0,
	ExpectContinueTimeout:  1 * time.Second,
	MaxResponseHeaderBytes: 0,
	WriteBufferSize:        0,
	ReadBufferSize:         0,
	ForceAttemptHTTP2:      true,
	DialContextTimeout:     30 * time.Second,
	DialContextKeepAlive:   30 * time.Second,

	RetryCount:       0,
	RetryWaitTime:    0,
	RetryMaxWaitTime: 0,
}

type Config struct {
	Log            string        `mapstructure:"logger"`
	InfoRecordTime time.Duration `mapstructure:"info_record_time"`
	WarnRecordTime time.Duration `mapstructure:"warn_record_time"`

	Timeout time.Duration `mapstructure:"timeout"`

	// HTTP Client Transport配置
	TLSHandshakeTimeout    time.Duration `mapstructure:"tls_handshake_timeout"`
	DisableKeepAlives      bool          `mapstructure:"disable_keep_alives"`
	DisableCompression     bool          `mapstructure:"disable_compression"`
	MaxIdleConns           int           `mapstructure:"max_idle_conns"`
	MaxIdleConnsPerHost    int           `mapstructure:"max_idle_conns_per_host"`
	MaxConnsPerHost        int           `mapstructure:"max_conns_per_host"`
	IdleConnTimeout        time.Duration `mapstructure:"idle_conn_timeout"`
	ResponseHeaderTimeout  time.Duration `mapstructure:"response_header_timeout"`
	ExpectContinueTimeout  time.Duration `mapstructure:"expect_continue_timeout"`
	MaxResponseHeaderBytes int64         `mapstructure:"max_response_header_bytes"`
	WriteBufferSize        int           `mapstructure:"write_buffer_size"`
	ReadBufferSize         int           `mapstructure:"read_buffer_size"`
	ForceAttemptHTTP2      bool          `mapstructure:"force_attempt_http2"`
	DialContextTimeout     time.Duration `mapstructure:"dial_context_timeout"`
	DialContextKeepAlive   time.Duration `mapstructure:"dial_context_keep_alive"`

	// resty Retry配置
	RetryCount       int           `mapstructure:"retry_count"`
	RetryWaitTime    time.Duration `mapstructure:"retry_wait_time"`
	RetryMaxWaitTime time.Duration `mapstructure:"retry_max_wait_time"`
}
