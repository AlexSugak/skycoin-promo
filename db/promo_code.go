package db

import "github.com/AlexSugak/skycoin-promo/db/models"

// GetPromoCodeByCode returns an entity of a PromoCode
func (s Storage) GetPromoCodeByCode(code string) (*models.PromoCode, error) {
	cmd := `SELECT ` +
		`pc.Id, ` +
		`pc.CreatedAt, ` +
		`pc.PromoId, ` +
		`pc.Code, ` +
		`p.AmountPerAccount as Amount, ` +
		`COALESCE(pc.Id = PromoCodeId, false) as Activated ` +
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
