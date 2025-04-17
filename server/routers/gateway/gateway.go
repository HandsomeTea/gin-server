package gateway

import (
	"gin-server/server/routers/gateway/api"
	"gin-server/server/routers/gateway/canary"

	"github.com/gin-gonic/gin"
)

func RegisterGateway(r *gin.Engine) {
	api.RegisterRoutes(r)
	canary.RegisterRoutes(r)
}
