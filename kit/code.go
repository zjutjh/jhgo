package kit

type Code int64

const CodeOK Code = 0 // CodeOK 成功

// 系统错误码
const (
	CodeUnknowkitor            Code = 10000 // 未知错误
	CodeThirdServiceError      Code = 10001 // 三方服务错误
	CodeDatabaseError          Code = 10002 // 数据库错误
	CodeRedisError             Code = 10003 // Redis错误
	CodeMiddlewareServiceError Code = 10004 // 中间件服务错误
)

// 业务通用错误码
const (
	CodeNotLogggedIn       Code = 20000 // 用户未登录
	CodeLoginExpired       Code = 20001 // 登录过期，请重新登录
	CodePermissionDenied   Code = 20002 // 用户无权限
	CodeParamterInvalid    Code = 20003 // 参数非法
	CodeDataParseError     Code = 20004 // 数据解析异常
	CodeDataNotFound       Code = 20005 // 数据不存在
	CodeDataConflict       Code = 20006 // 数据冲突
	CodeServiceMaintenance Code = 20007 // 系统维护中
	CodeTooFequently       Code = 20008 // 操作过于频繁/未获得锁
)
