package votestore

import (
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/vote"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (vs SQLStore) CreateTx(tx *sqlx.Tx, v *vote.Vote) error {
	if v.UserID == "" || v.VoteType == 0 {
		return ErrorInformationMissing
	}
	_, err := tx.NamedQuery("INSERT INTO _votes(user_id,vote_type) VALUES(:user_id,:vote_type)", v)
	if err != nil {
		return err
	}
	return nil
}

func (vs SQLStore) HasVoted(userID string, voteType vote.Category) (bool, error) {
	var count int
	err := vs.db.Get(&count, "SELECT COUNT(*) FROM _votes WHERE user_id=$1 AND vote_type=$2", userID, voteType)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
