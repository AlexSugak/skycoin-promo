package generator

import "github.com/jmoiron/sqlx"

// PromoGenerator represents a service that generates promocodes
type PromoGenerator interface {
	Generate(promoID string, count int) error
}

// Generator represents a service that generates promocodes
type Generator struct {
	DB *sqlx.DB
}

// NewGenerator creates a new instance of the Generator
func NewGenerator(DB *sqlx.DB) *Generator {
	return &Generator{DB: DB}
}

// Generate generates promo codes for specified promo campaign
func (g Generator) Generate(promoID string, count int) error {
	return nil
}
