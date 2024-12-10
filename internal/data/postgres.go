package data

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var db *sqlx.DB

func Init(url string) *sqlx.DB {
	var err error
	log.Info().Msg("connecting to postgresql database")
	db, err = sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatal().Err(err).Msg("postgresql connection failed")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("postgresql connection failed")
	}
	log.Info().Msg("connected to postgresql database")
	return db
}
