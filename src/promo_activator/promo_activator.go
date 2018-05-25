package activator

import "github.com/AlexSugak/skycoin-promo/src/db/models"

// PromoActivator is a type of a service that register users with theirs promo
type PromoActivator interface {
	GetPromo(int64) (*models.Promo, error)
	GetPromoCodeByCode(string) (*models.PromoCode, error)
	RegisterUser(models.RegisteredUser) (*models.RegisteredUser, error)
	UpdateRegistration(u models.RegisteredUser) error
	GetRegisteredCodesAmount(int64) (int, error)
}
