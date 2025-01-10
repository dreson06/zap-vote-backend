package classrep

import "time"

type ClassRep struct {
	ID          string    `db:"id" json:"id"`
	CourseCode  string    `db:"course_code" json:"course_code"`
	CandidateID string    `db:"candidate_id" json:"candidate_id"`
	Slogan      string    `db:"slogan" json:"slogan"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Simple struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	ElectionID string `db:"election_id" json:"election_id"`
	Slogan     string `db:"slogan" json:"slogan"`
	Thumbnail  string `db:"thumbnail" json:"thumbnail"`
	Votes      int64  `db:"votes" json:"votes"`
}

type Results struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	ElectionID  string `db:"election_id" json:"election_id"`
	CandidateID string `db:"candidate_id" json:"candidate_id"`
	Votes       int64  `db:"votes" json:"votes"`
}
