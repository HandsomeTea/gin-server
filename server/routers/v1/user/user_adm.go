package v1UserApi

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserAdmRoutes(r *gin.Engine) {
	v1AdmGroup := r.Group("/api/projectadm/v1")

	v1AdmGroup.GET("/user/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, gin-server v1 admin user!",
		})
	})
}
