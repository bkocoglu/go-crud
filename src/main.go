package main

import (
	"github.com/bilalkocoglu/go-crud/pkg/config"
	"github.com/bilalkocoglu/go-crud/pkg/entity"
	"github.com/pkg/errors"
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
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	err = config.DB.AutoMigrate(&entity.User{})

	if err != nil {
		errors.Wrap(err, "Db migration error !")
	}

	log.Info().Str("addr", cfg.Addr).Msg("starting http listener")
	err = e.Run(cfg.Addr)
	log.Fatal().Err(err).Msg("Server failed")
}
