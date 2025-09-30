package config

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/zjutjh/jhgo/kit"
)

var codeMsgMap map[kit.Code]string
var codeOnce sync.Once

// GetMessageByCode 获取业务状态码Code对应的Message
func GetMessageByCode(code kit.Code) string {
	codeOnce.Do(initCodeMsgMap)
	msg, ok := codeMsgMap[code]
	if ok {
		return msg
	}
	return fmt.Sprintf("Unknown code: %d", code)
}

func initCodeMsgMap() {
	codeMsgMap = make(map[kit.Code]string)
	for cs, msg := range CodeList() {
		c, err := strconv.ParseInt(cs, 10, 64)
		if err != nil {
			continue
		}
		codeMsgMap[kit.Code(c)] = msg
	}
}
