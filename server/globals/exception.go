package globals

import (
	"encoding/json"
	"fmt"

	httpError "gin-server/server/configs/error"
)

type HttpException struct {
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Data    any      `json:"data"`
	Source  []string `json:"source"`
}

func (e *HttpException) Error() string {
	jsons, _ := json.Marshal(e)

	return string(jsons)
}

/**
 * NewException()
 * NewException("error message") // 第一个参数为错误信息描述，string类型
 * NewException("error message", "INNER_SERVER_ERROR") // 第二个参数为错误码，string类型
 * NewException("error message", "INNER_SERVER_ERROR", "any data") // 第三个参数为任意数据，可以是任意类型
 * NewException("error message", "INNER_SERVER_ERROR", "any data", ["xxx"]) // 第四个参数为错误来源，[]string类型
 */
func NewException(args ...any) *HttpException {
	e := &HttpException{
		Message: "internal server error",
		Code:    "INTERNAL_SERVER_ERROR",
		Data:    nil,
		Source:  []string{},
	}

	for i, arg := range args {
		switch i {
		case 0: // message
			if msg, ok := arg.(string); ok {
				e.Message = msg
			} else {
				e.Message = fmt.Sprintf("%v", arg)
			}
		case 1: // code
			if code, ok := arg.(string); ok && httpError.ErrorCodeMap[code] != 0 {
				e.Code = code
			}
		case 2: // data
			e.Data = arg
		case 3: // source
			if sources, ok := arg.([]string); ok {
				e.Source = sources
			}
		}
	}
	hasSource := false

	for _, s := range e.Source {
		if s == SERVER_NAME {
			hasSource = true
			break
		}
	}

	if !hasSource {
		e.Source = append(e.Source, SERVER_NAME)
	}

	if httpError.ErrorCodeMap[e.Code] == 0 {
		e.Code = "INTERNAL_SERVER_ERROR"
	}

	return e
}
