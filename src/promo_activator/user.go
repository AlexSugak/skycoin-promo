package activator

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/shopspring/decimal"
)

// Email represents a string that contains email address
type Email string

// Value implements the driver.Valuer interface
func (r Email) Value() (driver.Value, error) {
	return string(r), nil
}

// Scan implements the sql.Scanner interface
func (r *Email) Scan(src interface{}) error {
	if src == nil {
		*r = ""
	} else {
		*r = Email(string(src.([]uint8)))
	}
	return nil
}

// Mobile represents a string that contains phone number
type Mobile string

// Value implements the driver.Valuer interface
func (r Mobile) Value() (driver.Value, error) {
	return string(r), nil
}

// Scan implements the sql.Scanner interface
func (r *Mobile) Scan(src interface{}) error {
	if src == nil {
		*r = ""
	} else {
		*r = Mobile(string(src.([]uint8)))
	}
	return nil
}

// RegisteredUser represents a gathered data about a user
type RegisteredUser struct {
	ID                    int64              `db:"Id"`
	CreatedAt             time.Time          `db:"CreatedAt"`
	UpdatedAt             time.Time          `db:"UpdatedAt"`
	PromoID               PromoID            `db:"PromoId"`
	PromoCodeID           PromoCodeID        `db:"PromoCodeId"`
	FirstName             string             `db:"FirstName"`
	LastName              string             `db:"LastName"`
	Email                 Email              `db:"Email"`
	Mobile                Mobile             `db:"Mobile"`
	AddressLine1          string             `db:"AddressLine1"`
	AddressLine2          string             `db:"AddressLine2"`
	City                  string             `db:"City"`
	State                 string             `db:"State"`
	Postcode              string             `db:"Postcode"`
	IP                    string             `db:"IP"`
	UserAgent             string             `db:"UserAgent"`
	CountryCode           string             `db:"CountryCode"`
	PublicKey             string             `db:"PublicKey"`
	Amount                decimal.Decimal    `db:"Amount"`
	Status                RegistrationStatus `db:"Status"`
	RejectionCode         RejectionCode      `db:"RejectionCode"`
	TransferError         sql.NullString     `db:"TransferError"`
	TransferTransactionID sql.NullString     `db:"TransferTransactionID"`
}

// RejectionCode is an entity that corresponds to a rejection reason of registration a promo code
type RejectionCode int

const (
	// None - registration success
	None RejectionCode = 0
	// MaxAccountsReached - there is no left promo codes
	MaxAccountsReached RejectionCode = 100
	// InvalidRedeemCode - something wrong with promo code
	InvalidRedeemCode RejectionCode = 101
	// Duplicate - such promo code has been already registered
	Duplicate RejectionCode = 102
	// Aborted - request is not validated
	Aborted RejectionCode = 103
)

// Value implements the driver.Valuer interface
func (r RejectionCode) Value() (driver.Value, error) {
	return int64(r), nil
}

// Scan implements the sql.Scanner interface
func (r *RejectionCode) Scan(src interface{}) error {
	if src == nil {
		*r = 0
	} else {
		*r = RejectionCode(src.(int64))
	}
	return nil
}

// RegistrationStatus is part of user data and it describes a status of promo code registration
type RegistrationStatus string

// Value implements the driver.Valuer interface
func (r RegistrationStatus) Value() (driver.Value, error) {
	return string(r), nil
}

// Scan implements the sql.Scanner interface
func (r *RegistrationStatus) Scan(src interface{}) error {
	if src == nil {
		*r = ""
	} else {
		*r = RegistrationStatus(string(src.([]uint8)))
	}
	return nil
}

const (
	// Completed - registration successfully finished
	Completed RegistrationStatus = "completed"
	// Pending - system waits for request validation
	Pending RegistrationStatus = "pending"
	// Rejected - registration is rejected due some reasons
	Rejected RegistrationStatus = "rejected"
)
