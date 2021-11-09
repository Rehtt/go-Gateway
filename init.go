/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/9 12:24
 */

package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_var "go-Gateway/var"
	"log"
)

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

// 获取service
func GetServices() {
	services := viper.Get("services")
	for _, s := range services.([]interface{}) {
		service := s.(map[interface{}]interface{})
		if service["name"] != nil && service["type"] != nil {
			//todo 优化
			switch service["type"].(string) {
			case "proxy":
				h := service["header"].([]interface{})
				header := map[string]string{}
				for _, v := range h {
					key_value := v.(map[interface{}]interface{})
					header[key_value["key"].(string)] = key_value["value"].(string)
				}
				_var.Services[service["name"].(string)] = _var.Service{
					Type: "proxy",
					Proxy: &_var.Proxy{
						Addr:   service["addr"].(string),
						Header: header,
					},
				}
			case "filebrowser":
				_var.Services[service["name"].(string)] = _var.Service{
					Type:        "filebrowser",
					FileBrowser: &_var.FileBrowser{Root: service["root"].(string)},
				}
			}
		}
	}
}

// 监听端口
func OpenPort(engine *gin.Engine) {
	for port, hosts := range _var.Listen {
		cert := map[string]*tls.Certificate{}
		for _, route := range hosts {
			if route.TLS != "" {
				if file, ok := _var.TLSFile[route.TLS]; ok {
					cert[route.TLS] = &file
				}
			}
		}
		if len(cert) == 0 {
			go engine.Run(":" + port.(string))
		} else {
			tlss := []tls.Certificate{}
			for _, file := range cert {
				tlss = append(tlss, *file)
			}
			listen, err := tls.Listen("tcp", "0.0.0.0:"+port.(string), &tls.Config{Certificates: tlss})
			if err != nil {
				log.Println(err)
				return
			}
			go engine.RunListener(listen)
		}
	}
	select {}
}
