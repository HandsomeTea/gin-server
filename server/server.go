package server

import (
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	server := gin.Default()
	server.SetTrustedProxies([]string{"127.0.0.1"})

	return server
}
