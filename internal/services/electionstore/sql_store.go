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

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{db: db}
}

func (es SQLStore) Create(e *election.Election) error {
	if e.ID == "" {
		e.ID = primitive.NewObjectID().Hex()
	}
	if e.Title == "" || e.StartAt.IsZero() || e.EndAt.IsZero() {
		return errors.New("information missing")
	}
	_, err := es.db.NamedExec("INSERT INTO _election(id,title,created_at,start_at,end_at) VALUES(:id,:title,:created_at,:start_at,:end_at)", e)
	if err != nil {
		return err
	}
	return nil
}
