package candidatestore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zapvote/internal/model/candidate"
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
	err := cs.db.Select(&candidates, "SELECT * FROM _candidate WHERE type=$1", candidate.General)
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