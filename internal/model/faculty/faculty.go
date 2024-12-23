package faculty

import "time"

type Faculty struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	CandidateID string    `db:"candidate_id" json:"candidate_id"`
	Slogan      string    `db:"slogan" json:"slogan"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
