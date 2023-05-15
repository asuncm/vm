package authentication

import (
	"github.com/asuncm/vm/service/config"
	"github.com/asuncm/vm/service/middleware"
	"github.com/gin-gonic/gin"
)

func GetUUID(c *gin.Context) {
	id := config.AuthCode()
	middleware.Json(c, id)
}
