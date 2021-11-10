package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_var "go-Gateway/var"
	"net/http"
	"strings"
)

func Block(context *gin.Context) {
	r_ip := strings.Split(context.Request.RemoteAddr, ":")[0]
	x_ip := context.Request.Header.Get("X-Forwarded-For")
	if x_ip != "" {
		r_ip = x_ip
	}
	// 真实ip
	context.Set("ip", r_ip)
	var BlackList []_var.IP
	if blackList, ok := context.Get("BlackList"); ok {
		BlackList = blackList.([]_var.IP)
	}

	var p _var.IP
	fmt.Sscanf(r_ip, "%d.%d.%d.%d", &p[0][0], &p[1][0], &p[2][0], &p[3][0])
	flag := false
	for _, v := range BlackList {
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

	context.Next()
}
