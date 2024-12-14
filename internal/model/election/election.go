package election

import "time"

type PresidentialElection struct {
	ID          string    `db:"id" json:"id"`
	PresidentID string    `db:"president_id" json:"president_id"`
	ViceID      string    `db:"vice_id" json:"vice_id"`
	Slogan      string    `db:"slogan" json:"slogan"`
	Votes       uint64    `db:"votes" json:"votes"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type FacultyElection struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"` //faculty name
	CandidateID string    `db:"candidate_id" json:"candidate_id"`
	Slogan      string    `db:"slogan" json:"slogan"`
	Votes       uint64    `db:"votes" json:"votes"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type SpecialElection struct {
	ID          string    `db:"id" json:"id"`
	CandidateID string    `db:"candidate_id" json:"candidate_id"`
	CourseCode  string    `db:"course_code" json:"course_code"`
	Slogan      string    `db:"slogan" json:"slogan"`
	Votes       uint64    `db:"votes" json:"votes"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type PresidentialCandidates struct {
	ID            string `db:"id" json:"id"`
	PresidentName string `db:"president_name" json:"president_name"`
	Department    string `db:"department" json:"department"`
	ViceName      string `db:"vice_name" json:"vice_name"`
	Slogan        string `db:"slogan" json:"slogan"`
}

//SELECT
//pe.id AS id,
//pres.name AS president_name,
//pres.department AS department,
//vice.name AS vice_name,
//pe.slogan
//FROM
//_presidential_election AS pe
//JOIN
//_candidate AS pres
//ON pe.president_id = pres.id
//JOIN
//_candidate AS vice
//ON pe.vice_id = vice.id;
