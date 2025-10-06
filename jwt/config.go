package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var DefaultConfig = Config{
	Secret:     "secret",
	Expiration: 30 * 24 * time.Hour,
	Issuer:     "",
	Subject:    "",
	Audience:   nil,

	Log: "",
}

type Config struct {
	Secret     string           `mapstructure:"secret"`
	Expiration time.Duration    `mapstructure:"expiration"`
	Issuer     string           `mapstructure:"issuer"`
	Subject    string           `mapstructure:"subject"`
	Audience   jwt.ClaimStrings `mapstructure:"audience"`

	Log string `mapstructure:"log"`
}
