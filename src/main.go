package main

import (
	"github.com/bilalkocoglu/go-crud/pkg/config"
	"github.com/bilalkocoglu/go-crud/pkg/database"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.ApplicationConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Load config failed")
	}

	e := config.PrepareServer(cfg)
	database.DB, err = gorm.Open(mysql.Open(database.DbURL(database.BuildDBConfig())), &gorm.Config{})
	database.Migration()

	log.Info().Str("addr", cfg.Addr).Msg("starting http listener")
	err = e.Run(cfg.Addr)
	log.Fatal().Err(err).Msg("Server failed")
}
