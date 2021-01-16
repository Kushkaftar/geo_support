package service

import "github.com/Kushkaftar/geo_support/pkg/repository"

type Domains struct {

}

type Prices struct {

}

type Service struct {
	Domains
	Prices
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}