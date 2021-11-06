package middleware

import "github.com/gin-gonic/gin"

var TLSMap map[string]tls

type tls struct {
	key  []byte
	cert []byte
}

func init() {
	TLSMap = make(map[string]tls)
}

// todo 实现端口与证书匹配映射
func TLS(context *gin.Context) {

}
