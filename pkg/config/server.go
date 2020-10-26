package config

import (
	"github.com/bilalkocoglu/go-crud/pkg/api"
	"github.com/bilalkocoglu/go-crud/pkg/mw"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg *Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func PrepareServer(config *Config) *gin.Engine {
	router := gin.Default()
	router.Use(mw.CORSMiddleware())

	g := router.Group("/v1")

	mw.SetInterceptors(g)
	api.RegisterHandlers(g)

	return router
}
