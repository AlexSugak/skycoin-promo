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
