package electionstore

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zapvote/internal/model/election"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{db: db}
}

func (es *SQLStore) CreatePresidentialElection(p *election.PresidentialElection) error {
	if p.ID == "" {
		p.ID = primitive.NewObjectID().Hex()
	}
	if p.PresidentID == "" || p.ViceID == "" || p.Slogan == "" {
		return errors.New("information missing")
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
		p.UpdatedAt = p.CreatedAt
	}
	_, err := es.db.NamedExec("INSERT INTO _presidential_election(id, president_id, vice_id, slogan,created_at,updated_at)VALUES(:id,:president_id,:vice_id,:slogan,:created_at,:updated_at)", p)
	if err != nil {
		return err
	}
	return nil
}

func (es *SQLStore) CreateFacultyElection(f *election.FacultyElection) error {
	if f.ID == "" {
		f.ID = primitive.NewObjectID().Hex()
	}
	if f.Name == "" || f.CandidateID == "" || f.Slogan == "" {
		return errors.New("information missing")
	}
	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now()
		f.UpdatedAt = f.CreatedAt
	}
	_, err := es.db.NamedExec("INSERT INTO _faculty_election(id, candidate_id, name, slogan,created_at,updated_at)VALUES(:id,:candidate_id,:name,:slogan,:created_at,:updated_at)", f)
	if err != nil {
		return err
	}
	return nil
}

func (es *SQLStore) CreateSpecificElection(s *election.SpecialElection) error {
	if s.ID == "" {
		s.ID = primitive.NewObjectID().Hex()
	}
	if s.CourseCode == "" || s.CandidateID == "" || s.Slogan == "" {
		return errors.New("information missing")
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
		s.UpdatedAt = s.CreatedAt
	}
	_, err := es.db.NamedExec("INSERT INTO _special_election(id, candidate_id, course_code, slogan,created_at,updated_at)VALUES(:id,:candidate_id,:course_code,:slogan,:created_at,:updated_at)", s)
	if err != nil {
		return err
	}
	return nil
}
