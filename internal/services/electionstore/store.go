package electionstore

import (
	"zapvote/internal/model/election"
	"zapvote/internal/model/presidential"
)

type Store interface {
	Create(e *election.Election) error
	GetPresidentialCandidates() ([]presidential.Simple, error)
}
