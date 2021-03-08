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

// CheckPrice ...
func (p *PricesMysql) CheckPrice(priceID int) (int, error) {
	var req int
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE id=?;", pricesTable)
	err := p.db.Get(&req, query, priceID)
	return req, err
}

// UpdatePrice ...
func (p *PricesMysql) UpdatePrice(pr modelsstruct.Price) (int, error) {

	check, err := p.CheckPrice(pr.ID)
	if err != nil {
		return 0, err
	}

	if check == 0 {
		return 0, err
	}

	query := fmt.Sprintf("UPDATE %s SET name=?, geo=?, country=?,"+
		"country_1=?, country_2=?, old=?, new=?, money=?, tel=? WHERE id=?;", pricesTable)
	_, err = p.db.Exec(query, pr.Name, pr.GEO, pr.Country, pr.Country1, pr.Country2, pr.Old, pr.New, pr.Money, pr.Tel, pr.ID)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return 1, nil
}
