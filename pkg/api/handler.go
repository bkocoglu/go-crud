package api

import (
	"github.com/bilalkocoglu/go-crud/pkg/mw"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/", hello, mw.BasicAuth)
	g.GET("/a", helloA, mw.JwtAuth)
}
