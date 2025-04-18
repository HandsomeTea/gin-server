package services

import (
	"encoding/json"
	"fmt"
	"gin-server/server/configs/env"
	"gin-server/server/configs/logger"
	"gin-server/server/configs/response"
	"math/rand"
	"strconv"

	httpError "gin-server/server/configs/error"

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

// =============================================== ab-test ================================================

type deployService struct{}

var DeployService = deployService{}

func (svc deployService) shouldBeAllocated() bool {
	_rate := env.GetEnv("AB_TEST_DIVERSION_PERCENTAGE")
	rate := 0

	if v, err := strconv.Atoi(_rate); err == nil {
		rate = v
	}

	if rate == 0 {
		return false
	}
	// 生成(0,100]的随机数
	randomNum := rand.Intn(100) + 1

	return randomNum <= rate
}

func (svc deployService) AllocatHandler(c *gin.Context) {
	path := c.Request.URL.Path

	headerJson, _ := json.MarshalIndent(c.Request.Header, "", "    ")
	logger.Log.Info(string(headerJson))
	allocated := svc.shouldBeAllocated()

	fmt.Print(map[string]any{"path": path, "auth": "allocat", "allocated": allocated})
	if allocated {
		if env.GetEnv("AB_TEST_MODE") == "redirect" {
			response.Ctx(c).Failed("When you receive this message you should redirect the request somewhere else", httpError.MOVED_PERMANENTLY)
			return
		} else if env.GetEnv("AB_TEST_MODE") == "proxy" {
			response.Ctx(c).Failed("When you get this message you should proxy the request somewhere else", httpError.USE_PROXY)
			return
		}
	}

	response.Ctx(c).Success()
}
