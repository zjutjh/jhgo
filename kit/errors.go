package kit

import (
	"errors"
	"fmt"
)

// 数据类通用错误
var (
	ErrNotFound       = errors.New("资源不存在")
	ErrAlreadyExists  = errors.New("资源已存在")
	ErrDataUnmarshal  = errors.New("反序列化错误")
	ErrDataMarshal    = errors.New("序列化错误")
	ErrDataFormat     = errors.New("格式错误")
	ErrNoAffectedRows = errors.New("没有影响行数")
)

// 登录态通用错误
var (
	ErrNotLogged          = errors.New("未登录")
	ErrLoginStatusExpired = errors.New("登录已失效")
)

// 请求类通用错误
var (
	ErrRequestInvalidParamter = errors.New("请求参数错误")
	ErrHttpStatusCodeNotOK    = errors.New("请求响应HTTP状态码非200")
	ErrRequestBizCodeNotOK    = errors.New("请求响应业务状态码非成功")
	ErrRequestTooFrequently   = errors.New("请求过于频繁")
)

// NewRequestInvalidParamterError 创建一个具体的HTTP Status Code非成功错误
func NewHttpStatusCodeNotOKError(code int) error {
	return fmt.Errorf("%w: HTTP Status Code[%d]", ErrHttpStatusCodeNotOK, code)
}

// NewRequestBizCodeNotOKError 创建一个具体的业务状态码非成功错误
func NewRequestBizCodeNotOKError(code int) error {
	return fmt.Errorf("%w: 业务状态码[%d]", ErrRequestBizCodeNotOK, code)
}
