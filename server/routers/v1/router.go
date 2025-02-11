package v1

import (
	v1User "gin-server/server/routers/v1/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesAll(r *gin.Engine) {
	v1User.RegisterUserRoutes(r)
	v1User.RegisterUserPubRoutes(r)
	v1User.RegisterUserAdmRoutes(r)
}
