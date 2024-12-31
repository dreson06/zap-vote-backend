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
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Slogan      string `db:"slogan" json:"slogan,omitempty"`
	Thumbnail   string `db:"thumbnail" json:"thumbnail"`
	FacultyName string `db:"faculty_name" json:"faculty_name,omitempty"`
	CourseCode  string `db:"course_code" json:"course_code"`
	Votes       int64  `db:"votes" json:"votes"`
}

type Results struct {
	ID            string `db:"id" json:"id"`
	CandidateID   string `db:"candidate_id" json:"candidate_id"`
	FacultyName   string `db:"faculty_name" json:"faculty_name"`
	CandidateName string `db:"candidate_name" json:"candidate_name"`
	Votes         int64  `db:"votes" json:"votes"`
}
