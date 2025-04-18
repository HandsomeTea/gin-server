package deploy

import (
	"gin-server/server/configs/env"
	"gin-server/server/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	if env.GetEnv("AB_TEST_MODE") == "redirect" {
		r.GET("/", services.DeployService.AllocatHandler)
	} else if env.GetEnv("AB_TEST_MODE") == "proxy" {
		r.GET("/index.html", services.DeployService.AllocatHandler)
	}
}
