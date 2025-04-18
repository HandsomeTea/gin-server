package services

import (
	"gin-server/server/configs/response"

	"github.com/gin-gonic/gin"
)

type apiGatewayService struct{}

var ApiGatewayService = apiGatewayService{}

func (svc apiGatewayService) CliAuth(c *gin.Context) {
	path := c.Param("path")

	// print(path[100])

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "cli"})
}

func (svc apiGatewayService) LoginAuth(c *gin.Context) {
	path := c.Param("path")

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "login"})
}

func (svc apiGatewayService) PubAuth(c *gin.Context) {
	path := c.Param("path")

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "pub"})
}
