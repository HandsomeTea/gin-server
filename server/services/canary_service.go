package services

import (
	"encoding/json"
	"gin-server/server/configs/logger"
	"gin-server/server/configs/response"

	"github.com/gin-gonic/gin"
)

type canaryService struct{}

var CanaryService = canaryService{}

func (user canaryService) AbTestSplit(c *gin.Context) {
	path := c.Param("path")

	headerJson, _ := json.MarshalIndent(c.Request.Header, "", "    ")
	logger.Log.Info(string(headerJson))

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "ab test"})
}
