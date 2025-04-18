package v1

import (
	v1UserApi "gin-server/server/routers/v1/user"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(r *gin.Engine) {
	v1UserApi.RegisterUserRoutes(r)
	v1UserApi.RegisterUserPubRoutes(r)
	v1UserApi.RegisterUserAdmRoutes(r)
}
