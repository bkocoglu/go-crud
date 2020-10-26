package main

import (
	"github.com/bilalkocoglu/go-crud/pkg/api"
	"github.com/bilalkocoglu/go-crud/pkg/config"
	"github.com/bilalkocoglu/go-crud/pkg/mw"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"io/ioutil"
)

func main() {
	cfg, err := config.ApplicationConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Load config failed")
	}

	/*
		server, err := server.NewServer(cfg)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed server")
		}
	*/

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))

	e.Logger.SetOutput(ioutil.Discard)
	g := e.Group("/v1")

	mw.SetInterceptors(g)
	api.RegisterHandlers(g)

	log.Info().Str("addr", cfg.Addr).Msg("starting http listener")
	err = e.Start(cfg.Addr)
	log.Fatal().Err(err).Msg("Server failed")
}
