package election

import "time"

type Election struct {
	ID         string    `db:"id" json:"id"`
	Title      string    `db:"title" json:"title"`
	TotalVotes int64     `db:"total_votes" json:"total_votes"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	StartAt    time.Time `db:"start_at" json:"start_at"`
	EndAt      time.Time `db:"end_at" json:"end_at"`
}
