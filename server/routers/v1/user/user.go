package v1UserApi

import (
	service "gin-server/server/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	v1Group := r.Group("/api/project/v1")

	v1Group.GET("/user/test", func(c *gin.Context) {
		query := c.DefaultQuery("ss", "123")

		print(query)
		c.JSON(200, service.UserService.Test())
	})

	v1Group.POST("/user/test", func(c *gin.Context) {
		c.JSON(200, service.UserService.Test())
	})
}
