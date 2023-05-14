package config

import (
	"github.com/asuncm/vm/basic/routers/auth"
	"github.com/gin-gonic/gin"
)

func AuthInit(route *gin.RouterGroup) {
	// 获取身份令牌
	route.GET("/verify", auth.Verify)
	// 获取验证码
	route.GET("/code", auth.Code)
	// 用户登录
	route.POST("/login", auth.Login)
	// 用户注册
	route.POST("/sign", auth.Sign)
	// 换取平台登录token
	route.PUT("/token", auth.Token)
	// 用户身份验证
	route.POST("/authentication", auth.Authentication)
	//	找回密码和用户名
	route.POST("/forget", auth.Forget)
}

func CodeInit(route *gin.RouterGroup) {

}
