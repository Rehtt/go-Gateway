package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	initConfig "go-Gateway/config"
	"go-Gateway/router"
	"go-Gateway/router/middleware"
	_var "go-Gateway/var"
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
	viper.Unmarshal(&_var.Config)
}

func main() {
	g := gin.New()
	router.LoadRouter(
		g,
		// 中间件
		middleware.Filter,
		middleware.Block,
	)
	initApp()
	openPort(g)
}
