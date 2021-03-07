package service

import (
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
)

// ReturnPriceService ...
type ReturnPriceService struct {
	repo repository.Prices
}

// NewReturnPriceService ...
func NewReturnPriceService(repo repository.Prices) *ReturnPriceService {
	return &ReturnPriceService{repo: repo}
}

// GetAllPrices ...
func (s *ReturnPriceService) GetAllPrices() (modelsStruct.Prices, error) {
	return s.repo.GetAllPrices()
}
