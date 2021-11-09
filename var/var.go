/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/9 10:09
 */

package _var

import "crypto/tls"

type (
	Ports     interface{}
	Hosts     interface{}
	RouteInfo struct {
		Name    string
		TLS     string
		Path    map[string]string
		Service []string
	}
	Service struct {
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
	Services = map[string]Service{}
	TLSFile  = map[string]tls.Certificate{}
)
