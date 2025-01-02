package user

import "time"

type User struct {
	ID         string    `db:"id" json:"id"`
	DeviceID   string    `db:"device_id" json:"device_id"`
	Password   string    `db:"password" json:"-"`
	CourseCode string    `db:"course_code" json:"course_code"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

type Simple struct {
	ID       string `db:"id" json:"id"`
	Password string `db:"password" json:"-"`
	DeviceID string `db:"device_id" json:"device_id,omitempty"`
}

type MeData struct {
	ID         string `db:"id" json:"id"`
	CourseCode string `db:"course_code" json:"course_code"`
}
