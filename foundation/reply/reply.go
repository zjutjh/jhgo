package reply

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zjutjh/mygo/config"
	"github.com/zjutjh/mygo/kit"
)

// Response 通用标准响应
type Response struct {
	Code    kit.Code `json:"code"`
	Message string   `json:"message"`
	Data    any      `json:"data"`
}

// Success 标准成功HTTP API响应
func Success(ctx *gin.Context, data any) {
	Reply(ctx, kit.CodeOK, data)
}

// Fail 标准错误HTTP API响应
func Fail(ctx *gin.Context, code kit.Code) {
	Reply(ctx, code, nil)
}

// Reply 标准HTTP API响应
func Reply(ctx *gin.Context, code kit.Code, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: config.GetMessageByCode(code),
		Data:    data,
	})
	ctx.Abort()
}
