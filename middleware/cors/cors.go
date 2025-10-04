package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/zjutjh/mygo/config"
	"github.com/zjutjh/mygo/kit"
)

const defaultConfigKey = "mid_cors"

// Pick 获取指定实例
func Pick(keys ...string) gin.HandlerFunc {
	key := defaultConfigKey
	if len(keys) != 0 && keys[0] != "" {
		key = keys[0]
	}
	conf := Config{}
	err := copier.Copy(&conf, DefaultConfig)
	if err != nil {
		panic(err)
	}
	app := config.Pick()
	if !app.IsSet(key) {
		panic(kit.ErrNotFound)
	}
	err = app.UnmarshalKey(key, &conf)
	if err != nil {
		panic(err)
	}
	return cors.New(cors.Config{
		AllowAllOrigins:           conf.AllowAllOrigins,
		AllowOrigins:              conf.AllowOrigins,
		AllowMethods:              conf.AllowMethods,
		AllowPrivateNetwork:       conf.AllowPrivateNetwork,
		AllowHeaders:              conf.AllowHeaders,
		AllowCredentials:          conf.AllowCredentials,
		ExposeHeaders:             conf.ExposeHeaders,
		MaxAge:                    conf.MaxAge,
		AllowWildcard:             conf.AllowWildcard,
		AllowBrowserExtensions:    conf.AllowBrowserExtensions,
		CustomSchemas:             conf.CustomSchemas,
		AllowWebSockets:           conf.AllowWebSockets,
		AllowFiles:                conf.AllowFiles,
		OptionsResponseStatusCode: conf.OptionsResponseStatusCode,
	})
}
