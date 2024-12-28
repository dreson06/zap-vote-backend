package electionstore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zapvote/internal/model/election"
	"zapvote/internal/model/faculty"
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
	query := `SELECT p.id,p.slogan,p.votes,c.name as president_name,v.name as vice_name,c.department FROM _presidential p JOIN _election e ON e.id = p.election_id JOIN _candidate c ON c.id = p.president_id JOIN _candidate v ON v.id = p.vice_id ORDER BY p.votes DESC;`
	err := es.db.Select(&candidates, query)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetFacultyCandidates(name string) ([]faculty.Simple, error) {
	candidates := make([]faculty.Simple, 0)
	query := `SELECT f.id,f.slogan,f.votes,c.name,c.course_code,c.thumbnail FROM _faculty f JOIN _candidate c ON f.candidate_id = c.id WHERE f.name = $1`
	err := es.db.Select(&candidates, query, name)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetClassElection(courseCode string) {
	panic("not implemented")
}
