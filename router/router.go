package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 加载中间件
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.Any("/", func(context *gin.Context) {
		context.Writer.WriteString("123")
	})
	return g
}
