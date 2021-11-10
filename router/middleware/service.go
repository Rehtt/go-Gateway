/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/10 11:48
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	_var "go-Gateway/var"
	"log"
	"net/http"
	"regexp"
)

func Filter(ctx *gin.Context) {
	host := ctx.Request.Host
	path := ctx.Request.RequestURI
	for p, pp := range _var.Listen[_var.Host(host)].Path {
		matched, err := regexp.MatchString(p, path)
		if err != nil {
			log.Panicln(err)
		}
		if matched {
			ctx.Set("BlackList", pp.BlackList)
			ctx.Set("Service", pp.ServiceName)
			ctx.Next()
			return
		}
	}
	ctx.String(http.StatusNotFound, "The incorrect Path route.")
	ctx.Abort()
}
