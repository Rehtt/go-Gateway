package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	initConfig "go-Gateway/config"
	"go-Gateway/router"
	"go-Gateway/router/middleware"
	"os"
)

var (
	configFile = flag.String("c", "./config/config.yaml", "配置文件地址")
)

func init() {
	flag.Parse()
	if err := initConfig.InitConfig(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	g := gin.New()
	router.LoadRouter(
		g,
		// 中间件
		middleware.Block,
	)
	GetTLS()
	GetServices()
	GetApps()
	OpenPort(g)
}
