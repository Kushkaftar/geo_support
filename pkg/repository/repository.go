package repository

import "github.com/jmoiron/sqlx"

type Domains struct {
}

type Prices struct {
}

type Repository struct {
	Domains
	Prices
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
