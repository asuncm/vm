package middleware

import (
	"fmt"
	"github.com/asuncm/vm/service/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

/*@package Proxy 	 接口反向代理
* @param   json	 	 返回接口数据
* @param   code      接口编码
* @param   message   错误信息
* @param   file		 文件路径
* @param   value     map数据key值
 */
func Proxy(c *gin.Context, file string, value string) {
	conf, err := config.Config(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "读取配置文件错误",
			"data":    fmt.Errorf("获取配置文件失败"),
		})
	} else {
		options := conf.Services
		auth := options[value]
		// 拼接string
		target := strings.Join([]string{auth["host"], auth["port"]}, ":")
		target = strings.Join([]string{"http:", target}, "//")
		// 创建反向代理路径
		proxyUrl, _ := url.Parse(target)
		// 创建反向代理服务
		proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
		// 发送代理请求
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
