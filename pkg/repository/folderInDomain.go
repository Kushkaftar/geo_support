package repository

import "github.com/jmoiron/sqlx"

// FoldersInDomainMysql ...
type FoldersInDomainMysql struct {
	db *sqlx.DB
}

// NewFoldersInDomainMysql ...
func NewFoldersInDomainMysql(db *sqlx.DB) *FoldersInDomainMysql {
	return &FoldersInDomainMysql{db: db}
}

// TODO: доделать
// InsertFolder ...
func (f *FoldersInDomainMysql) InsertFolder() error {
	return nil
}
