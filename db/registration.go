package db

import "github.com/AlexSugak/skycoin-promo/db/models"

// RegisterPromo accepts a full user model and inserts it into a database
func (s Storage) RegisterPromo(u models.RegisteredUser) error {
	cmd := `INSERT INTO Registration ` +
		`(PromoId, ` +
		`PromoCodeId, ` +
		`FirstName, ` +
		`LastName, ` +
		`Email, ` +
		`Mobile, ` +
		`AddressLine1, ` +
		`AddressLine2, ` +
		`City, ` +
		`State, ` +
		`Postcode, ` +
		`IP, ` +
		`UserAgent, ` +
		`CountryCode, ` +
		`PublicKey, ` +
		`Amount) ` +
		`VALUES ` +
		`(:PromoId, ` +
		`:PromoCodeId, ` +
		`:FirstName, ` +
		`:LastName, ` +
		`:Email, ` +
		`:Mobile, ` +
		`:AddressLine1, ` +
		`:AddressLine2, ` +
		`:City, ` +
		`:State, ` +
		`:Postcode, ` +
		`:IP, ` +
		`:UserAgent, ` +
		`:CountryCode, ` +
		`:PublicKey, ` +
		`:Amount) `

	_, err := s.DB.NamedExec(cmd, u)

	return err
}
