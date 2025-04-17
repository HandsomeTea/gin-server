package apigateway

import (
	// "gin-server/server/services"

	// "gin-server/server/middlewares"

	// ginregex "github.com/jxskiss/ginregex"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// apiGroup := r.Group("/api")
	// methods := []string{"GET", "POST", "PUT", "DELETE"}
	// matchers := make([]*ginregex.Matcher, 0, len(methods)*3)

	// for _, method := range methods {
	// 	matchers = append(matchers, ginregex.NewMatcher(method, `/cppbuildcli/v(\d+)/.*`, services.ApiGatewayService.CliAuth))
	// 	matchers = append(matchers, ginregex.NewMatcher(method, `/cppbuild/v(\d+)/.*`, services.ApiGatewayService.LoginAuth))
	// 	matchers = append(matchers, ginregex.NewMatcher(method, `/cppbuildpub/v(\d+)/.*`, services.ApiGatewayService.PubAuth))
	// }

	// apiGroup.Any("/*path", ginregex.Dispatch(matchers...), middlewares.NoRouteHandle)
}
