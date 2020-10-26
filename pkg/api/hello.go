package api

import (
	"github.com/bilalkocoglu/go-crud/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func helloA(c echo.Context) error {
	m := &model.Response{
		Message: "Hello world A!",
	}
	return c.JSON(http.StatusOK, m)
}
