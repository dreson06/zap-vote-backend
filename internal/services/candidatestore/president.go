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
	query := `SELECT p.id,p.slogan,p.votes,c.thumbnail as president_thumbnail,v.thumbnail as vice_thumbnail,c.course_code as president_course,v.course_code as vice_course,c.name as president_name,v.name as vice_name,c.department FROM _presidential p JOIN _election e ON e.id = p.election_id JOIN _candidate c ON c.id = p.president_id JOIN _candidate v ON v.id = p.vice_id WHERE p.id= $1;`
	err := cs.db.Get(c, query, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
