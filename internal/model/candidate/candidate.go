package candidate

import "time"

type Position string
type Type string

const (
	President     Position = "president"
	VicePresident Position = "vice-president"
	FR            Position = "faculty-representative"
	CR            Position = "class-representative"
)

const (
	Presidential Type = "presidential"
	Specific     Type = "specific"
	Faculty      Type = "faculty"
)

type Candidate struct {
	ID         string    `db:"id" json:"id,omitempty"`
	Name       string    `db:"name" json:"name"`
	ElectionID string    `db:"election_id" json:"election_id"`
	CourseCode string    `db:"course_code" json:"course_code"`
	Position   Position  `db:"position" json:"position"`
	Type       Type      `db:"type" json:"type"`
	Thumbnail  string    `db:"thumbnail" json:"thumbnail"`
	Department string    `db:"department" json:"department,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

func New() *Candidate {
	return &Candidate{}
}
