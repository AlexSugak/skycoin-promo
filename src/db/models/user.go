package models

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

// RegisteredUser represents a gathered data about a user
type RegisteredUser struct {
	ID                    int64           `db:"Id"`
	CreatedAt             time.Time       `db:"CreatedAt"`
	UpdatedAt             time.Time       `db:"UpdatedAt"`
	PromoID               int64           `db:"PromoId"`
	PromoCodeID           int64           `db:"PromoCodeId"`
	FirstName             string          `db:"FirstName"`
	LastName              string          `db:"LastName"`
	Email                 string          `db:"Email"`
	Mobile                string          `db:"Mobile"`
	AddressLine1          string          `db:"AddressLine1"`
	AddressLine2          string          `db:"AddressLine2"`
	City                  string          `db:"City"`
	State                 string          `db:"State"`
	Postcode              string          `db:"Postcode"`
	IP                    string          `db:"IP"`
	UserAgent             string          `db:"UserAgent"`
	CountryCode           string          `db:"CountryCode"`
	PublicKey             string          `db:"PublicKey"`
	Amount                decimal.Decimal `db:"Amount"`
	Status                string          `db:"Status"`
	RejectionCode         int             `db:"RejectionCode"`
	TransferError         sql.NullString  `db:"TransferError"`
	TransferTransactionID sql.NullString  `db:"TransferTransactionID"`
}
