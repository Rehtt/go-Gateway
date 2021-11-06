package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestV1(ctx *gin.Context) {
	ctx.String(http.StatusOK, "RemoteAddr :%s", ctx.Request.RemoteAddr)
}
