package userstore

import (
	"errors"
	"zapvote/internal/model/user"
)

var ErrorUserExists = errors.New("user exists")
var ErrorUserNotFound = errors.New("user not found")

type Store interface {
	Create(u *user.User) error
	GetUserSimple(id string) (*user.Simple, error)
	GetMe(id string) (*user.MeData, error)
}
