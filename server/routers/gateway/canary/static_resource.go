package canary

import (
	"gin-server/server/configs/env"
	"gin-server/server/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	if env.GetEnv("CANARY_MODE") == "redirect" {
		r.GET("/", services.CanaryService.AllocatHandler)
	} else if env.GetEnv("CANARY_MODE") == "proxy" {
		r.GET("/index.html", services.CanaryService.AllocatHandler)
	}
}
