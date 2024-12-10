package main

import (
	"github.com/rs/zerolog/log"
	"zapvote/config"
	"zapvote/internal/api"
	"zapvote/internal/data"
)

func main() {
	cfg := config.Init()
	initGlobalLogger(cfg.Mode.IsRelease())

	db := data.Init(config.Cfg.PostgresURL)
	apiConfig := &api.ConfigParams{
		DB:   db,
		Mode: config.Cfg.Mode,
	}
	server := api.Init(apiConfig)

	log.Info().Msg("starting server on port " + config.Cfg.Port)
	log.Fatal().Err(server.Start(":" + config.Cfg.Port)).Send()
}
