package apigateway

import (
	"gin-server/server/routers/api-gateway/gateway"

	"github.com/gin-gonic/gin"
	ginregex "github.com/jxskiss/ginregex"
)

func RegisterApiGateway(r *gin.Engine) {
	var handlers []*ginregex.Matcher

	server1Matchers := gateway.GetServer1AuthMatcher()
	otherMatchers := gateway.GetOtherAuthMatcher()
	handlers = append(handlers, server1Matchers...)
	handlers = append(handlers, otherMatchers...)

	r.Any("/api/*path", ginregex.Dispatch(handlers...))
}
