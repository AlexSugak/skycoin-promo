package models

import (
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
