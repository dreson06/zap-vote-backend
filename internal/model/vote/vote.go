package vote

type Vote struct {
	UserID     string `db:"user_id" json:"user_id"`
	ElectionID string `db:"election_id" json:"election_id"`
}
