package models

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

// PromoCode represents a promo code entry in the db
// Field 'Activated' is calculated by a query
type PromoCode struct {
	ID        int64           `db:"Id"`
	CreatedAt time.Time       `db:"CreatedAt"`
	PromoID   int64           `db:"PromoId"`
	Code      string          `db:"Code"`
	Activated bool            `db:"Activated"`
	Amount    decimal.Decimal `db:"Amount"`
}

// Promo represents a promo entry in the db
type Promo struct {
	ID               int64            `db:"Id"`
	CreatedAt        time.Time        `db:"CreatedAt"`
	UpdatedAt        time.Time        `db:"UpdatedAt"`
	Name             string           `db:"Name"`
	Description      string           `db:"Description"`
	Tandc            sql.NullString   `db:"Tandc"`
	StartAt          *time.Time       `db:"StartAt"`
	EndAt            *time.Time       `db:"EndAt"`
	AmountPerAccount decimal.Decimal  `db:"AmountPerAccount"`
	MaxAccounts      int              `db:"MaxAccounts"`
	EnabledYN        bool             `db:"EnabledYN"`
	ShowKeyYN        bool             `db:"ShowKeyYN"`
	EmailKeyYN       bool             `db:"EmailKeyYN"`
	AdminEmail       string           `db:"AdminEmail"`
	SourceKey        string           `db:"SourceKey"`
	CleanupKey       sql.NullString   `db:"CleanupKey"`
	CleanedUpAmount  *decimal.Decimal `db:"CleanedUpAmount"`
	CleanedUpAt      *time.Time       `db:"CleanedUpAt"`
}
