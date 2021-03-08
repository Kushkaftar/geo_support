package service

import (
	"github.com/Kushkaftar/geo_support/modelsstruct"
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
func (s *ReturnPriceService) GetAllPrices() ([]modelsstruct.Price, error) {
	return s.repo.GetAllPrices()
}

// UpdatePrice ...
func (s *ReturnPriceService) UpdatePrice(pr modelsstruct.Price) (int, error) {
	return s.repo.UpdatePrice(pr)
}
