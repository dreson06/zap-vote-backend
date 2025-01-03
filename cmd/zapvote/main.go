package main

import (
	"github.com/rs/zerolog/log"
	"zapvote/config"
	"zapvote/internal/api"
	"zapvote/internal/data"
	"zapvote/internal/services/adminstore"
	"zapvote/internal/services/candidatestore"
	"zapvote/internal/services/electionstore"
	"zapvote/internal/services/userstore"
	"zapvote/internal/services/votestore"
)

func main() {
	cfg := config.Init()
	initGlobalLogger(cfg.Mode.IsRelease())

	db := data.Init(config.Cfg.PostgresURL)
	userService := userstore.NewSqlStore(db)
	adminService := adminstore.NewSqlStore(db)
	candidateService := candidatestore.NewSqlStore(db)
	electionService := electionstore.NewSqlStore(db)
	voteService := votestore.NewSqlStore(db)

	apiConfig := &api.ConfigParams{
		DB:             db,
		Mode:           config.Cfg.Mode,
		UserStore:      userService,
		AdminStore:     adminService,
		CandidateStore: candidateService,
		ElectionStore:  electionService,
		VoteStore:      voteService,
	}
	server := api.Init(apiConfig)

	log.Info().Msg("starting server on port " + config.Cfg.Port)
	log.Fatal().Err(server.Start(":" + config.Cfg.Port)).Send()
}
