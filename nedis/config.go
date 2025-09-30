package nedis

import "time"

const (
	ModeSingle = "single"
	// ModeCluster = "cluster"
	// ModeFailover = "failover"
)

var DefaultConfig = Config{
	Logger:         "",
	InfoRecordTime: 200 * time.Millisecond,
	WarnRecordTime: 1 * time.Second,

	Addrs:    []string{"localhost:6379"},
	DB:       0,
	Username: "",
	Password: "",

	Mode: ModeSingle,
}

type Config struct {
	Logger         string        `mapstructure:"logger"`
	InfoRecordTime time.Duration `mapstructure:"info_record_time"`
	WarnRecordTime time.Duration `mapstructure:"warn_record_time"`

	Addrs    []string `mapstructure:"addrs"`
	DB       int      `mapstructure:"db"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`

	Mode string `mapstructure:"mode"`
}
