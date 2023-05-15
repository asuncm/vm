package routers

import "github.com/gin-gonic/gin"

func add(route *gin.Engine) {
	authRouter(route)
}

// 初始化gin
func Init() *gin.Engine {
	router := gin.Default()
	add(router)
	return router
}
