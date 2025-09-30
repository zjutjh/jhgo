package nesty

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func onBeforeRequest() func(client *resty.Client, request *resty.Request) error {
	return func(client *resty.Client, request *resty.Request) error {
		// 设置X-Request-ID
		if request.Header.Get("X-Request-ID") == "" {
			ctx := request.Context()
			if ctx2, ok := ctx.(*gin.Context); ok {
				if rid := ctx2.GetString("X-Request-ID"); rid != "" {
					request.SetHeader("X-Request-ID", rid)
				}
			}
		}
		return nil
	}
}

func onAfterResponse(logger *logrus.Logger, infoRecordTime, warnRecordTime time.Duration) func(client *resty.Client, response *resty.Response) error {
	return func(client *resty.Client, response *resty.Response) error {
		if logger == nil {
			return nil
		}

		// 5xx或4xx错误
		request := response.Request
		code := response.StatusCode()
		if code >= 400 && code <= 599 {
			logger.WithContext(request.Context()).WithFields(logrus.Fields{
				"method":           request.Method,
				"url":              request.URL,
				"req_body":         request.Body,
				"resp_status_code": code,
				"resp_body":        string(response.Body()),
				"cost":             response.Time().String(),
			}).Error("发送HTTP请求失败")
			return nil
		}

		// 耗时日志记录
		cost := response.Time()
		if warnRecordTime >= 0 && cost >= warnRecordTime {
			logger.WithContext(request.Context()).WithFields(logrus.Fields{
				"method":           request.Method,
				"url":              request.URL,
				"req_body":         request.Body,
				"resp_status_code": code,
				"resp_body":        string(response.Body()),
				"cost":             cost.String(),
				"threshold":        warnRecordTime.String(),
			}).Warn("请求HTTP接口时长超过期望")
		} else if infoRecordTime >= 0 && cost >= infoRecordTime {
			logger.WithContext(request.Context()).WithFields(logrus.Fields{
				"method":           request.Method,
				"url":              request.URL,
				"req_body":         request.Body,
				"resp_status_code": code,
				"resp_body":        string(response.Body()),
				"cost":             cost.String(),
			}).Info("请求HTTP接口时长记录")
		}

		return nil
	}
}

func onError(logger *logrus.Logger) func(request *resty.Request, err error) {
	return func(request *resty.Request, err error) {
		if logger == nil {
			return
		}

		// 基础信息
		entry := logger.WithContext(request.Context()).WithError(err).WithFields(logrus.Fields{
			"method":   request.Method,
			"url":      request.URL,
			"req_body": request.Body,
		})

		// 拿到response的error
		var v *resty.ResponseError
		if errors.As(err, &v) {
			entry = entry.WithFields(logrus.Fields{
				"resp_status_code": v.Response.StatusCode(),
				"resp_body":        string(v.Response.Body()),
				"cost":             v.Response.Time().String(),
			})
		}

		// 记录日志
		entry.Error("发送HTTP请求失败")
	}
}
