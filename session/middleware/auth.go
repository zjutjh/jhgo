package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/zjutjh/mygo/foundation/reply"
	"github.com/zjutjh/mygo/kit"
	"github.com/zjutjh/mygo/session"
)

// Auth Session 鉴权中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := session.GetUid(ctx)
		if err != nil {
			reply.Fail(ctx, kit.CodeNotLoggedIn)
			return
		}
	}
}
