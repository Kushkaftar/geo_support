package repository

import (
	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/jmoiron/sqlx"
)

// Domains ...
type Domains interface {
	GetAllDomains() ([]modelsstruct.Domain, error)
	InsertDomain(domain string) error
	CheckDomain(domain string) (int, error)
	SetFlag(flag, id int) error
}

// Prices ...
type Prices interface {
	GetAllPrices() ([]modelsstruct.Price, error)
	UpdatePrice(pr modelsstruct.Price) (int, error)
	SetPrice(pr modelsstruct.Price) (int, error)
}

// Offer ...
type Offer interface {
	SetOfferName(n modelsstruct.Name) error
	GetOfferName(id int) (string, error)
}

//Repository ...
type Repository struct {
	Domains
	Offer
	Prices
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Domains: NewDomainsMysql(db),
		Offer:   NewOfferMysql(db),
		Prices:  NewPricesMysql(db),
	}
}
