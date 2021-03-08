package repository

import (
	"fmt"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/jmoiron/sqlx"
)

// MyError ...
type MyError struct{}

// DomainsMysql ...
type DomainsMysql struct {
	db *sqlx.DB
}

// Error ...
func (m *MyError) Error() string {
	return " id domain is missing"
}

// NewDomainsMysql ...
func NewDomainsMysql(db *sqlx.DB) *DomainsMysql {
	return &DomainsMysql{db: db}
}

// GetAllDomains ...
func (d *DomainsMysql) GetAllDomains() ([]modelsstruct.Domain, error) {
	var domains []modelsstruct.Domain

	query := fmt.Sprintf("SELECT * FROM %s", domainsTable)

	err := d.db.Select(&domains, query)

	return domains, err
}

// InsertDomain ...
func (d *DomainsMysql) InsertDomain(domain string) error {
	query := fmt.Sprintf("INSERT INTO %s (domain_name) VALUES (?);", domainsTable)
	row := d.db.QueryRow(query, domain)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

// CheckDomain ...
func (d DomainsMysql) CheckDomain(domain string) (int, error) {
	var req int
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE domain_name=?;", domainsTable)
	err := d.db.Get(&req, query, domain)
	return req, err
}

// SetFlag ...
func (d *DomainsMysql) SetFlag(flag, id int) error {
	var check int

	queryChek := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE id=?;", domainsTable)
	err := d.db.Get(&check, queryChek, id)
	if err != nil {
		return err
	}

	if check == 0 {
		return &MyError{}
	}

	query := fmt.Sprintf("UPDATE %s SET checked=? WHERE id=?;", domainsTable)
	_, err = d.db.Exec(query, flag, id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// GetDomain ...
func (d *DomainsMysql) GetDomain(id int) (modelsstruct.Domain, error) {
	var domain modelsstruct.Domain

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?;", domainsTable)

	err := d.db.Get(&domain, query, id)

	return domain, err
}
