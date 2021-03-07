package repository

import (
	"fmt"
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/jmoiron/sqlx"
)

type MyError struct{}

type DomainsMysql struct {
	db *sqlx.DB
}

func (m *MyError) Error() string {
	return " id domain is missing"
}

func NewDomainsMysql(db *sqlx.DB) *DomainsMysql {
	return &DomainsMysql{db: db}
}

func (d *DomainsMysql) GetAllDomains() ([]modelsStruct.Domain, error) {
	var domains []modelsStruct.Domain

	query := fmt.Sprintf("SELECT * FROM %s", domainsTable)

	err := d.db.Select(&domains, query)

	return domains, err
}

func (d *DomainsMysql) InsertDomain(domain string) error {
	query := fmt.Sprintf("INSERT INTO %s (domain_name) VALUES (?);", domainsTable)
	row := d.db.QueryRow(query, domain)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

func (d DomainsMysql) CheckDomain(domain string) (int, error) {
	var req int
	query := fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE domain_name=?;", domainsTable)
	err := d.db.Get(&req, query, domain)
	return req, err
}

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
