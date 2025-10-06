package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/zjutjh/mygo/kit"
)

type result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Feishu struct {
	conf   Config
	client *resty.Client
}

// New 以指定配置创建实例
func New(conf Config) *Feishu {
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

	return &Feishu{
		conf:   conf,
		client: client,
	}
}

// Send 发送飞书Bot消息 https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot
// 参数:
// title: 消息标题
// message: 消息内容
func (f *Feishu) Send(title, message string) error {
	// 未启用
	if !f.conf.Enable {
		return nil
	}

	// 参数检查
	if title == "" {
		return fmt.Errorf("%w: 发送飞书消息时必须指定消息标题", kit.ErrRequestInvalidParamter)
	}
	if message == "" {
		return fmt.Errorf("%w: 发送飞书消息时必须指定消息内容", kit.ErrRequestInvalidParamter)
	}

	// 计算签名
	timestamp := time.Now().Unix()
	sign, err := f.genSign(f.conf.NoticeSecret, timestamp)
	if err != nil {
		return fmt.Errorf("签名计算发生错误: %w", err)
	}

	params := map[string]any{
		"timestamp": timestamp,
		"sign":      sign,
		"msg_type":  "post",
		"content": map[string]any{
			"post": map[string]any{
				"zh_cn": map[string]any{
					"title": title,
					"content": [][]map[string]string{
						{
							{
								"tag":  "text",
								"text": message,
							},
						},
					},
				},
			},
		},
	}

	// 发送消息
	res := result{Code: -1}
	resp, err := f.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&res).
		Post(f.conf.NoticeWebhook)
	if err != nil {
		return fmt.Errorf("发送飞书Bot消息请求错误: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("发送飞书Bot消息状态码错误: %w", err)
	}
	if res.Code != 0 {
		return fmt.Errorf("发送飞书Bot消息业务码错误: %v", res)
	}

	return nil
}

// genSign 计算签名 https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot?lang=zh-CN#3c6592d6
func (f *Feishu) genSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
