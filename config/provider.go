package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samber/do"
	"github.com/spf13/viper"
)

const (
	iocPrefix    = "_cfg_:"
	defaultScope = "app"
)

var extMap = map[string]string{
	".yaml": "yaml",
	// ".toml": "toml",
	// ".ini":  "ini",
	// ".yml":  "yml",
	// ".json": "json",
}

// Boot 预加载指定目录下全部配置文件实例
func Boot(path string) error {
	des, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("读取开发框架配置目录[%s]错误: %w", path, err)
	}
	// 挂载配置 ({path/xx.yaml})
	for _, de := range des {
		if de.IsDir() {
			continue
		}
		ext := filepath.Ext(de.Name())
		file := de.Name()[:len(de.Name())-len(ext)]
		ext2, ok := extMap[ext]
		if !ok {
			continue
		}
		do.ProvideNamed(nil, iocPrefix+file, func(injector *do.Injector) (*viper.Viper, error) {
			v := viper.New()
			v.SetConfigName(file)
			v.SetConfigType(ext2)
			v.AddConfigPath(path)
			err = v.ReadInConfig()
			if err != nil {
				return nil, err
			}
			return v, nil
		})
	}
	return nil
}

// Exist 判断指定scope实例是否挂载 (被Boot过) 且类型正确
func Exist(scope string) bool {
	_, err := do.InvokeNamed[*viper.Viper](nil, iocPrefix+scope)
	return err == nil
}

// Pick 获取指定scope实例
func Pick(scopes ...string) *viper.Viper {
	scope := defaultScope
	if len(scopes) != 0 && scopes[0] != "" {
		scope = scopes[0]
	}
	return do.MustInvokeNamed[*viper.Viper](nil, iocPrefix+scope)
}
