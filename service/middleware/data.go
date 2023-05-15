package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回json数据格式
func Json(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "请求成功",
		"data":    body,
	})
}

// 返回form-data数据格式
func Form(c *gin.Context, body interface{}) {

}
