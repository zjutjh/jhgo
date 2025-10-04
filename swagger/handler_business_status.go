package swagger

import (
	"errors"
	"slices"

	"github.com/zjutjh/mygo/kit"
)

var statusCodeMap = map[string][]kit.Code{}

var ErrBusinessStatusAlreadyRegistered = errors.New("该处理器已经注册业务状态码")

func RegisterBusinessStatusCodes(funcName string, codes []kit.Code) error {
	slices.SortFunc(codes, func(a, b kit.Code) int {
		return int(a.Code) - int(b.Code)
	})
	if v, ok := statusCodeMap[funcName]; ok {
		if slices.Equal(v, codes) {
			return nil
		}
		return ErrBusinessStatusAlreadyRegistered
	}
	statusCodeMap[funcName] = codes
	return nil
}

func MustRegisterBusinessStatusCodes(funcName string, codes []kit.Code) {
	err := RegisterBusinessStatusCodes(funcName, codes)
	if err != nil {
		panic(err)
	}
}

func getAllBusinessStatusCodes(handerNames ...string) []kit.Code {
	allCodes := map[kit.Code]struct{}{}
	for i, name := range handerNames {
		codes, ok := statusCodeMap[name]
		if !ok && i == len(handerNames)-1 {
			Output("发现未注册业务状态码的处理器[%s]\n", name)
			continue
		}
		for _, code := range codes {
			allCodes[code] = struct{}{}
		}
	}
	ans := make([]kit.Code, 0, len(allCodes))
	for code := range allCodes {
		ans = append(ans, code)
	}
	slices.SortFunc(ans, func(a, b kit.Code) int {
		return int(a.Code) - int(b.Code)
	})
	return ans
}
