package electionstore

import (
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/classrep"
	"zapvote/internal/model/election"
	"zapvote/internal/model/faculty"
	"zapvote/internal/model/presidential"
)

type Store interface {
	Create(e *election.Election) error
	GetElection(et election.TypeElection) (*election.Election, error)

	GetAllPresidentialCandidates() ([]presidential.Candidate, error)
	GetFacultyCandidates(faculty string) ([]faculty.Simple, error)
	GetClassRepCandidates(courseCode string) ([]classrep.Simple, error)

	GetFacultyResults() ([]faculty.Results, error)
	GetClassRepResults(code string) ([]classrep.Results, error)

	UpdatePresidentVoteTx(tx *sqlx.Tx, id string, vote int64) error
	UpdateFacultyVote(tx *sqlx.Tx, id string, vote int64) error
	UpdateClassRepVote(tx *sqlx.Tx, id string, vote int64) error

	GetClassVotes(tx *sqlx.Tx, id string) (int64, error)
	GetPresidentVotesTx(tx *sqlx.Tx, id string) (int64, error)
	GetFacultyVotesTx(tx *sqlx.Tx, id string) (int64, error)
}
