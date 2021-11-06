package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-Gateway/config"
	"go-Gateway/router"
	"go-Gateway/router/middleware"
	"os"
)

var (
	configFile = flag.String("c", "./config/config.yaml", "配置文件地址")
)

func init() {
	flag.Parse()
	if err := config.InitConfig(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	//// 初始化mysql数据库
	//if err := mysql.DB.InitDB(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//// 初始化redis
	//if err := redis.Init(); err != nil {
	//	fmt.Println(err)
	//}
	g := gin.New()
	router.LoadRouter(
		g,
		// 中间件
		middleware.Block,
		middleware.TLS,
		//middleware.Options,
		//middleware.NoCache,
	)
	go g.Run(":8080")
	//todo 读取配置文件
	//configs:=viper.Get("app")
	//for _,config:=range configs.([]interface{}){
	//}

	//http.ListenAndServe(viper.GetString("server.addr")+":"+viper.GetString("server.port"), g)
	select {}
}
