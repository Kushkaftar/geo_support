package repository

import (
	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/jmoiron/sqlx"
)

type OfferMysql struct {
	db *sqlx.DB
}

func NewOfferMysql(db *sqlx.DB) *OfferMysql {
	return &OfferMysql{db: db}
}

func (o OfferMysql) SetOfferName(n modelsstruct.Name) error {
	return nil
}

func (o OfferMysql) GetOfferName(id int) (string, error) {
	return "", nil
}
