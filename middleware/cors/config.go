package cors

import (
	"time"
)

// DefaultConfig 默认配置
var DefaultConfig = Config{
	AllowAllOrigins:           false,
	AllowOrigins:              nil,
	AllowMethods:              []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	AllowPrivateNetwork:       false,
	AllowHeaders:              []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	AllowCredentials:          true,
	ExposeHeaders:             []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "X-Request-Id"},
	MaxAge:                    12 * time.Hour,
	AllowWildcard:             true,
	AllowBrowserExtensions:    false,
	CustomSchemas:             nil,
	AllowWebSockets:           true,
	AllowFiles:                false,
	OptionsResponseStatusCode: 0,
}

// Config 对应cors.Config中的同名配置项
type Config struct {
	AllowAllOrigins           bool          `mapstructure:"allow_all_origins"`
	AllowOrigins              []string      `mapstructure:"allow_origins"`
	AllowMethods              []string      `mapstructure:"allow_methods"`
	AllowPrivateNetwork       bool          `mapstructure:"allow_private_network"`
	AllowHeaders              []string      `mapstructure:"allow_headers"`
	AllowCredentials          bool          `mapstructure:"allow_credentials"`
	ExposeHeaders             []string      `mapstructure:"expose_headers"`
	MaxAge                    time.Duration `mapstructure:"max_age"`
	AllowWildcard             bool          `mapstructure:"allow_wildcard"`
	AllowBrowserExtensions    bool          `mapstructure:"allow_browser_extensions"`
	CustomSchemas             []string      `mapstructure:"custom_schemas"`
	AllowWebSockets           bool          `mapstructure:"allow_web_sockets"`
	AllowFiles                bool          `mapstructure:"allow_files"`
	OptionsResponseStatusCode int           `mapstructure:"options_response_status_code"`
}
