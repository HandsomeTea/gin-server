package middlewares

import (
	"github.com/gin-gonic/gin"
)

/*
* 不建议使用后置中间件，因为接口逻辑和前置中间件中的c.Abort()并不会阻止洋葱模型从当前层向外执行
* 如果要阻止洋葱模型继续向外执行，则需要zaiv.Abort()处c.Error(xxx)，并在后置中间件中做错误判断，才能终止，这样增加了逻辑控制的复杂性
 **/
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// errorData := c.Errors[0]
			// response.Ctx(c).Failed(errorData)
		}
	}
}
