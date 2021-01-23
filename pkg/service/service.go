package service

import (
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
)

type Domains interface {
	GetAllDomains() ([]modelsStruct.Domain, error)
	InsertDomain(domain string) error
	CheckDomain(domain string) (int, error)
}

type Prices interface {
}

type Service struct {
	Domains
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Domains: NewReturnDomainsService(repos.Domains),
	}
}
