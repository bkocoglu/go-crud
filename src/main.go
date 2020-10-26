package main

import (
	"fmt"
	"github.com/bilalkocoglu/go-crud/pkg/api"
	"github.com/bilalkocoglu/go-crud/pkg/config"
	"github.com/bilalkocoglu/go-crud/pkg/middleware"
	"github.com/bilalkocoglu/go-crud/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"io/ioutil"
)

func main() {
	cfg, err := config.ApplicationConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Load config failed")
	}

	server, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed server")
	}

	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	g := e.Group("/v1")

	middleware.SetInterceptors(g)
	api.RegisterHandlers(g)

	log.Info().Msg(fmt.Sprintf("%#v", server))
	log.Info().Str("addr", cfg.Addr).Msg("starting http listener")
	err = e.Start(cfg.Addr)
	log.Fatal().Err(err).Msg("Server failed")
}
