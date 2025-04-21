package services

import (
	"gin-server/server/configs/response"

	"github.com/gin-gonic/gin"
)

type apiGatewayService struct{}

var ApiGatewayService = apiGatewayService{}

func (svc apiGatewayService) Server1AdmAuth(c *gin.Context) {
	path := c.Param("path")

	// print(path[100])

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "admin校验"})
}

func (svc apiGatewayService) Server1LoginAuth(c *gin.Context) {
	path := c.Param("path")

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "登录校验"})
}

func (svc apiGatewayService) OtherAuth(c *gin.Context) {
	path := c.Param("path")

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "其它"})
}
