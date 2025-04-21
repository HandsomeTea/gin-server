package gateway

import (
	"gin-server/server/services"

	ginregex "github.com/jxskiss/ginregex"
)

func GetOtherAuthMatcher() []*ginregex.Matcher {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	matchers := make([]*ginregex.Matcher, 0, len(methods)*1)

	for _, method := range methods {
		matchers = append(matchers, ginregex.NewMatcher(method, `^/api/v(\d+)/.*`, services.ApiGatewayService.OtherAuth))
	}

	return matchers
}
