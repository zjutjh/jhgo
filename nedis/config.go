package nedis

import "time"

const (
	ModeSingle   = "single"
	ModeCluster  = "cluster"
	ModeFailover = "failover"
)

var DefaultConfig = Config{
	Log:            "",
	InfoRecordTime: 1 * time.Millisecond,
	WarnRecordTime: 10 * time.Millisecond,

	Addrs:      nil,
	ClientName: "",
	DB:         0,

	Protocol:         0,
	Username:         "",
	Password:         "",
	SentinelUsername: "",
	SentinelPassword: "",

	MaxRetries:      0,
	MinRetryBackoff: 0,
	MaxRetryBackoff: 0,

	DialTimeout:           0,
	ReadTimeout:           0,
	WriteTimeout:          0,
	ContextTimeoutEnabled: false,

	PoolFIFO:        false,
	PoolSize:        0,
	PoolTimeout:     0,
	MinIdleConns:    0,
	MaxIdleConns:    0,
	MaxActiveConns:  0,
	ConnMaxIdleTime: 0,
	ConnMaxLifetime: 0,

	MaxRedirects:   0,
	ReadOnly:       false,
	RouteByLatency: false,
	RouteRandomly:  false,

	MasterName: "",

	DisableIdentity: false,
	IdentitySuffix:  "",
	UnstableResp3:   false,

	Mode: ModeSingle,
}

type Config struct {
	Log            string        `mapstructure:"logger"`
	InfoRecordTime time.Duration `mapstructure:"info_record_time"`
	WarnRecordTime time.Duration `mapstructure:"warn_record_time"`

	Addrs      []string `mapstructure:"addrs"`
	ClientName string   `mapstructure:"client_name"`
	DB         int      `mapstructure:"db"`

	Protocol         int    `mapstructure:"protocol"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	SentinelUsername string `mapstructure:"sentinel_username"`
	SentinelPassword string `mapstructure:"sentinel_password"`

	MaxRetries      int           `mapstructure:"max_retries"`
	MinRetryBackoff time.Duration `mapstructure:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `mapstructure:"max_retry_backoff"`

	DialTimeout           time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout           time.Duration `mapstructure:"read_timeout"`
	WriteTimeout          time.Duration `mapstructure:"write_timeout"`
	ContextTimeoutEnabled bool          `mapstructure:"context_timeout_enabled"`

	PoolFIFO        bool          `mapstructure:"pool_fifo"`
	PoolSize        int           `mapstructure:"pool_size"`
	PoolTimeout     time.Duration `mapstructure:"pool_timeout"`
	MinIdleConns    int           `mapstructure:"min_idle_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxActiveConns  int           `mapstructure:"max_active_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`

	MaxRedirects   int  `mapstructure:"max_redirects"`
	ReadOnly       bool `mapstructure:"read_only"`
	RouteByLatency bool `mapstructure:"route_by_latency"`
	RouteRandomly  bool `mapstructure:"route_randomly"`

	MasterName string `mapstructure:"master_name"`

	DisableIdentity bool   `mapstructure:"disable_identity"`
	IdentitySuffix  string `mapstructure:"identity_suffix"`
	UnstableResp3   bool   `mapstructure:"unstable_resp3"`

	Mode string `mapstructure:"mode"`
}
