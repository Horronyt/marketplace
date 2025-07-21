package repository

import (
	"fmt"
	"github.com/Horronyt/marketplace"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user marketplace.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, salt) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Salt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserSalt(username string) (string, error) {
	var user marketplace.User
	query := fmt.Sprintf("SELECT salt FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)

	return user.Salt, err
}

func (r *AuthPostgres) GetUser(username, password string) (marketplace.User, error) {
	var user marketplace.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
