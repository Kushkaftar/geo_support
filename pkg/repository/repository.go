package repository

import (
	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/jmoiron/sqlx"
)

type Domains interface {
	GetAllDomains() ([]modelsstruct.Domain, error)
	InsertDomain(domain string) error
	CheckDomain(domain string) (int, error)
	SetFlag(flag, id int) error
}

type Prices interface {
	GetAllPrices() ([]modelsstruct.Price, error)
}

type Offer interface {
	SetOfferName(n modelsstruct.Name) error
	GetOfferName(id int) (string, error)
}

type Repository struct {
	Domains
	Offer
	Prices
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Domains: NewDomainsMysql(db),
		Offer:   NewOfferMysql(db),
		Prices:  NewPricesMysql(db),
	}
}
