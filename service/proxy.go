/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2021/11/10 14:09
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_var "go-Gateway/var"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(ctx *gin.Context, service *_var.Service) {
	u, _ := url.Parse(service.Proxy.Addr)
	proxy := httputil.NewSingleHostReverseProxy(u)

	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("http: proxy error: %v", err)
		ret := fmt.Sprintf("http proxy error %v", err)

		rw.Write([]byte(ret))
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
