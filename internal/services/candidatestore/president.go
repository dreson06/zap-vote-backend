package candidatestore

import (
	"zapvote/internal/model/candidate"
	"zapvote/internal/model/presidential"
)

func (cs *SQLStore) GetCandidateByDepartment(department string) ([]candidate.Candidate, error) {
	candidates := make([]candidate.Candidate, 0)
	err := cs.db.Select(&candidates, "SELECT * FROM _candidate WHERE department=$1", department)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (cs *SQLStore) GetPresidentialOne(id string) (*presidential.Candidate, error) {
	c := &presidential.Candidate{}
	query := `SELECT p.id,p.election_id,p.slogan,p.votes,c.name as president_name,c.course_code as president_course,c.thumbnail as president_thumbnail,v.name as vice_name,v.course_code as vice_course,v.thumbnail as vice_thumbnail,c.department FROM _presidential p JOIN _election e ON e.id = p.election_id JOIN _candidate c ON c.id = p.president_id JOIN _candidate v ON v.id = p.vice_id WHERE p.id= $1;`
	err := cs.db.Get(c, query, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
