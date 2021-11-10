package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_var "go-Gateway/var"
	"net/http"
	"strings"
)

func Block(context *gin.Context) {
	// todo 黑名单
	r_ip := strings.Split(context.Request.RemoteAddr, ":")[0]
	x_ip := context.Request.Header.Get("X-Forwarded-For")
	host := context.Request.Host
	if host == "" {
		context.Next()
		return
	}
	path := context.Request.RequestURI
	if x_ip != "" {
		r_ip = x_ip
	}
	var p _var.IP
	fmt.Sscanf(r_ip, "%d.%d.%d.%d", &p[0][0], &p[1][0], &p[2][0], &p[3][0])
	flag := false
	for _, v := range _var.Listen[_var.Host(host)].Path[path].BlackList {
		f := 0
		for i, vv := range v {
			if p[i][0] >= vv[0] && p[i][0] <= vv[1] {
				f++
			}
		}
		if f == 4 {
			flag = true
			break
		}
	}

	if flag {
		context.String(http.StatusBadRequest, "bad request")
		context.Abort()
		return
	}
	context.Set("ip", r_ip)
	context.Next()
}
