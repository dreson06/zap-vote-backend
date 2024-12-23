package classrep

import "time"

type ClassRep struct {
	ID          string    `db:"id" json:"id"`
	CourseCode  string    `db:"course_code" json:"course_code"`
	CandidateID string    `db:"candidate_id" json:"candidate_id"`
	Slogan      string    `db:"slogan" json:"slogan"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
