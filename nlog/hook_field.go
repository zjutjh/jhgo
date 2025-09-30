package nlog

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type hookField struct {
	app string
}

func (h *hookField) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *hookField) Fire(entry *logrus.Entry) error {
	// 业务字段初始化
	bizFields := entry.Data
	entry.Data = logrus.Fields{}

	// 全局字段处理
	entry.Data["app"] = h.app
	entry.Data["time"] = entry.Time.UnixMilli()

	// error字段处理
	if _, exist := bizFields[logrus.ErrorKey]; exist {
		entry.Data[logrus.ErrorKey] = bizFields[logrus.ErrorKey]
		delete(bizFields, logrus.ErrorKey)
	}

	// 业务字段处理
	entry.Data["body"] = bizFields

	// 作用域字段处理
	if entry.Context != nil {
		if ctx, ok := entry.Context.(*gin.Context); ok {
			entry.Data["client_ip"] = ctx.ClientIP()
			entry.Data["uri"] = ctx.Request.Host + ctx.Request.RequestURI
			entry.Data["method"] = ctx.Request.Method
			if rid := ctx.GetHeader("X-Request-ID"); rid != "" {
				entry.Data["request_id"] = rid
			}
		}
	}

	return nil
}
