package services

import (
	"encoding/json"
	"fmt"
	"gin-server/server/configs/env"
	httpError "gin-server/server/configs/error"
	"gin-server/server/configs/logger"
	"gin-server/server/configs/response"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

type canaryService struct{}

var CanaryService = canaryService{}

func (svc canaryService) shouldBeAllocated() bool {
	_rate := env.GetEnv("CANARY_DIVERSION_PERCENTAGE")
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

func (svc canaryService) AllocatHandler(c *gin.Context) {
	path := c.Request.URL.Path

	headerJson, _ := json.MarshalIndent(c.Request.Header, "", "    ")
	logger.Log.Info(string(headerJson))
	allocated := svc.shouldBeAllocated()

	fmt.Print(map[string]any{"path": path, "auth": "allocat", "allocated": allocated})
	if allocated {
		if env.GetEnv("CANARY_MODE") == "redirect" {
			response.Ctx(c).Failed("When you receive this message you should redirect the request somewhere else", httpError.MOVED_PERMANENTLY)
			return
		} else if env.GetEnv("CANARY_MODE") == "proxy" {
			response.Ctx(c).Failed("When you get this message you should proxy the request somewhere else", httpError.USE_PROXY)
			return
		}
	}

	response.Ctx(c).Success()
}
