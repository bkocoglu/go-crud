package api

import (
	"github.com/bilalkocoglu/go-crud/pkg/mw"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(g *gin.RouterGroup) {
	g.GET("/", mw.BasicAuth(), hello)
	g.GET("/a", mw.JwtAuth(), helloA)
}
