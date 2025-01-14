package electionstore

import (
	"github.com/jmoiron/sqlx"
	"zapvote/internal/model/presidential"
)

func (es SQLStore) GetPresidentialCandidates() ([]presidential.Simple, error) {
	candidates := make([]presidential.Simple, 0)
	query := `SELECT p.id,p.election_id,c.department,c.name as president_name,v.name as vice_name,p.slogan,p.votes FROM _presidential p JOIN _candidate c ON c.id = p.president_id JOIN _candidate v on v.id = p.vice_id ORDER BY p.votes DESC`
	err := es.db.Select(&candidates, query)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (es SQLStore) GetPresidentVotesTx(tx *sqlx.Tx, id string) (int64, error) {
	var votes int64
	err := tx.Get(&votes, "SELECT votes FROM _presidential WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return votes, nil
}

func (es SQLStore) UpdatePresidentVoteTx(tx *sqlx.Tx, id string, vote int64) error {
	_, err := tx.Exec("UPDATE _presidential SET votes = $1 WHERE id=$2", vote, id)
	return err
}
