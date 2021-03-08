package service

import (
	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
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
}

// Offer ...
type Offer interface {
	SetOfferName(n modelsstruct.Name) error
	GetOfferName(id int) (string, error)
}

// Service ...
type Service struct {
	Domains
	Offer
	Prices
}

// NewService ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Domains: NewReturnDomainsService(repos.Domains),
		Offer:   NewReturnOfferService(repos.Offer),
		Prices:  NewReturnPriceService(repos.Prices),
	}
}
