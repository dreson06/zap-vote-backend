package votestore

import (
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/vote"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSqlStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (vs SQLStore) CreateTx(tx *sqlx.Tx, v *vote.Vote) error {
	if v.UserID == "" || v.ElectionID == "" {
		return ErrorInformationMissing
	}
	_, err := tx.NamedQuery("INSERT INTO _votes(user_id,election_id) VALUES(:user_id,:election_id)", v)
	if err != nil {
		return err
	}
	return nil
}

func (vs SQLStore) HasVoted(userID, electionID string) (bool, error) {
	var count int
	err := vs.db.Get(&count, "SELECT COUNT(*) FROM _votes WHERE user_id=$1 AND election_id=$2", userID, electionID)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
