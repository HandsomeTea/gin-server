package response

import (
	"encoding/json"
	"net/http"

	httpError "gin-server/server/configs/error"
	logger "gin-server/server/configs/logger"

	"github.com/gin-gonic/gin"
)

type response struct {
	ctx *gin.Context
}

func Ctx(ctx *gin.Context) *response {
	return &response{ctx: ctx}
}

func (response *response) Success(data ...any) {
	if len(response.ctx.Errors) > 0 {
		return
	}
	var result any = data

	if len(data) == 1 {
		result = data[0]
	}
	if len(data) == 0 {
		result = gin.H{}
	}
	responseDataJson, _ := json.MarshalIndent(result, "", "    ")
	logger.TraceLog.Info("[http-response] " + response.ctx.Request.Method + ": " + response.ctx.Request.URL.RequestURI() + " => " + string(responseDataJson))

	response.ctx.JSON(http.StatusOK, result)
	response.ctx.Abort()
}

type httpException struct {
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Data    any      `json:"data"`
	Source  []string `json:"source"`
}

/** 参数只取前四个，其余参数将被忽略，所有参数均为可选参数，必须按顺序传入。示例如下：
 * response.Failed()
 * response.Failed("error message") // 第一个参数为错误信息描述，string类型
 * response.Failed("error message", "INNER_SERVER_ERROR") // 第二个参数为错误码，string类型
 * response.Failed("error message", "INNER_SERVER_ERROR", "any data") // 第三个参数为任意数据，可以是任意类型
 * response.Failed("error message", "INNER_SERVER_ERROR", "any data", ["xxx"]) // 第四个参数为错误来源，[]string类型
 */
func (response *response) Failed(data ...any) {
	if len(response.ctx.Errors) > 0 {
		return
	}
	errorException := httpException{
		Message: "internal server error",
		Code:    "INTERNAL_SERVER_ERROR",
		Data:    nil,
		Source:  []string{},
	}
	if len(data) >= 1 {
		message, isString := data[0].(string)

		if isString {
			errorException.Message = message
		}
	}

	if len(data) >= 2 {
		code, isString := data[1].(string)

		if isString && httpError.ErrorCodeMap[code] != 0 {
			errorException.Code = code
		}
	}

	if len(data) >= 3 {
		errorException.Data = data[2]
	}

	if len(data) >= 4 {
		source, isArray := data[3].([]string)

		if isArray {
			errorException.Source = source
		}
	}

	serverName := "gin-server"
	sourceIncludeServerName := false

	for _, source := range errorException.Source {
		if source == serverName {
			sourceIncludeServerName = true
		}
	}
	if !sourceIncludeServerName {
		errorException.Source = append(errorException.Source, serverName)
	}

	if httpError.ErrorCodeMap[errorException.Code] == 0 {
		errorException.Code = "INTERNAL_SERVER_ERROR"
	}

	status := httpError.ErrorCodeMap[errorException.Code]

	responseDataJson, _ := json.MarshalIndent(errorException, "", "    ")
	logger.TraceLog.Info("[http-response] " + response.ctx.Request.Method + ": " + response.ctx.Request.URL.RequestURI() + " => " + string(responseDataJson))
	response.ctx.JSON(status, errorException)
	response.ctx.Abort()
}
