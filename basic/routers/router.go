package routers

import (
	"github.com/asuncm/vm/basic/routers/config"
	"github.com/gin-gonic/gin"
)

// 用户中心路由
func userRouter(route *gin.Engine) {
	user := route.Group("/user")
	{
		go config.UserInit(user)
	}
}

// web组件路由
func webRouter(route *gin.Engine) {

}

// 表单路由
func formRouter(route *gin.Engine) {}

// 供应链路由
func scmRouter(route *gin.Engine) {}

// 教务路由
func teachRouter(route *gin.Engine) {}

// web生产路由
func websiteRouter(route *gin.Engine) {

}

// 流程路由
func flowRouter(route *gin.Engine) {}

// 教育资源路由
func oerRouter(route *gin.Engine) {

}

// 认证服务路由
func authRouter(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		config.AuthInit(auth)
	}
	code := route.Group("/code")
	{
		config.CodeInit(code)
	}
}

// 设计服务路由
func designRouter(route *gin.Engine) {}

// 事务管理路由
func taskRouter(route *gin.Engine) {}

// 信息资源路由
func irmRouter(route *gin.Engine) {

}

// 文档管理路由
func docRouter(route *gin.Engine) {

}

// 音视频资源路由
func avpRouter(route *gin.Engine) {}

// 营销服务路由
func semRouter(route *gin.Engine) {}
