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
	Ports       string
	Hosts       string
	IP          [4][2]uint8
	Name        string
	ServiceName string
	TLSName     string
	RouteInfo   struct {
		Name      string
		TLS       string
		Path      map[string]ServiceName
		BlackList []IP
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
	Listen   = map[Ports]map[Hosts]*RouteInfo{}
	Services = map[ServiceName]Service{}
	TLSFile  = map[TLSName]tls.Certificate{}
	Config   = models.Config{}
)
