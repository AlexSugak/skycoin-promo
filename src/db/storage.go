package db

import (
	"github.com/jmoiron/sqlx"
)

// Storage stands for main DB storage
type Storage struct {
	DB *sqlx.DB
}

// NewStorage creates new Storage
func NewStorage(db *sqlx.DB) Storage {
	return Storage{db}
}

// EmptyResult is sql error string that occurs when Get function doesn't return any results
const EmptyResult = "sql: no rows in result set"
