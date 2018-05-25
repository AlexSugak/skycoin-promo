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
		`CountryCode, ` +
		`State, ` +
		`City, ` +
		`Postcode, ` +
		`IP, ` +
		`UserAgent, ` +
		`PublicKey, ` +
		`Amount, ` +
		`Status, ` +
		`RejectionCode) ` +
		`VALUES ` +
		`(:PromoId, ` +
		`:PromoCodeId, ` +
		`:FirstName, ` +
		`:LastName, ` +
		`:Email, ` +
		`:Mobile, ` +
		`:AddressLine1, ` +
		`:AddressLine2, ` +
		`:CountryCode, ` +
		`:State, ` +
		`:City, ` +
		`:Postcode, ` +
		`:IP, ` +
		`:UserAgent, ` +
		`:PublicKey, ` +
		`:Amount, ` +
		`:Status, ` +
		`:RejectionCode)`

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
		`CountryCode=:CountryCode, ` +
		`State=:State, ` +
		`City=:City, ` +
		`Postcode=:Postcode, ` +
		`IP=:IP, ` +
		`UserAgent=:UserAgent, ` +
		`PublicKey=:PublicKey, ` +
		`Amount=:Amount ` +
		`Status=:Status ` +
		`RejectionCode=:RejectionCode ` +
		`WHERE Id = :Id `

	_, err := s.DB.NamedExec(cmd, u)
	return err
}

// GetRegistrationsByEmailOrPhone returns registration with such email and mobile
func (s Storage) GetRegistrationsByEmailOrPhone(email string, mobile string) (*models.RegisteredUser, error) {
	cmd := `SELECT ` +
		`r.PromoId, ` +
		`r.PromoCodeId, ` +
		`r.FirstName, ` +
		`r.LastName, ` +
		`r.Email, ` +
		`r.Mobile, ` +
		`r.AddressLine1, ` +
		`r.AddressLine2, ` +
		`r.CountryCode, ` +
		`r.State, ` +
		`r.City, ` +
		`r.Postcode, ` +
		`r.IP, ` +
		`r.UserAgent, ` +
		`r.PublicKey, ` +
		`r.Amount, ` +
		`r.Status, ` +
		`r.RejectionCode ` +
		`FROM Registration r ` +
		`WHERE r.Email = ? OR r.Mobile = ?`

	registration := models.RegisteredUser{}
	err := s.DB.Get(&registration, cmd, email, mobile)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return &registration, nil
}
