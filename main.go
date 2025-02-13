package main

import (
	"encoding/json"
	"gin-server/server"
	env "gin-server/server/configs/env"
	logger "gin-server/server/configs/logger"
	v1 "gin-server/server/routers/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	goEnv := env.GetEnv("GO_ENV")
	if goEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.InitLogger()

	r := server.CreateServer()

	v1.RegisterV1Routes(r)
	logger.Log.Info("Server is running on port 8084")

	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	user := User{ID: 1001, Username: "johndoe"}
	jsonBytes, _ := json.MarshalIndent(user, "", "    ")
	logger.TraceLog.Info("[http-request] GET: /api/project/v1/user/test => " + string(jsonBytes))

	logger.SystemLog.Debug("Server is running on port 8084")
	r.Run(":" + env.GetEnv("PORT"))
}
