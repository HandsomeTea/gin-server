package canary

import (
	"gin-server/server/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/index.html", services.CanaryService.AbTestSplit)
}
