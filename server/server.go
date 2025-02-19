package server

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	return router
}
