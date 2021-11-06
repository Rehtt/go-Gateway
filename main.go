package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	initConfig "go-Gateway/config"
	"go-Gateway/router"
	"go-Gateway/router/middleware"
	"os"
)

var (
	configFile = flag.String("c", "./config/config.yaml", "配置文件地址")
)

type (
	port string
	host string
)

func init() {
	flag.Parse()
	if err := initConfig.InitConfig(*configFile); err != nil {
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
		//middleware.Options,
		//middleware.NoCache,
	)
	//var a, b, c tls.Certificate
	//conf := &tls.Config{Certificates: []tls.Certificate{a, b, c}}
	//t, _ := tls.Listen("tcp", "0.0.0.0:443", conf)
	//g.RunListener(t)
	go g.Run(":8080")
	//todo 读取配置文件
	configs := viper.Get("app")

	tlss := map[port]map[host]tls.Certificate{}
	for _, c := range configs.([]interface{}) {
		config := c.(map[string]interface{})
		var addr []string
		var port []string
		listenAddress(config, &addr, &port, &tlss)

	}

	//http.ListenAndServe(viper.GetString("server.addr")+":"+viper.GetString("server.port"), g)
	select {}
}
func listenAddress(config map[string]interface{}, a, p *[]string, tlss interface{}) error {
	listen := config["listen"].(map[string]interface{})
	if listen != nil {
		if listen["addr"] != nil {
			*a = listen["addr"].([]string)
		}
		if listen["port"] != nil {
			*p = listen["port"].([]string)
		}
	}
	t := config["tls"].(map[string]interface{})
	if t != nil && t["cert"] != nil && t["key"] != nil {
		f, err := tls.LoadX509KeyPair(t["cert"].(string), t["key"].(string))
		if err != nil {
			return err
		}
		for _, po := range *p {
			// todo
			//tlss.(map[port]map[host]tls.Certificate)[po]=f
		}
	}
}
