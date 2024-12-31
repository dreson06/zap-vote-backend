package candidatestore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zapvote/internal/model/candidate"
	"zapvote/internal/model/classrep"
	"zapvote/internal/model/faculty"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (cs *SQLStore) Create(c *candidate.Candidate) error {
	if c.ID == "" {
		c.ID = primitive.NewObjectID().Hex()
	}
	if c.Name == "" || c.CourseCode == "" || c.Position == "" || c.Type == "" {
		return errors.New("information missing")
	}
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
		c.UpdatedAt = c.CreatedAt
	}

	_, err := cs.db.NamedExec("INSERT INTO _candidate (id, name, course_code, position, type, department, thumbnail, created_at, updated_at) VALUES (:id,:name,:course_code,:position,:type,:department,:thumbnail,:created_at,:updated_at)", c)
	if err != nil {
		return err
	}
	return nil
}

func (cs *SQLStore) GetGeneral() ([]candidate.Candidate, error) {
	candidates := make([]candidate.Candidate, 0)
	err := cs.db.Select(&candidates, "SELECT * FROM _candidate WHERE type=$1", candidate.Presidential)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (cs *SQLStore) GetSpecific(courseCode string) ([]candidate.Candidate, error) {
	candidates := make([]candidate.Candidate, 0)
	err := cs.db.Select(&candidates, "SELECT * FROM _candidate WHERE type=$1 AND course_code=$2", candidate.Specific, courseCode)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (cs *SQLStore) GetCandidateByDepartment(department string) ([]candidate.Candidate, error) {
	candidates := make([]candidate.Candidate, 0)
	err := cs.db.Select(&candidates, "SELECT * FROM _candidate WHERE department=$1", department)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (cs *SQLStore) GetFacultyCandidateByID(id string) (*faculty.Simple, error) {
	res := &faculty.Simple{}
	query := `SELECT f.id,f.name as faculty_name,f.slogan,f.votes,c.name,c.course_code,c.thumbnail FROM _faculty f JOIN _candidate c ON f.candidate_id = c.id WHERE f.candidate_id = $1;`
	err := cs.db.Get(res, query, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cs *SQLStore) GetClassRepByID(id string) (*classrep.Simple, error) {
	res := &classrep.Simple{}
	query := `SELECT cr.id,cr.slogan,cr.votes,c.name,c.thumbnail FROM _classrep cr JOIN _candidate c ON c.id = cr.candidate_id WHERE cr.candidate_id = $1;`
	err := cs.db.Get(res, query, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
