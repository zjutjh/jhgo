package feishu

import (
	"github.com/samber/do"
)

const (
	iocPrefix    = "_feishu_:"
	defaultScope = "feishu"
	configScope  = "feishu"
)

// Exist 判断scope实例是否挂载 (被Boot过) 且类型正确
func Exist(scope string) bool {
	_, err := do.InvokeNamed[*Feishu](nil, iocPrefix+scope)
	return err == nil
}

// Pick 获取指定scope实例
func Pick(scopes ...string) *Feishu {
	scope := defaultScope
	if len(scopes) != 0 && scopes[0] != "" {
		scope = scopes[0]
	}
	return do.MustInvokeNamed[*Feishu](nil, iocPrefix+scope)
}
