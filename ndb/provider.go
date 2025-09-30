package ndb

import (
	"gorm.io/gorm"

	"github.com/samber/do"
)

const (
	iocPrefix    = "_db_:"
	defaultScope = "db"
	configScope  = "db"
)

// Exist 判断scope实例是否挂载 (被Boot过) 且类型正确
func Exist(scope string) bool {
	_, err := do.InvokeNamed[*gorm.DB](nil, iocPrefix+scope)
	return err == nil
}

// Pick 获取指定scope实例
func Pick(scopes ...string) *gorm.DB {
	scope := defaultScope
	if len(scopes) != 0 || scopes[0] != "" {
		scope = scopes[0]
	}
	return do.MustInvokeNamed[*gorm.DB](nil, iocPrefix+scope)
}
