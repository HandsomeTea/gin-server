package middlewares

import (
	"fmt"
	"gin-server/server/configs/logger"
	"gin-server/server/configs/response"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

/*
* 不建议使用后置中间件，因为接口逻辑和前置中间件中的c.Abort()并不会阻止洋葱模型从当前层向外执行
* 如果要阻止洋葱模型继续向外执行，则需要zaiv.Abort()处c.Error(xxx)，并在后置中间件中做错误判断，才能终止，这样增加了逻辑控制的复杂性
 **/
func ErrorHandle(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		// errorData := c.Errors[0]
		// response.Ctx(c).Failed(errorData)
	}
}

func ExceptionHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%v", err)
			response.Ctx(c).Failed(message)
			logger.SystemLog.Error(message + "\n" + string(debug.Stack()))
		}
	}()

	c.Next()
}

func NoRouteHandle(c *gin.Context) {
	response.Ctx(c).Failed("404 not found", "URL_NOT_FOUND", c.Request.Method+": "+c.Request.URL.RequestURI())
	c.Abort()
}
