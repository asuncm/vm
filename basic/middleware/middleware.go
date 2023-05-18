package middleware

import (
	"github.com/asuncm/vm/service/badger"
	user "github.com/asuncm/vm/service/badger/userInfo"
	"github.com/asuncm/vm/service/config"
	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack"
	"net/http"
	"strconv"
)

type hashMap map[string]interface{}

func Middleware(options config.ComConf) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		contentType := c.Request.Header.Get("Content-Type")
		token := c.Request.Header.Get("token")
		uid := c.Request.Header.Get("uid")
		config := user.Authorization{
			Token:  token,
			Uid:    uid,
			Origin: origin,
		}
		userInfo, _ := Authorization(config)
		configMap, _ := msgpack.Marshal(options)
		c.Request.Header.Set("configs", string(configMap))
		status := strconv.FormatBool(userInfo.Status)
		c.Request.Header.Set("status", status)
		if method != "OPTIONS" {
			// 设置接口跨域信息，订制域名限制
			if userInfo.Status {
				users := userInfo.Users
				c.Request.Header.Set("timestamp", users["Timestamp"].(string))
				c.Request.Header.Set("hash", users["Hash"].(string))
				c.Header("Access-Control-Allow-Origin", userInfo.Origin)
				if users["nowTime"] != nil {
					tNum := TimeLag(users)
					if tNum > 20 {
						var baConf hashMap
						baData, _ := msgpack.Marshal(users)
						msgpack.Unmarshal(baData, &baConf)
						badger.Update([]byte(uid), baConf, "/basic", users["TTL"])
					}
				}
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
