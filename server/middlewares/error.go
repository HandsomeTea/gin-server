package middlewares

import (
	"encoding/json"
	"fmt"
	httpError "gin-server/server/configs/error"
	"gin-server/server/configs/logger"
	"gin-server/server/configs/response"
	"gin-server/server/globals"
	"html"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ExceptionHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			errStr := fmt.Sprintf("%v", err)
			e := globals.HttpException{}
			json.Unmarshal([]byte(errStr), &e)

			logger.SystemLog.Error(e.Message + "\n" + string(debug.Stack()))
			response.Ctx(c).Failed(e.Message, e.Code, e.Data)
			c.Abort()
		}
	}()

	c.Next()
}

func NoRouteHandle(c *gin.Context) {
	response.Ctx(c).Failed("url not found", httpError.URL_NOT_FOUND, c.Request.Method+": "+html.EscapeString(c.Request.URL.RequestURI()))
	c.Abort()
}
