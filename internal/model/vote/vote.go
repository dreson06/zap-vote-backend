package vote

type Category int

const (
	Presidential Category = 1
	Banking      Category = 2
	Finance      Category = 3
	Computer     Category = 4
	Class        Category = 5
)

type Vote struct {
	UserID   string   `db:"user_id" json:"user_id"`
	VoteType Category `db:"vote_type" json:"vote_type"`
}
