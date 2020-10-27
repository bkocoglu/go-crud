package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveUser(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
