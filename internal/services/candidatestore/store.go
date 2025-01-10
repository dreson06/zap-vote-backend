package candidatestore

import (
	"zapvote/internal/model/candidate"
	"zapvote/internal/model/classrep"
	"zapvote/internal/model/faculty"
	"zapvote/internal/model/presidential"
)

type Store interface {
	Create(c *candidate.Candidate) error
	GetCandidateByDepartment(department string) ([]candidate.Candidate, error)

	GetFacultyCandidateByID(id string) (*faculty.Simple, error)
	GetClassRepByID(id string) (*classrep.Simple, error)

	GetPresidentialOne(id string) (*presidential.Candidate, error)
}
