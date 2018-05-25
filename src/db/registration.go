package db

import "github.com/AlexSugak/skycoin-promo/src/db/models"

// RegisterUser accepts a full user model and inserts it into a database
func (s Storage) RegisterUser(u models.RegisteredUser) (*models.RegisteredUser, error) {
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

	res, err := s.DB.NamedExec(cmd, u)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return &u, nil
}

// UpdateRegistration updates a registration model
func (s Storage) UpdateRegistration(u models.RegisteredUser) error {
	cmd := `UPDATE Registration SET ` +
		`PromoId=:PromoId, ` +
		`PromoCodeId=:PromoCodeId, ` +
		`FirstName=:FirstName, ` +
		`LastName=:LastName, ` +
		`Email=:Email, ` +
		`Mobile=:Mobile, ` +
		`AddressLine1=:AddressLine1, ` +
		`AddressLine2=:AddressLine2, ` +
		`City=:City, ` +
		`State=:State, ` +
		`Postcode=:Postcode, ` +
		`IP=:IP, ` +
		`UserAgent=:UserAgent, ` +
		`CountryCode=:CountryCode, ` +
		`PublicKey=:PublicKey, ` +
		`Amount=:Amount ` +
		`WHERE Id = :Id `

	_, err := s.DB.NamedExec(cmd, u)
	return err
}
