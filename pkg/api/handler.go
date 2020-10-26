package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterHandlers(g *echo.Group) {
	g.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
