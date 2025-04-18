package gateway

import (
	"gin-server/server/configs/env"
	"gin-server/server/routers/gateway/api"
	"gin-server/server/routers/gateway/deploy"

	"github.com/gin-gonic/gin"
)

func RegisterGateway(r *gin.Engine) {
	if env.GetEnv("DEPLOY_MODE") == "api-gateway" {
		api.RegisterRoutes(r)
	} else if env.GetEnv("DEPLOY_MODE") == "ab-test" {
		deploy.RegisterRoutes(r)
	} else {
		deploy.RegisterRoutes(r)
		api.RegisterRoutes(r)
	}
}
