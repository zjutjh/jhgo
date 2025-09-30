package config

import (
	"fmt"

	"github.com/zjutjh/jhgo/kit"
)

// 应用环境
const (
	AppEnvProd = "prod"
	AppEnvTest = "test"
	AppEnvDev  = "dev"
)

// AppConfig 应用基础配置结构
type AppConfig struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
}

// GetAppConf 获取应用基础配置
func GetAppConf() (AppConfig, error) {
	ac := AppConfig{}
	err := Pick().UnmarshalKey(defaultScope, &ac)
	if err != nil {
		return ac, fmt.Errorf("%w: 解析应用基础配置错误: %w", kit.ErrDataUnmarshal, err)
	}
	return ac, nil
}

// AppName 获取配置的应用Name
func AppName() string {
	return Pick().GetString("app.name")
}

// AppEnv 获取配置的应用Env
func AppEnv() string {
	env := Pick().GetString("app.env")
	if env == "" {
		env = AppEnvDev
	}
	if env != AppEnvTest && env != AppEnvProd {
		env = AppEnvDev
	}
	return env
}

// CodeList 获取配置的Code列表
func CodeList() map[string]string {
	return Pick().GetStringMapString("code")
}
