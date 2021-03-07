package service

import (
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
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

type Service struct {
	Domains
	Offer
	Prices
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Domains: NewReturnDomainsService(repos.Domains),
		Offer:   NewReturnOfferService(repos.Offer),
		Prices:  NewReturnPriceService(repos.Prices),
	}
}
