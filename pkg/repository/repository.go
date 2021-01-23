package repository

import (
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/jmoiron/sqlx"
)

type Domains interface {
	GetAllDomains() ([]modelsStruct.Domain, error)
	InsertDomain(domain string) error
	CheckDomain(domain string) (int, error)
}

type Prices interface {
}

type Repository struct {
	Domains
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Domains: NewDomainsMysql(db),
	}
}
