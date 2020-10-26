package main

import (
	"fmt"
	"github.com/bilalkocoglu/go-crud/pkg/config"
	"github.com/bilalkocoglu/go-crud/pkg/server"
	"github.com/rs/zerolog/log"
	"net/http"
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

	log.Info().Msg(fmt.Sprintf("%#v", server))
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":10000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
