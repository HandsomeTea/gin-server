package middlewares

import (
	"fmt"
	httpError "gin-server/server/configs/error"
	"gin-server/server/configs/logger"
	"gin-server/server/configs/response"
	"html"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ExceptionHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%v", err)
			logger.SystemLog.Error(message + "\n" + string(debug.Stack()))
			response.Ctx(c).Failed(message)
			c.Abort()
		}
	}()

	c.Next()
}

func NoRouteHandle(c *gin.Context) {
	response.Ctx(c).Failed("url not found", httpError.URL_NOT_FOUND, c.Request.Method+": "+html.EscapeString(c.Request.URL.RequestURI()))
	c.Abort()
}
