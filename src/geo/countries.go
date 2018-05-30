package geo

import (
	"github.com/jmoiron/sqlx"
)

// Geo represents geolocation functions
type Geo interface {
	Countries() ([]Country, error)
}

// Storage represents a low-level entity that provides access to data base
type Storage struct {
	DB *sqlx.DB
}

// NewGeoStorage creates a new GeoStorage
func NewGeoStorage(db *sqlx.DB) *Storage {
	return &Storage{db}
}

// Countries returns all available countries
func (s Storage) Countries() ([]Country, error) {
	countries := []Country{}
	err := s.DB.Select(&countries, "SELECT * FROM Country")
	return countries, err
}
