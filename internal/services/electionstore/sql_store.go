package electionstore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zapvote/internal/model/election"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSqlStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{db: db}
}

func (es SQLStore) Create(e *election.Election) error {
	if e.ID == "" {
		e.ID = primitive.NewObjectID().Hex()
	}
	if e.Title == "" || e.StartAt.IsZero() || e.EndAt.IsZero() {
		return errors.New("information missing")
	}
	_, err := es.db.NamedExec("INSERT INTO _election(id,title,election_type,created_at,start_at,end_at) VALUES(:id,:title,:election_type,:created_at,:start_at,:end_at)", e)
	if err != nil {
		return err
	}
	return nil
}

func (es SQLStore) GetElection(id string) (*election.Election, error) {
	var e election.Election
	err := es.db.Get(&e, "SELECT * FROM _election WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (es SQLStore) GetTotalVotes(id string) (int64, error) {
	var votes int64
	err := es.db.Get(&votes, "SELECT total_votes FROM _election WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return votes, nil
}
