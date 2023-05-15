package userInfo

import (
	"github.com/asuncm/vm/auth/authentication"
	"github.com/gin-gonic/gin"
)

func AuthCode(c *gin.RouterGroup) {
	c.GET("/code", authentication.GetUUID)
}
