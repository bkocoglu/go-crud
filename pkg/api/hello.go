package api

import (
	"github.com/bilalkocoglu/go-crud/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func HelloA(c *gin.Context) {
	m := &model.Response{
		Message: "Hello world A!",
	}
	c.JSON(http.StatusOK, m)
}
