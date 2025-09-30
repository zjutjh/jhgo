package kernel

import (
	"errors"
	"fmt"
	"os"

	"github.com/zjutjh/jhgo/config"
)

type BootList []func() error

func Bootstrap(confPath string, bootRegister func() BootList) {
	// 加载配置
	if err := config.Boot(confPath); err != nil {
		fmt.Fprintln(os.Stdout, "引导加载配置错误:", err)
		os.Exit(1)
	}

	// 检查配置
	if err := checkConfig(); err != nil {
		fmt.Fprintln(os.Stdout, "引导配置发现错误:", err)
		os.Exit(1)
	}

	// 引导与加载资源
	bs := bootRegister()
	for _, boot := range bs {
		if err := boot(); err != nil {
			fmt.Fprintln(os.Stdout, "引导加载资源错误:", err)
			os.Exit(1)
		}
	}
}

func checkConfig() error {
	// 检测app.yaml是否正确
	config.Pick()
	if config.AppName() == "" {
		return errors.New("未配置应用Name, 请在app.yaml[app.name]中配置")
	}
	return nil
}
