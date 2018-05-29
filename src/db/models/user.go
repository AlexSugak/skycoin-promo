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
	PromoID               PromoID         `db:"PromoId"`
	PromoCodeID           PromoCodeID     `db:"PromoCodeId"`
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

// RejectionCode is an entity that corresponds to a rejection reason of registration a promo code
type RejectionCode int

const (
	// MaxAccountsReached means that there is no left promo codes
	MaxAccountsReached RejectionCode = 100
	// InvalidRedeemCode means that something wrong with promo code
	InvalidRedeemCode RejectionCode = 101
	// Duplicate means that such promo code has been already registered
	Duplicate RejectionCode = 102
	// Aborted means that request is not validated
	Aborted RejectionCode = 103
)

// RegistrationStatus is part of user data and it describes a status of promo code registration
type RegistrationStatus string

const (
	// Completed - registration successfully finished
	Completed RegistrationStatus = "completed"
	// Pending - system waits for request validation
	Pending RegistrationStatus = "pending"
	// Rejected - registration is rejected due some reasons
	Rejected RegistrationStatus = "rejected"
)
