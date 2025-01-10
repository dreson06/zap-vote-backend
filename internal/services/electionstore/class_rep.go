package electionstore

import (
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/classrep"
)

func (es SQLStore) GetClassRepResults(code string) ([]classrep.Results, error) {
	var candidates []classrep.Results
	query := `SELECT cr.id,e.id as election_id,cr.candidate_id,cr.votes,c.name FROM _classrep cr JOIN _candidate c ON c.id = cr.candidate_id JOIN _election e ON e.id = cr.election_id WHERE cr.course_code=$1 ORDER BY cr.votes DESC`
	err := es.db.Select(&candidates, query, code)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetClassRepCandidates(courseCode string) ([]classrep.Simple, error) {
	candidates := make([]classrep.Simple, 0)
	query := `SELECT cr.id,cr.election_id,c.name,cr.slogan,cr.votes,c.thumbnail FROM _classrep cr LEFT JOIN _candidate c ON cr.candidate_id = c.id WHERE cr.course_code=$1 ORDER BY cr.votes DESC`
	err := es.db.Select(&candidates, query, courseCode)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) UpdateClassRepVote(tx *sqlx.Tx, id string, vote int64) error {
	_, err := tx.Exec("UPDATE _classrep SET votes = $1 WHERE id=$2", vote, id)
	return err
}

func (es SQLStore) GetClassVotes(tx *sqlx.Tx, id string) (int64, error) {
	var votes int64
	err := tx.Get(&votes, "SELECT votes FROM _classrep WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return votes, nil
}
