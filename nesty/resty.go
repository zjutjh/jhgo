package nesty

import (
	"net"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/zjutjh/mygo/nlog"
)

// New 以指定配置创建实例
func New(conf Config) *resty.Client {
	// 选中logger
	l := nlog.Pick(conf.Log)

	// 初始化HTTP Client实例
	hc := &http.Client{
		Timeout: conf.Timeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   conf.DialContextTimeout,
				KeepAlive: conf.DialContextKeepAlive,
			}).DialContext,
			TLSHandshakeTimeout:    conf.TLSHandshakeTimeout,
			DisableKeepAlives:      conf.DisableKeepAlives,
			DisableCompression:     conf.DisableCompression,
			MaxIdleConns:           conf.MaxIdleConns,
			MaxIdleConnsPerHost:    conf.MaxIdleConnsPerHost,
			MaxConnsPerHost:        conf.MaxConnsPerHost,
			IdleConnTimeout:        conf.IdleConnTimeout,
			ResponseHeaderTimeout:  conf.ResponseHeaderTimeout,
			ExpectContinueTimeout:  conf.ExpectContinueTimeout,
			MaxResponseHeaderBytes: conf.MaxResponseHeaderBytes,
			WriteBufferSize:        conf.WriteBufferSize,
			ReadBufferSize:         conf.ReadBufferSize,
			ForceAttemptHTTP2:      conf.ForceAttemptHTTP2,
		},
	}

	// 初始化resty Client
	client := resty.NewWithClient(hc)

	// 设置重试属性
	client.SetRetryCount(conf.RetryCount)
	client.SetRetryWaitTime(conf.RetryWaitTime)
	client.SetRetryMaxWaitTime(conf.RetryMaxWaitTime)

	// 设置日志记录器
	client.SetLogger(newLogger(l))

	// 设置Hook
	client.OnBeforeRequest(onBeforeRequest())
	client.OnAfterResponse(onAfterResponse(l, conf.InfoRecordTime, conf.WarnRecordTime))
	client.OnError(onError(l))

	return client
}
