package presidential

type Candidate struct {
	ID                 string `db:"id" json:"id"`
	PresidentName      string `db:"president_name" json:"president_name"`
	ElectionID         string `db:"election_id" json:"election_id"`
	ViceName           string `db:"vice_name" json:"vice_name"`
	Slogan             string `db:"slogan" json:"slogan"`
	Department         string `db:"department" json:"department"`
	PresidentThumbnail string `db:"president_thumbnail" json:"president_thumbnail,omitempty"`
	ViceThumbnail      string `db:"vice_thumbnail" json:"vice_thumbnail,omitempty"`
	PresidentCourse    string `db:"president_course" json:"president_course,omitempty"`
	ViceCourse         string `db:"vice_course" json:"vice_course,omitempty"`
	Votes              int    `db:"votes" json:"votes"`
}

type Simple struct {
	ID            string `db:"id" json:"id"`
	PresidentName string `db:"president_name" json:"president_name"`
	ElectionID    string `db:"election_id" json:"election_id"`
	ViceName      string `db:"vice_name" json:"vice_name"`
	Slogan        string `db:"slogan" json:"slogan"`
	Department    string `db:"department" json:"department"`
	Votes         int    `db:"votes" json:"votes"`
}
