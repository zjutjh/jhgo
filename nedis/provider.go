package nedis

import (
	"github.com/redis/go-redis/v9"
	"github.com/samber/do"
)

const (
	iocPrefix    = "_redis_:"
	defaultScope = "redis"
	configScope  = "redis"
)

// Exist 判断scope实例是否挂载 (被Boot过) 且类型正确
func Exist(scope string) bool {
	_, err := do.InvokeNamed[redis.UniversalClient](nil, iocPrefix+scope)
	return err == nil
}

// Pick 获取指定scope实例
func Pick(scopes ...string) redis.UniversalClient {
	scope := defaultScope
	if len(scopes) != 0 && scopes[0] != "" {
		scope = scopes[0]
	}
	return do.MustInvokeNamed[redis.UniversalClient](nil, iocPrefix+scope)
}
