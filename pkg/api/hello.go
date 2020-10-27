package api

import (
	_const "github.com/bilalkocoglu/go-crud/pkg/const"
	"github.com/bilalkocoglu/go-crud/pkg/database"
	"github.com/bilalkocoglu/go-crud/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	value, exists := c.Get(_const.CurrentUser)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User must be not null",
		})
		return
	}
	user := value.(database.User)

	c.String(http.StatusOK, "Hello, World ! => "+user.Name)
}

func HelloA(c *gin.Context) {
	m := &model.Response{
		Message: "Hello world A!",
	}
	c.JSON(http.StatusOK, m)
}
