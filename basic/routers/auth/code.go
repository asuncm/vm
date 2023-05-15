package auth

import (
	"github.com/asuncm/vm/service/middleware"
	"github.com/gin-gonic/gin"
)

func CodeVerify(c *gin.Context) {
	middleware.Proxy(c, "/basic", "auth")
}
