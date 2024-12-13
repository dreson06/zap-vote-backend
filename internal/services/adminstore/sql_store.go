package adminstore

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zapvote/internal/model/admin"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (as *SQLStore) Create(a *admin.Admin) error {
	if a.ID == "" {
		a.ID = primitive.NewObjectID().Hex()
	}
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now()
		a.UpdatedAt = a.CreatedAt
	}
	_, err := as.db.NamedExec("INSERT INTO _admin (id,username,password,created_at,updated_at)VALUES (:id,:username,:password,:created_at,:updated_at)", a)
	if err != nil {
		return err
	}
	return nil
}

func (as *SQLStore) GetMe(username string) (*admin.Admin, error) {
	a := admin.New()
	err := as.db.Get(a, "SELECT * FROM _admin WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return a, nil
}
