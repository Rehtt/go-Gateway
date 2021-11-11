package router

import (
	"github.com/gin-gonic/gin"
	service2 "go-Gateway/service"
	_var "go-Gateway/var"
	"net/http"
)

func LoadRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 加载中间件
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.Any("/*router", func(ctx *gin.Context) {
		ctx.Writer.WriteString(ctx.Request.RequestURI)
		serviceName, _ := ctx.Get("Service")
		service := _var.Services[serviceName.(_var.ServiceName)]
		switch service.Type {
		case "proxy":
			service2.ReverseProxy(ctx, service)
		case "filebrowser":
			service2.FileBrowser(ctx, service)
		}
	})
	return g
}
