package electionstore

import "zapvote/internal/model/election"

type Store interface {
	CreatePresidentialElection(p *election.PresidentialElection) error
	CreateFacultyElection(f *election.FacultyElection) error
	CreateSpecificElection(s *election.SpecialElection) error

	GetPresidential() ([]election.PresidentialCandidates, error)
}
