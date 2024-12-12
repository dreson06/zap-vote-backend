package userstore

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
	"zapvote/internal/model/user"
)

type SqlStore struct {
	db *sqlx.DB
}

func NewSqlStore(db *sqlx.DB) *SqlStore {
	return &SqlStore{db: db}
}

func (us *SqlStore) Create(u *user.User) error {
	if u.ID == "" || u.DeviceID == "" {
		return errors.New("information missing")
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
		u.UpdatedAt = u.CreatedAt
	}
	_, err := us.db.NamedQuery("INSERT INTO _user(id, device_id,password, course_code, created_at, updated_at) VALUES (:id,:device_id,:password, :course_code, :created_at, :updated_at)", u)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrorUserExists
		}
		return err
	}
	return nil
}

func (us *SqlStore) GetUserSimple(id string) (*user.Simple, error) {
	u := &user.Simple{}
	err := us.db.Get(u, "SELECT id,password FROM _user WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func (us *SqlStore) GetMe(id string) (*user.MeData, error) {
	u := &user.MeData{}
	err := us.db.Get(u, "SELECT id,course_code FROM _user WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
