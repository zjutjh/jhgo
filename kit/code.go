package kit

var CodeOK = NewCode(0, "成功")

// 系统错误码
var (
	CodeUnknownError           = NewCode(10000, "未知错误")
	CodeThirdServiceError      = NewCode(10001, "三方服务错误")
	CodeDatabaseError          = NewCode(10002, "数据库错误")
	CodeRedisError             = NewCode(10003, "Redis错误")
	CodeMiddlewareServiceError = NewCode(10004, "中间件服务错误")
)

// 业务通用错误码
var (
	CodeNotLoggedIn        = NewCode(20000, "用户未登录")
	CodeLoginExpired       = NewCode(20001, "登录过期，请重新登录")
	CodePermissionDenied   = NewCode(20002, "用户无权限")
	CodeParameterInvalid   = NewCode(20003, "参数非法")
	CodeDataParseError     = NewCode(20004, "数据解析异常")
	CodeDataNotFound       = NewCode(20005, "数据不存在")
	CodeDataConflict       = NewCode(20006, "数据冲突")
	CodeServiceMaintenance = NewCode(20007, "系统维护中")
	CodeTooFrequently      = NewCode(20008, "操作过于频繁/未获得锁")
)

type Code struct {
	Code    int64
	Message string
}

func NewCode(code int64, message string) Code {
	return Code{
		Code:    code,
		Message: message,
	}
}
