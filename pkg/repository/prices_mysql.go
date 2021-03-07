package repository

import (
	"fmt"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/jmoiron/sqlx"
)

// PricesMysql ...
type PricesMysql struct {
	db *sqlx.DB
}

// NewPricesMysql ...
func NewPricesMysql(db *sqlx.DB) *PricesMysql {
	return &PricesMysql{db: db}
}

// GetAllPrices ...
func (p *PricesMysql) GetAllPrices() ([]modelsstruct.Price, error) {
	var prices []modelsstruct.Price

	query := fmt.Sprintf("SELECT * FROM %s", pricesTable)

	err := p.db.Select(&prices, query)

	return prices, err
}
