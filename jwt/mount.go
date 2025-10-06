package jwt

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/zjutjh/mygo/kit"
)

const MountKey = "_jwt_uid_"

// MountUid 挂载uid至上下文
func MountUid(ctx *gin.Context, uid string) {
	ctx.Set(MountKey, uid)
}

// GetUid 获取上下文中挂载的uid
// 注意：该函数需在 MountUid 函数进行挂载后使用
func GetUid(ctx *gin.Context) (string, error) {
	v, ok := ctx.Get(MountKey)
	if !ok {
		return "", fmt.Errorf("%w: 当前上下文未挂载[%s]", kit.ErrNotFound, MountKey)
	}
	uid, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("%w: 当前上下文挂载[%s]类型错误", kit.ErrDataFormat, MountKey)
	}
	return uid, nil
}
