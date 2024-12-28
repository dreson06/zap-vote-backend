package faculty

import "time"

type Faculty struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	CandidateID string    `db:"candidate_id" json:"candidate_id"`
	Votes       int64     `db:"votes" json:"votes"`
	Slogan      string    `db:"slogan" json:"slogan"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Simple struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Slogan     string `db:"slogan" json:"slogan"`
	Thumbnail  string `db:"thumbnail" json:"thumbnail"`
	Votes      int64  `db:"votes" json:"votes"`
	CourseCode string `db:"course_code" json:"course_code"`
}
