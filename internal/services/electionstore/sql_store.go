package electionstore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zapvote/internal/model/election"
	"zapvote/internal/model/presidential"
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

func (es SQLStore) GetPresidentialCandidates() ([]presidential.Simple, error) {
	candidates := make([]presidential.Simple, 0)
	query := `SELECT p.id,p.slogan,c.name as president_name,v.name as vice_name,c.department FROM _presidential p JOIN _election e ON e.id = p.election_id JOIN _candidate c ON c.id = p.president_id JOIN _candidate v ON v.id = p.vice_id;`
	err := es.db.Select(&candidates, query)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetFacultyElection() {
	panic("not implemented")
}

func (es SQLStore) GetClassElection(courseCode string) {
	panic("not implemented")
}
