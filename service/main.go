package main

import (
	"github.com/asuncm/vm/service/config"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()
	conf, err := config.Config("/service")
	if err != nil {
		router.Run(":36003")
	}
	options := conf.Services
	serve := options["service"]
	router.Run(strings.Join([]string{serve["host"], serve["port"]}, ":"))
}
