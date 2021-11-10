/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/9 10:09
 */

package _var

import (
	"crypto/tls"
	"go-Gateway/models"
)

type (
	Port        string
	Host        string
	IP          [4][2]uint8
	Name        string
	ServiceName string
	TLSName     string
	RouteInfo   struct {
		Name string
		Port Port
		TLS  TLSName
		Path map[string]Path
	}
	Path struct {
		BlackList []IP
		ServiceName
	}
	Service struct {
		Name string
		Type string
		*Proxy
		*FileBrowser
	}
	Proxy struct {
		Addr   string
		Header map[string]string
	}
	FileBrowser struct {
		Root string
	}
)

var (
	Listen = map[Host]*RouteInfo{}

	Ports    = map[Port]map[TLSName]struct{}{}
	Services = map[ServiceName]*Service{}
	TLSFile  = map[TLSName]tls.Certificate{}
	Config   = models.Config{}
)
