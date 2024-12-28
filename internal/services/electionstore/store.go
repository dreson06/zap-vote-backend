package electionstore

import (
	"zapvote/internal/model/classrep"
	"zapvote/internal/model/election"
	"zapvote/internal/model/faculty"
	"zapvote/internal/model/presidential"
)

type Store interface {
	Create(e *election.Election) error
	GetPresidentialCandidates() ([]presidential.Simple, error)
	GetFacultyCandidates(faculty string) ([]faculty.Simple, error)
	GetClassRepCandidates(courseCode string) ([]classrep.Simple, error)
}
