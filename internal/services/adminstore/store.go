package adminstore

import "zapvote/internal/model/admin"

type Store interface {
	Create(a *admin.Admin) error
	GetMe(username string) (*admin.Admin, error)
}
