package repository

import (
	"github.com/Horronyt/marketplace"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user marketplace.User) (int, error)
	GetUser(username, password string) (marketplace.User, error)
	GetUserSalt(username string) (string, error)
}

type Listing interface {
}

type Repository struct {
	Authorization
	Listing
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
