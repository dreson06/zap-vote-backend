package candidatestore

import "zapvote/internal/model/faculty"

func (cs *SQLStore) GetFacultyCandidateByID(id string) (*faculty.Simple, error) {
	res := &faculty.Simple{}
	query := `SELECT f.id,f.name as faculty_name,f.slogan,f.votes,c.name,c.course_code,c.thumbnail FROM _faculty f JOIN _candidate c ON f.candidate_id = c.id WHERE f.candidate_id = $1;`
	err := cs.db.Get(res, query, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
