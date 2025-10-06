package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/zjutjh/mygo/kit"
)

type JWT struct {
	conf Config
}

// New 以指定配置创建实例
func New(conf Config) *JWT {
	return &JWT{
		conf: conf,
	}
}

// GenerateToken 生成 JWT Token
func (j *JWT) GenerateToken(uid string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    j.conf.Issuer,
		Subject:   uid,
		Audience:  j.conf.Audience,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.conf.Expiration)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.conf.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析 JWT Token
func (j *JWT) ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.conf.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("解析JWT Token错误: %w", err)
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("%w: 转化JWT Token为指定Claims结构失败", kit.ErrDataFormat)
	}
	return claims, nil
}
