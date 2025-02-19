package v1UserApi

import (
	services "gin-server/server/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	v1Group := r.Group("/api/project/v1")

	v1Group.GET("/user/:test", services.UserService.TestApi)
}
