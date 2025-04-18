package gateway

import (
	"gin-server/server/configs/env"
	"gin-server/server/routers/gateway/api"
	"gin-server/server/routers/gateway/canary"

	"github.com/gin-gonic/gin"
)

func RegisterGateway(r *gin.Engine) {
	if env.GetEnv("DEPLOY_MODE") == "api-gateway" {
		api.RegisterRoutes(r)
	} else if env.GetEnv("DEPLOY_MODE") == "canary" {
		canary.RegisterRoutes(r)
	} else {
		canary.RegisterRoutes(r)
		api.RegisterRoutes(r)
	}
}
