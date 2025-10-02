package nlog

import (
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/zjutjh/mygo/feishu"
)

type FeishuHook struct {
	feishu *feishu.Feishu
	levels []logrus.Level
}

func newFeishuHook(feishu *feishu.Feishu, levels []logrus.Level) *FeishuHook {
	return &FeishuHook{
		feishu: feishu,
		levels: levels,
	}
}

func (f *FeishuHook) Levels() []logrus.Level {
	return f.levels
}

func (f *FeishuHook) Fire(entry *logrus.Entry) error {
	// 组装告警信息
	data := entry.Data

	dataContent := ""
	if method, ok := data["method"]; ok {
		if uri, ok := data["uri"]; ok {
			dataContent = fmt.Sprintf("%s\n%-16s[%s] %s", dataContent, "API:", method, uri)
		}
	}
	if ip, ok := data["client_ip"].(string); ok {
		dataContent = fmt.Sprintf("%s\n%-15s%s", dataContent, "ClientIP:", ip)
	}
	if id, ok := data["request_id"].(string); ok {
		dataContent = fmt.Sprintf("%s\n%-12s%s", dataContent, "RequestID:", id)
	}
	if err, ok := data[logrus.ErrorKey].(error); ok {
		dataContent = fmt.Sprintf("%s\n%-16s%s", dataContent, "Error:", err.Error())
	}
	if body, ok := data["body"].(logrus.Fields); ok {
		dataContent = fmt.Sprintf("%s\n%s", dataContent, "Body:\t(")
		for k, v := range body {
			dataContent = fmt.Sprintf("%s\n\t%-10s%#v", dataContent, k+":", v)
		}
		dataContent = fmt.Sprintf("%s\n)", dataContent)
	}

	message := fmt.Sprintf("%-15s%s\n%-13s[%s] %s", "Time:", entry.Time.Format(time.DateTime), "Message:", entry.Level.String(), entry.Message)
	if dataContent != "" {
		message = fmt.Sprintf("%s%s", message, dataContent)
	}

	// 发送报警
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("请求飞书Bot发送报警发生了panic", err)
			}
		}()
		title := "应用告警"
		if app, ok := data["app"].(string); ok {
			title = fmt.Sprintf("[%s]%s", app, title)
		}
		f.feishu.Send(title, message)
	}()

	return nil
}
