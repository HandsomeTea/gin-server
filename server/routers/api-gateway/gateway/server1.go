package gateway

import (
	"gin-server/server/services"

	ginregex "github.com/jxskiss/ginregex"
)

func GetServer1AuthMatcher() []*ginregex.Matcher {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	matchers := make([]*ginregex.Matcher, 0, len(methods)*2)

	for _, method := range methods {
		matchers = append(matchers, ginregex.NewMatcher(method, `^/api/server1adm/v(\d+)/.*`, services.ApiGatewayService.Server1AdmAuth))
		matchers = append(matchers, ginregex.NewMatcher(method, `^/api/server1/v(\d+)/.*`, services.ApiGatewayService.Server1LoginAuth))
	}

	return matchers
}
