package votestore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/vote"
)

var ErrorInformationMissing = errors.New("information missing")

type Store interface {
	CreateTx(tx *sqlx.Tx, v *vote.Vote) error
	HasVoted(userID string, voteType vote.Category) (bool, error)
}
