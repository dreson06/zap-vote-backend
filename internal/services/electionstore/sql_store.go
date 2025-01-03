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

func (es SQLStore) GetElection(et election.TypeElection) (*election.Election, error) {
	var e election.Election
	err := es.db.Get(&e, "SELECT * FROM _election WHERE election_type=$1", et)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
