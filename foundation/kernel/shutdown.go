package kernel

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// ListenStop 监听结束信号并注册处理逻辑器
func ListenStop(handler func() error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err := handler()
	if err != nil {
		fmt.Fprintln(os.Stdout, "关闭处理错误:", err)
		// os.Exit(1)
	}
}
