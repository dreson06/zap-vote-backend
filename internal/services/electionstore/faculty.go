package electionstore

import (
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/faculty"
)

func (es SQLStore) GetFacultyResults() ([]faculty.Results, error) {
	var candidates []faculty.Results
	query := `SELECT f.id,e.id as election_id,f.candidate_id, f.name as faculty_name,f.votes,c.name as candidate_name FROM _faculty f JOIN _candidate c ON f.candidate_id = c.id JOIN _election e ON e.id = f.election_id ORDER BY f.votes DESC`
	err := es.db.Select(&candidates, query)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetFacultyCandidates(name string) ([]faculty.Simple, error) {
	candidates := make([]faculty.Simple, 0)
	query := `SELECT f.id,f.election_id,f.slogan,f.votes,c.name,c.course_code,c.thumbnail FROM _faculty f JOIN _candidate c ON f.candidate_id = c.id WHERE f.name = $1`
	err := es.db.Select(&candidates, query, name)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetFacultyVotesTx(tx *sqlx.Tx, id string) (int64, error) {
	var votes int64
	err := tx.Get(&votes, "SELECT votes FROM _faculty WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return votes, nil
}

func (es SQLStore) UpdateFacultyVote(tx *sqlx.Tx, id string, vote int64) error {
	_, err := tx.Exec("UPDATE _faculty SET votes = $1 WHERE id=$2", vote, id)
	return err
}
