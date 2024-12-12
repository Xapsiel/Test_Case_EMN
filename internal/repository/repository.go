package repository

import "github.com/jmoiron/sqlx"

type User interface {
	Register(nickname, name, email string) error
	UpdateToken(token, email string) error
	VerifyToken(token string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{User: NewUserPostgres(db)}
}

type Repository struct {
	User
}
