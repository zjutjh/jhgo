package nedis

import (
	"github.com/redis/go-redis/v9"

	"github.com/zjutjh/mygo/nlog"
)

// New 以指定配置创建实例
func New(conf Config) redis.UniversalClient {
	// 选中logger
	l := nlog.Pick(conf.Logger)

	// 创建hook
	h := &hook{
		logger:         l,
		InfoRecordTime: conf.InfoRecordTime,
		WarnRecordTime: conf.WarnRecordTime,
	}

	// 初始化options
	options := &redis.UniversalOptions{
		Addrs:    conf.Addrs,
		DB:       conf.DB,
		Username: conf.Username,
		Password: conf.Password,
	}

	// 创建client并挂载hook
	var client redis.UniversalClient
	switch conf.Mode {
	// case ModeCluster:
	// 	client = redis.NewClusterClient(options.Cluster())
	// case ModeFailover:
	// 	client = redis.NewFailoverClient(options.Failover())
	case ModeSingle:
		client = redis.NewClient(options.Simple())
	default:
		client = redis.NewClient(options.Simple())
	}
	client.AddHook(h)

	return client
}
