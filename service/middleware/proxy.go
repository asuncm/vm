package middleware

import (
	"github.com/asuncm/vm/service/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Proxy(c *gin.Context, file string, value string) {
	conf, err := config.Config(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "读取配置文件错误",
			"data":    err.Error(),
		})
	} else {
		options := conf.Services
		auth := options[value]
		target := strings.Join([]string{auth["host"], auth["port"]}, ":")
		target = strings.Join([]string{"http:", target}, "//")
		proxyUrl, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
