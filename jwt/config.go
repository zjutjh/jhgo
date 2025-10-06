package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var DefaultConfig = Config{
	Secret:     "secret",
	Expiration: 30 * 24 * time.Hour,
	Issuer:     "",
	Audience:   nil,
}

type Config struct {
	Secret     string           `mapstructure:"secret"`
	Expiration time.Duration    `mapstructure:"expiration"`
	Issuer     string           `mapstructure:"issuer"`
	Audience   jwt.ClaimStrings `mapstructure:"audience"`
}
