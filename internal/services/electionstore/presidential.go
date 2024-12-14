package electionstore

import "zapvote/internal/model/election"

func (es *SQLStore) GetPresidential() ([]election.PresidentialCandidates, error) {
	p := make([]election.PresidentialCandidates, 0)
	err := es.db.Select(&p, "SELECT pe.id AS id,pres.name AS president_name,pres.department AS department,vice.name AS vice_name,pe.slogan FROM _presidential_election AS pe JOIN _candidate AS pres ON pe.president_id = pres.id JOIN _candidate AS vice ON pe.vice_id = vice.id;")
	if err != nil {
		return nil, err
	}
	return p, nil
}
