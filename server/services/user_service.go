package service

import (
	response "gin-server/server/configs/response"

	"github.com/gin-gonic/gin"
)

type userService struct {
}

var UserService = userService{}

func (user userService) TestApi(c *gin.Context) {
	print("进入接口逻辑")
	// query := c.DefaultQuery("ss", "123")
	// query := c.Query("ss")
	// param := c.Param("test")
	// body := make(map[string]interface{})
	// c.BindJSON(&body)
	// fmt.Println(query, param, body, body["user"])

	// 以下获取参数的方法需要在receive中间件中设置
	// _query, _ := c.Get("query")
	// query, _ := _query.(map[string]interface{})
	// _body, _ := c.Get("body")
	// body := _body.(map[string]interface{})
	// _param, _ := c.Get("param")
	// param := _param.(map[string]interface{})
	// fmt.Println(query["ss"], query, body, param)

	// if true {
	// 	response.Ctx(c).Failed("演示错误返回信息", httpError.NOT_FOUND)
	// 	return
	// }

	print("已响应客户端还继续执行了")
	response.Ctx(c).Success()
}
