package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Uid string // 用户唯一标识

	jwt.RegisteredClaims
}
