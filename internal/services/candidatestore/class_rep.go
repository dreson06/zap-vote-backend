package candidatestore

import "zapvote/internal/model/classrep"

func (cs *SQLStore) GetClassRepByID(id string) (*classrep.Simple, error) {
	res := &classrep.Simple{}
	query := `SELECT cr.id,cr.slogan,cr.votes,c.name,c.thumbnail FROM _classrep cr JOIN _candidate c ON c.id = cr.candidate_id WHERE cr.candidate_id = $1;`
	err := cs.db.Get(res, query, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
