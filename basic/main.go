package main

import (
	"github.com/asuncm/vm/basic/middleware"
	"github.com/asuncm/vm/basic/routers"
	"github.com/asuncm/vm/service/config"
	"strings"
)

func main() {
	// 初始化gin服务实例
	router := routers.Init()
	// 获取配置文件
	conf, err := config.Config()
	// 设置中间件
	go router.Use(middleware.Middleware(conf))
	if err != nil {
		router.Run(":36001")
	} else {
		// 启动服务
		options := conf.Services
		serve := options["service"]
		router.Run(strings.Join([]string{serve["host"], serve["port"]}, ":"))
	}
}
