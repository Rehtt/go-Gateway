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
	ports interface{}
	hosts interface{}
)

func init() {
	flag.Parse()
	if err := initConfig.InitConfig(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var listen = map[ports]map[hosts]*tls.Certificate{}

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
	//todo 读取配置文件
	apps := viper.Get("app")
	for _, app := range apps.([]interface{}) {
		config := app.(map[interface{}]interface{})
		listenAddress(config)
	}
	for port, f := range listen {
		cert := []tls.Certificate{}
		for _, file := range f {
			// todo null
			if file != nil {
				cert = append(cert, *file)
			}
		}
		if len(cert) != 0 {
			conf := &tls.Config{Certificates: cert}
			t, err := tls.Listen("tcp", "0.0.0.0:"+port.(string), conf)
			if err != nil {
				panic(err)
			}
			go g.RunListener(t)
		} else {
			go g.Run(":" + port.(string))
		}

	}

	//http.ListenAndServe(viper.GetString("server.addr")+":"+viper.GetString("server.port"), g)

	select {}
}

// 获取配置文件的端口、hosts以及tls证书
func listenAddress(config map[interface{}]interface{}) error {
	portHost := func(f *tls.Certificate) error {
		por := config["ports"]
		var port []interface{}
		if por == nil {
			if f == nil {
				port = []interface{}{"80"}
			} else {
				port = []interface{}{"443"}
			}
		} else {
			port = por.([]interface{})
		}
		for _, p := range port {
			if listen[ports(p)] == nil {
				listen[ports(p)] = make(map[hosts]*tls.Certificate)
			}
			if config["hosts"] == nil || len(config["hosts"].([]interface{})) == 0 {
				return fmt.Errorf("null host")
			}
			for _, h := range config["hosts"].([]interface{}) {
				listen[ports(p)][hosts(h)] = f
			}
		}
		return nil
	}
	t := config["tls"]
	if t != nil {
		tl := t.(map[interface{}]interface{})
		if tl["cert"] == nil || tl["key"] == nil {
			return fmt.Errorf("not tls file")
		}
		f, err := tls.LoadX509KeyPair(tl["cert"].(string), tl["key"].(string))
		if err != nil {
			return err
		}
		return portHost(&f)
	}
	return portHost(nil)

}
