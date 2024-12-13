package candidatestore

import "zapvote/internal/model/candidate"

type Store interface {
	Create(c *candidate.Candidate) error
	GetSpecific(courseCode string) ([]candidate.Candidate, error)
	GetGeneral() ([]candidate.Candidate, error)
}
