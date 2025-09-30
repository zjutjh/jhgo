package nesty

import (
	"github.com/go-resty/resty/v2"
	"github.com/samber/do"
)

const (
	iocPrefix    = "_resty_:"
	defaultScope = "resty"
)

// Exist 判断scope实例是否挂载 (被Boot过) 且类型正确
func Exist(scope string) bool {
	_, err := do.InvokeNamed[*resty.Client](nil, iocPrefix+scope)
	return err == nil
}

// Pick 获取指定scope实例
func Pick(scopes ...string) *resty.Client {
	scope := defaultScope
	if len(scopes) != 0 && scopes[0] != "" {
		scope = scopes[0]
	}
	return do.MustInvokeNamed[*resty.Client](nil, iocPrefix+scope)
}
