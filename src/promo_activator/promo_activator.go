package activator

import (
	"github.com/AlexSugak/skycoin-promo/src/db"
	"github.com/jmoiron/sqlx"
)

// PromoActivator is a type of a service that register users with theirs promo
type PromoActivator interface {
	GetPromo(PromoID) (*Promo, error)
	GetPromoCodeByCode(Code) (*PromoCode, error)
	RegisterUser(RegisteredUser) (*RegisteredUser, error)
	UpdateRegistration(RegisteredUser) error
	GetRegisteredCodesAmount(PromoID) (int, error)
	GetRegistrationByEmailOrPhone(Email, Mobile) (*RegisteredUser, error)
}

// Activator represents a service that generates promocodes
type Activator struct {
	DB *sqlx.DB
}

// NewActivator creates new Activator
func NewActivator(db *sqlx.DB) Activator {
	return Activator{db}
}

// RegisterUser accepts a full user model and inserts it into a database
func (s Activator) RegisterUser(u RegisteredUser) (*RegisteredUser, error) {
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
func (s Activator) UpdateRegistration(u RegisteredUser) error {
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

// GetRegistrationByEmailOrPhone returns registration with such email and mobile
func (s Activator) GetRegistrationByEmailOrPhone(email Email, mobile Mobile) (*RegisteredUser, error) {
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
		`WHERE (r.Email = ? OR r.Mobile = ?) AND r.Status = 'completed'`

	registration := RegisteredUser{}
	err := s.DB.Get(&registration, cmd, email, mobile)
	if err != nil {
		if err.Error() == db.EmptyResult {
			return nil, nil
		}
		return nil, err
	}

	return &registration, nil
}

// GetPromo returns an entity of a Promo from the DB
func (s Activator) GetPromo(ID PromoID) (*Promo, error) {
	cmd := `SELECT ` +
		`p.Id, ` +
		`p.CreatedAt, ` +
		`p.UpdatedAt, ` +
		`p.Name, ` +
		`p.Description, ` +
		`p.Tandc, ` +
		`p.StartAt, ` +
		`p.EndAt, ` +
		`p.AmountPerAccount, ` +
		`p.MaxAccounts, ` +
		`p.EnabledYN, ` +
		`p.ShowKeyYN, ` +
		`p.EmailKeyYN, ` +
		`p.AdminEmail, ` +
		`p.SourceKey, ` +
		`p.CleanupKey, ` +
		`p.CleanedUpAmount, ` +
		`p.CleanedUpAt ` +
		`FROM Promo p ` +
		`WHERE p.ID = ?`

	promo := Promo{}
	err := s.DB.Get(&promo, cmd, ID)
	if err != nil {
		if err.Error() == db.EmptyResult {
			return nil, nil
		}
		return &promo, err
	}

	return &promo, nil
}

// GetPromoCodeByCode returns an entity of a PromoCode from the DB
func (s Activator) GetPromoCodeByCode(code Code) (*PromoCode, error) {
	cmd := `SELECT ` +
		`pc.Id, ` +
		`pc.CreatedAt, ` +
		`pc.PromoId, ` +
		`pc.Code, ` +
		`p.AmountPerAccount as Amount ` +
		`FROM Promo p ` +
		`INNER JOIN PromoCode pc on p.Id = pc.PromoId ` +
		`WHERE pc.Code = ?`

	promoCode := PromoCode{}
	err := s.DB.Get(&promoCode, cmd, code)
	if err != nil {
		if err.Error() == db.EmptyResult {
			return nil, nil
		}
		return &promoCode, err
	}

	registered := []int{}
	cmd = `SELECT Id FROM Registration WHERE Status = 'completed' AND PromoCodeId = ?`
	err = s.DB.Select(&registered, cmd, promoCode.ID)
	if err != nil {
		return nil, err
	}
	promoCode.Activated = len(registered) > 0
	return &promoCode, nil
}

// GetRegisteredCodesAmount calculates amount registered promo codes for such promoID
func (s Activator) GetRegisteredCodesAmount(promoID PromoID) (int, error) {
	cmd := `SELECT DISTINCT COUNT(r.ID) ` +
		`FROM Registration r ` +
		`WHERE r.PromoId = ? AND r.Status = 'completed'`

	registeredCodesAmount := 0
	err := s.DB.Get(&registeredCodesAmount, cmd, promoID)
	if err != nil {
		if err.Error() == db.EmptyResult {
			return 0, nil
		}
		return 0, err
	}

	return registeredCodesAmount, nil
}
