package routers

import (
	"github.com/asuncm/vm/auth/routes/code"
	"github.com/asuncm/vm/auth/routes/userInfo"
	"github.com/gin-gonic/gin"
)

func authRouter(route *gin.Engine) {
	authRoute := route.Group("/auth")
	{
		userInfo.AuthCode(authRoute)
	}
	codeRoute := route.Group("/code")
	{
		code.AuthCode(codeRoute)
	}
}
