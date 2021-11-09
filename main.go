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
	_var "go-Gateway/var"
	"log"
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

	select {}
}

// 获取tls文件
func GetTLS() {
	config := viper.Get("tls")
	for _, c := range config.([]interface{}) {
		t := c.(map[interface{}]interface{})
		if t["cert"] == nil || t["key"] == nil || t["name"] == nil {
			continue
		}
		file, err := tls.LoadX509KeyPair(t["cert"].(string), t["key"].(string))
		if err != nil {
			log.Println(err)
			continue
		}
		_var.TLSFile[t["name"].(string)] = file
	}
}

func GetApps() {
	//todo 读取配置文件
	apps := viper.Get("app")
	for _, app := range apps.([]interface{}) {
		config := app.(map[interface{}]interface{})

	}
}

// todo 读取service
func GetServices() {

}
