package swagger

import (
	"errors"
	"reflect"
	"runtime"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/zjutjh/mygo/swagger/internal/utils"
)

type authSchemeRegistry struct {
	registry map[string]*SecurityScheme
	keyMap   map[string][]string
}

var (
	ErrSecuritySchemeExisted      = errors.New("security scheme existed")
	ErrSecurityMappingKeyNotFound = errors.New("security mapping key not found")
	ErrSchemeNameEmpty            = errors.New("scheme name empty")
	ErrSchemeInUse                = errors.New("scheme in use")
)

type errSecuritySchemeNotFound string

var _ error = errSecuritySchemeNotFound("")

func (e errSecuritySchemeNotFound) Error() string {
	return "security scheme not found: " + string(e)
}

type securitySchemeInfo struct {
	key    string
	scheme SecurityScheme
}

func (a *authSchemeRegistry) Add(schemeName string, scheme *SecurityScheme) error {
	if schemeName == "" {
		return ErrSchemeNameEmpty
	}
	if v, ok := a.registry[schemeName]; ok {
		if *v == *scheme {
			return nil
		}
		return ErrSecuritySchemeExisted
	}
	copy := *scheme
	a.registry[schemeName] = &copy
	return nil
}

// Map 将身份验证方案名称映射到指定的 key 上。
// 如果 schemeNames 为空，则删除该 key 的所有映射关系。
// 如果 schemeNames 中有空字符串，则返回 ErrSchemeNameEmpty。
// 如果 schemeNames 中有未注册的方案，则返回 errSecuritySchemeNotFound。
// 如果 key 已经存在，则覆盖原有的映射关系。
func (a *authSchemeRegistry) Map(key string, schemeNames ...string) error {
	if len(schemeNames) == 0 {
		delete(a.keyMap, key)
		return nil
	}
	for _, schemeName := range schemeNames {
		if schemeName == "" {
			return ErrSchemeNameEmpty
		}
		if _, ok := a.registry[schemeName]; !ok {
			return errSecuritySchemeNotFound(schemeName)
		}
	}
	a.keyMap[key] = schemeNames
	return nil
}

func (a *authSchemeRegistry) Get(schemeName string) (*SecurityScheme, error) {
	scheme, ok := a.registry[schemeName]
	if ok {
		return scheme, nil
	}
	return nil, errSecuritySchemeNotFound(schemeName)
}

func (a *authSchemeRegistry) GetFromKey(key string) ([]string, []*SecurityScheme, error) {
	if v, ok := authSechemaInfo.keyMap[key]; ok {
		schemeNames := make([]string, 0, len(v))
		schemes := make([]*SecurityScheme, 0, len(v))
		for _, schemeName := range v {
			scheme, err := a.Get(schemeName)
			if err != nil {
				return nil, nil, err
			}
			schemeNames = append(schemeNames, schemeName)
			schemes = append(schemes, scheme)
		}
		return schemeNames, schemes, nil
	}
	return nil, nil, ErrSecurityMappingKeyNotFound
}

func newAuthSchemeInfo() *authSchemeRegistry {
	return &authSchemeRegistry{
		registry: make(map[string]*SecurityScheme),
		keyMap:   make(map[string][]string),
	}
}

var authSechemaInfo = newAuthSchemeInfo()

// UnregisterAuthScheme 取消注册身份验证方案。
// 如果该方案仍在使用，则返回 ErrSchemeInUse
func UnregisterAuthScheme(schemeName string) error {
	if _, ok := authSechemaInfo.registry[schemeName]; !ok {
		return nil
	}
	// check reference
	for _, v := range authSechemaInfo.keyMap {
		if slices.Contains(v, schemeName) {
			return ErrSchemeInUse
		}
	}
	delete(authSechemaInfo.registry, schemeName)
	return nil
}

// RegisterAuthScheme 注册身份验证方案
//
// 如果已注册同名方案，若方案内容相同，则直接返回 nil；若不同，则返回 ErrSecuritySchemeExisted
func RegisterAuthScheme(schemeName string, scheme *SecurityScheme) error {
	return authSechemaInfo.Add(schemeName, scheme)
}

// MustRegisterAuthScheme 必须注册身份验证方案，若注册失败，则 panic
func MustRegisterAuthScheme(schemeName string, scheme *SecurityScheme) {
	if err := RegisterAuthScheme(schemeName, scheme); err != nil {
		panic(err)
	}
}

// RegisterMidAuthScheme 为中间件注册特定身份验证方案，所有的方案必须已经被注册。
// 若未注册该方案，则返回 ErrSecuritySchemeNotFound。
// 若方案为空，则注销该中间件的所有身份验证方案。
func RegisterMidAuthScheme(middlewareFunc gin.HandlerFunc, schemeNames ...string) error {
	schemeNames = utils.Dedup(schemeNames)
	funcName := runtime.FuncForPC(reflect.ValueOf(middlewareFunc).Pointer()).Name()
	return authSechemaInfo.Map(funcName, schemeNames...)
}

// MustRegisterMidAuthScheme 必须为中间件注册特定身份验证方案，所有的方案必须已经被注册，若注册失败，则 panic。
// 若方案为空，则注销该中间件的所有身份验证方案。
func MustRegisterMidAuthScheme(middlewareFunc gin.HandlerFunc, schemeNames ...string) {
	if err := RegisterMidAuthScheme(middlewareFunc, schemeNames...); err != nil {
		panic(err)
	}
}

func getMidAuthScheme(funcName string) []*securitySchemeInfo {
	k, v, _ := authSechemaInfo.GetFromKey(funcName)
	if len(v) == 0 {
		return nil
	}
	schemeInfos := make([]*securitySchemeInfo, 0, len(v))
	for i := range v {
		schemeInfos = append(schemeInfos, &securitySchemeInfo{
			key:    k[i],
			scheme: *v[i],
		})
	}
	return schemeInfos
}
