package nlog

import (
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

const (
	iocPrefix    = "_log_:"
	defaultScope = "log"
)

// Exist 判断scope实例是否挂载 (被Boot过) 且类型正确
func Exist(scope string) bool {
	_, err := do.InvokeNamed[*logrus.Logger](nil, iocPrefix+scope)
	return err == nil
}

// Pick 获取指定scope实例
func Pick(scopes ...string) *logrus.Logger {
	scope := defaultScope
	if len(scopes) != 0 && scopes[0] != "" {
		scope = scopes[0]
	}
	return do.MustInvokeNamed[*logrus.Logger](nil, iocPrefix+scope)
}
