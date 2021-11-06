package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

type ip [4][2]uint8

var blockList []ip

func initBlockList() {
	blockList = make([]ip, 0, 10)
	list := viper.GetStringSlice("block")
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
		p := ip{}
		for i := range arr {
			p[i] = arr[i]
		}
		blockList = append(blockList, p)

	}
}
func Block(context *gin.Context) {
	if blockList == nil {
		initBlockList()
	}
	r_ip := strings.Split(context.Request.RemoteAddr, ":")[0]
	x_ip := context.Request.Header.Get("X-Forwarded-For")
	if x_ip != "" {
		r_ip = x_ip
	}
	var p ip
	fmt.Sscanf(r_ip, "%d.%d.%d.%d", &p[0][0], &p[1][0], &p[2][0], &p[3][0])
	flag := false
	for _, v := range blockList {
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
	fmt.Println(blockList)
	if flag {
		context.String(http.StatusBadRequest, "bad request")
		context.Abort()
		return
	}
	context.Set("ip", r_ip)
	context.Next()
}
