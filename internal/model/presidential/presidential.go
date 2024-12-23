package presidential

import "time"

type Presidential struct {
	ID          string    `db:"id" json:"id"`
	ElectionID  string    `db:"election_id" json:"election_id"`
	PresidentID string    `db:"president_id" json:"president_id"`
	ViceID      string    `db:"vice_id" json:"vice_id"`
	Slogan      string    `db:"slogan" json:"slogan"`
	Votes       int       `db:"votes" json:"votes"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Simple struct {
	ID            string `db:"id" json:"id"`
	PresidentName string `db:"president_name" json:"president_name"`
	ViceName      string `db:"vice_name" json:"vice_name"`
	Slogan        string `db:"slogan" json:"slogan"`
	Department    string `db:"department" json:"department"`
}
