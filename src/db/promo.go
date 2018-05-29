package db

import "github.com/AlexSugak/skycoin-promo/src/db/models"

// GetPromo returns an entity of a Promo from the DB
func (s Storage) GetPromo(ID int64) (*models.Promo, error) {
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

	promo := models.Promo{}
	err := s.DB.Get(&promo, cmd, ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return &promo, err
	}

	return &promo, nil
}

// GetPromoCodeByCode returns an entity of a PromoCode from the DB
func (s Storage) GetPromoCodeByCode(code string) (*models.PromoCode, error) {
	cmd := `SELECT ` +
		`pc.Id, ` +
		`pc.CreatedAt, ` +
		`pc.PromoId, ` +
		`pc.Code, ` +
		`p.AmountPerAccount as Amount, ` +
		`COALESCE(pc.Id = r.PromoCodeId, false) as Activated ` +
		`FROM Promo p ` +
		`INNER JOIN PromoCode pc on p.Id = pc.PromoId ` +
		`LEFT JOIN Registration r on r.PromoCodeId = pc.Id ` +
		`WHERE pc.Code = ?`

	promoCode := models.PromoCode{}
	err := s.DB.Get(&promoCode, cmd, code)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return &promoCode, err
	}

	return &promoCode, nil
}

// GetRegisteredCodesAmount calculates amount registered promo codes for such promoID
func (s Storage) GetRegisteredCodesAmount(promoID int) (int, error) {
	cmd := `SELECT DISTINCT COUNT(r.ID) ` +
		`FROM Registration r ` +
		`WHERE r.PromoId = ?`

	registeredCodesAmount := 0
	err := s.DB.Get(&registeredCodesAmount, cmd, promoID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, nil
		}
		return 0, err
	}

	return registeredCodesAmount, nil
}

// InsertPromoCode inserts a new promocode to the database
func (s Storage) InsertPromoCode(code *models.PromoCode) (*models.PromoCode, error) {
	cmd := "INSERT INTO PromoCode (PromoId, Code) VALUES(:PromoId, :Code)"

	res, err := s.DB.NamedExec(cmd, code)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	code.ID = id
	return code, nil
}
