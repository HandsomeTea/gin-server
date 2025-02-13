package v1UserApi

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserPubRoutes(r *gin.Engine) {
	v1PubGroup := r.Group("/api/projectpub/v1")

	v1PubGroup.GET("/user/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, gin-server v1 public user!",
		})
	})
}
