package response

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	"gin-server/server/configs/env"
	httpError "gin-server/server/configs/error"
	"gin-server/server/configs/logger"
	"gin-server/server/globals"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	responseDataJson, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		panic(globals.NewException(err.Error()))
	}
	logMsg := "[http-response] " + response.ctx.Request.Method + ": " + response.ctx.Request.URL.RequestURI() + " => " + string(responseDataJson)
	otelEnabled, _ := env.GetEnv("OTEL_ENABLED")

	if otelEnabled == "yes" {
		logger.TraceLog.Info(logMsg, zap.Any("ctx", response.ctx.Request.Context()))
	} else {
		logger.TraceLog.Info(logMsg)
	}
	response.ctx.JSON(http.StatusOK, result)
}

/**
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
	errorException := globals.NewException(data...)
	status := httpError.ErrorCodeMap[errorException.Code]
	errorLogMsg := "[http-response] " + response.ctx.Request.Method + ": " + response.ctx.Request.URL.RequestURI() + " => "
	otelEnabled, _ := env.GetEnv("OTEL_ENABLED")
	responseDataJson, err := json.MarshalIndent(errorException, "", "    ")

	if err != nil {
		message := err.Error()

		logger.SystemLog.Error(message + "\n" + string(debug.Stack()))
		errorException = globals.NewException(message)
		jsonData, _ := json.MarshalIndent(errorException, "", "    ")
		errorLogMsg = errorLogMsg + string(jsonData)

		if otelEnabled == "yes" {
			logger.TraceLog.Error(errorLogMsg, zap.Any("ctx", response.ctx.Request.Context()))
		} else {
			logger.TraceLog.Error(errorLogMsg)
		}

		response.ctx.JSON(status, errorException)
		return
	}

	errorLogMsg = errorLogMsg + string(responseDataJson)
	if otelEnabled == "yes" {
		logger.TraceLog.Error(errorLogMsg, zap.Any("ctx", response.ctx.Request.Context()))
	} else {
		logger.TraceLog.Error(errorLogMsg)
	}

	response.ctx.JSON(status, errorException)
}
