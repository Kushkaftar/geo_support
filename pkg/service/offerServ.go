package service

import (
	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
)

type ReturnOfferService struct {
	repo repository.Offer
}

func NewReturnOfferService(repo repository.Offer) *ReturnOfferService {
	return &ReturnOfferService{repo: repo}
}

func (ros ReturnOfferService) SetOfferName(n modelsstruct.Name) error {
	return ros.SetOfferName(n)
}

func (ros ReturnOfferService) GetOfferName(id int) (string, error) {
	return ros.GetOfferName(id)
}
