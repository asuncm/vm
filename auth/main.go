package main

import (
	"github.com/asuncm/vm/auth/routes"
	"github.com/asuncm/vm/service/config"
	"strings"
)

func main() {
	// 初始化gin服务实例
	router := routers.Init()
	// 获取配置文件
	conf, err := config.Config("/auth")
	// 设置中间件
	if err != nil {
		router.Run(":36003")
	} else {
		// 启动服务
		options := conf.Services
		serve := options["auth"]
		router.Run(strings.Join([]string{serve["host"], serve["port"]}, ":"))
	}
}
