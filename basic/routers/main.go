package routers

import (
	"github.com/gin-gonic/gin"
)

func add(route *gin.Engine) {
	webRouter(route)
	userRouter(route)
	formRouter(route)
	scmRouter(route)
	semRouter(route)
	flowRouter(route)
	webRouter(route)
	websiteRouter(route)
	teachRouter(route)
	authRouter(route)
	docRouter(route)
	designRouter(route)
	oerRouter(route)
	taskRouter(route)
	irmRouter(route)
	avpRouter(route)
}

// 初始化gin
func Init() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	add(router)
	return router
}
