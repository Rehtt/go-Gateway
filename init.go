/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/9 12:24
 */

//todo 待优化

package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	_var "go-Gateway/var"
	"log"
	"runtime"
)

func initApp() {
	getTLS()
	getServices()
	getApps()
	runtime.GC()
}

// 获取tls文件
func getTLS() {
	for _, c := range _var.Config.TLS {
		if c.Cert == "" || c.Key == "" || c.Name == "" {
			continue
		}
		file, err := tls.LoadX509KeyPair(c.Cert, c.Key)
		if err != nil {
			log.Println(err)
			continue
		}
		_var.TLSFile[_var.TLSName(c.Name)] = file
	}
}

func getApps() {
	for _, app := range _var.Config.App {
		if app.Name == "" {
			log.Fatalln("app name is null")
		}
		path := map[string]_var.Path{}
		for _, p := range app.Path {
			path[p.Path] = _var.Path{
				BlackList:   initBlockList(p.Block),
				ServiceName: "",
			}
		}
		route := _var.RouteInfo{
			Name: app.Name,
			Port: _var.Port(app.Port),
			TLS:  _var.TLSName(app.TLS),
			Path: path,
		}
		// 收集监听端口
		if _var.Ports[route.Port] == nil {
			_var.Ports[route.Port] = make(map[_var.TLSName]struct{})
		}
		_var.Ports[route.Port][route.TLS] = struct{}{}

		for _, h := range app.Hosts {
			_var.Listen[_var.Host(h)] = &route
		}
	}
}

// 获取service
func getServices() {
	for _, service := range _var.Config.Services {
		if service.Name != "" && service.Type != "" {
			//todo 优化
			switch service.Type {
			case "proxy":
				h := service.Header
				header := map[string]string{}
				for _, key_value := range h {
					header[key_value.Key] = key_value.Value
				}
				_var.Services[_var.ServiceName(service.Name)] = &_var.Service{
					Name: service.Name,
					Type: "proxy",
					Proxy: &_var.Proxy{
						Addr:   service.Addr,
						Header: header,
					},
				}
			case "filebrowser":
				_var.Services[_var.ServiceName(service.Name)] = &_var.Service{
					Name:        service.Name,
					Type:        "filebrowser",
					FileBrowser: &_var.FileBrowser{Root: service.Root},
				}
			default:
				log.Println("config error service type:", service.Type)
			}
		}
	}
}

// 监听端口
func openPort(engine *gin.Engine) {
	for port, tlsName := range _var.Ports {
		certs := []tls.Certificate{}
		for name := range tlsName {
			if t, ok := _var.TLSFile[name]; ok {
				certs = append(certs, t)
			}

		}
		if len(certs) == 0 {
			go engine.Run(":" + string(port))
		} else {
			listen, err := tls.Listen("tcp", "0.0.0.0:"+string(port), &tls.Config{Certificates: certs})
			if err != nil {
				log.Println(err)
				return
			}
			go engine.RunListener(listen)
		}
	}
	select {}
}

func initBlockList(list []string) []_var.IP {
	blockList := make([]_var.IP, 0, 10)
	for _, l := range list {
		// 根据网段计算ip范围
		var ips [5]uint8
		n, err := fmt.Sscanf(l, "%d.%d.%d.%d/%d", &ips[0], &ips[1], &ips[2], &ips[3], &ips[4])
		if n < 4 {
			panic(err)
		}
		if n == 4 {
			ips[4] = 32
		}
		var arr [4][2]uint8
		for i := range arr {
			if 8 <= ips[4] {
				arr[i][0] = ips[i]
				arr[i][1] = ips[i]
			} else {
				if ips[4] < 0 {
					ips[4] = 0
				}
				arr[i][0] = 1
				arr[i][1] = 1<<(8-ips[4]) - 2
			}
			ips[4] -= 8
		}
		p := _var.IP{}
		for i := range arr {
			p[i] = arr[i]
		}
		blockList = append(blockList, p)
	}
	return blockList
}
