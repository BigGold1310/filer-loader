package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB - Holds the sql.DB object pointer
type DB struct {
	*sqlx.DB
}

// NewDB - Creats a new sqlx.DB
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
