package repository

import (
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/jmoiron/sqlx"
)

type Domains interface {
	GetAllDomains() ([]modelsStruct.Domain, error)
	InsertDomain(domain string) error
	CheckDomain(domain string) (int, error)
	SetFlag(flag, id int) error
}

type Prices interface {
	GetAllPrices() (modelsStruct.Prices, error)
}

type Offer interface {
	SetOfferName(n modelsStruct.Name) error
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
