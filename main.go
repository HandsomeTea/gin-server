package main

import (
	"gin-server/server"
	env "gin-server/server/configs/env"
	logger "gin-server/server/configs/logger"
	middleware "gin-server/server/middlewares"
	v1 "gin-server/server/routers/v1"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	goEnv := env.GetEnv("GO_ENV")

	if goEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InitLogger()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.SystemLog.Debug(httpMethod + ": " + absolutePath + " --> " + handlerName + " (" + strconv.Itoa(nuHandlers) + " handlers)")
	}

	router := server.CreateRouter()

	// 中间件中的c.Next()方法，表示执行接口中的逻辑，一个中间件不调用c.Next()，则全部当成前置中间件，执行顺序与注册顺序一致
	// 一个中间件调用c.Next()，则中间件中c.Next()调用前的代码遵循先注册的中间件先执行，c.Next()调用后的代码遵循先注册的中间件后执行
	// 所有中间件c.Next()前的代码，c.Next()，c.Next()后的代码，组成了一个洋葱模型：
	// 请求先进入洋葱的最外层，然后逐层向内(每个中间件c.Next()前的代码按照中间件的注册顺序依次执行)，
	// 最后到达最内层(由c.Next()在中间件中指代，调用c.Next()即执行接口逻辑)，
	// 处理完接口逻辑后再从洋葱模型中逐层由内向外执行(每个中间件c.Next()后的代码按照中间件的注册顺序反向依次执行)
	// c.Abort()方法会终止洋葱模型向内层继续执行，直接从当前中间件c.Next()处开始向外执行

	// 如果想在前置中间件前终止业务逻辑直接向客户端返回，则需要：
	// c.Abort() // 保证终止洋葱模型向内层继续执行
	// c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success"}) // 示例直接向客户端返回
	// return // 终止当前中间件后续代码执行，包括后续的c.Next()，c.Next()不执行则洋葱模型当前层的后续后置中间件不会执行

	// 在接口逻辑中c.Abort()只会组织洋葱模型向内层继续执行，不会组织洋葱模型向外层继续执行，即不会组织洋葱模型中后续后置中间件的执行
	// 一般在接口逻辑中c.Abort()，同时还需要再c.Error(xxx)，return，并在后置中间件中处理c.Errors，这样就可以在接口逻辑中向客户端返回错误信息，并在后置中间件中记录错误日志

	router.Use(middleware.AcceptRequestHandle())

	v1.RegisterV1Routes(router)

	logger.SystemLog.Debug("Server is running on port 8084")
	http.ListenAndServe(":"+env.GetEnv("PORT"), router)
}
