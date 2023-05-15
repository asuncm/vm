package routers

import (
	"github.com/asuncm/vm/basic/middleware"
	"github.com/asuncm/vm/service/config"
	"github.com/gin-gonic/gin"
)

func add(route *gin.Engine) {
	go webRouter(route)
	go userRouter(route)
	go formRouter(route)
	go scmRouter(route)
	go semRouter(route)
	go flowRouter(route)
	go webRouter(route)
	go websiteRouter(route)
	go teachRouter(route)
	go authRouter(route)
	go docRouter(route)
	go designRouter(route)
	go oerRouter(route)
	go taskRouter(route)
	go irmRouter(route)
	go avpRouter(route)
}

// 初始化gin
func Init(options config.ComConf) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(middleware.Middleware(options))
	add(router)
	return router
}
