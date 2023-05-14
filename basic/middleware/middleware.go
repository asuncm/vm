package middleware

import (
	user "github.com/asuncm/vm/service/badger/userInfo"
	"github.com/asuncm/vm/service/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(options config.ComConf) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		contentType := c.Request.Header.Get("Content-Type")
		token := c.Request.Header.Get("Token")
		verify := c.Request.Header.Get("verify")
		config := user.Authorization{
			Token:  token,
			Verify: verify,
			Origin: origin,
		}
		userInfo, _ := Authorization(options, config)
		if method != "OPTIONS" && userInfo.Status {
			// 设置接口跨域信息，订制域名限制
			if userInfo.Status {
				c.Header("Access-Control-Allow-Origin", userInfo.Origin)
			}
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Content-Type", "application/json; charset=utf-8")
			if contentType != "" {
				c.Header("Content-Type", contentType)
			}
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
