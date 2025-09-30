package ndb

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/samber/do"

	"github.com/zjutjh/jhgo/config"
	"github.com/zjutjh/jhgo/kit"
)

// Boot 预加载默认实例 同时加载指定实例列表
func Boot(scopes ...string) func() error {
	return func() error {
		if err := provide(defaultScope); err != nil {
			return fmt.Errorf("加载资源[%s]错误: %w", defaultScope, err)
		}
		for _, scope := range scopes {
			if err := provide(scope); err != nil {
				return fmt.Errorf("加载资源[%s]错误: %w", scope, err)
			}
		}
		return nil
	}
}

// provide 提供指定scope实例
func provide(scope string) error {
	// 获取配置
	conf, err := getConf(scope)
	if err != nil {
		return err
	}

	// 初始化实例
	instance, err := New(conf)
	if err != nil {
		return fmt.Errorf("初始化DB实例错误: %w", err)
	}

	// 挂载实例
	do.ProvideNamedValue(nil, iocPrefix+scope, instance)

	return nil
}

// getConf 获取配置
func getConf(scope string) (conf Config, err error) {
	// 先尝试核心配置 {configScope}.yaml
	if config.Exist(configScope) {
		core := config.Pick(configScope)
		if core.IsSet(scope) {
			// 初始化默认配置
			conf, err = defaultConfig()
			if err != nil {
				return
			}
			// 解析 {configScope}.yaml[{scope}]
			err = core.UnmarshalKey(scope, &conf)
			if err != nil {
				return conf, fmt.Errorf("%w: 解析%s.yaml [%s=>%s]错误: %w", kit.ErrDataUnmarshal, configScope, scope, scope, err)
			}
			return conf, nil
		}
	}

	// 再尝试应用配置 app.yaml
	app := config.Pick()
	if !app.IsSet(scope) {
		return conf, fmt.Errorf("%w: 配置app.yaml[%s]不存在", kit.ErrNotFound, scope)
	}
	// 初始化默认配置
	conf, err = defaultConfig()
	if err != nil {
		return
	}
	// 解析 app.yaml[{scope}]
	err = app.UnmarshalKey(scope, &conf)
	if err != nil {
		return conf, fmt.Errorf("%w: 解析app.yaml[%s]错误: %w", kit.ErrDataUnmarshal, scope, err)
	}
	return conf, nil
}

// defaultConfig 获取默认配置
func defaultConfig() (conf Config, err error) {
	err = copier.CopyWithOption(&conf, &DefaultConfig, copier.Option{DeepCopy: true})
	return conf, err
}
