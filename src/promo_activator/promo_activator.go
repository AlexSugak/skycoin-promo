package activator

import "github.com/AlexSugak/skycoin-promo/src/db/models"

// PromoActivator is a type of a service that register users with theirs promo
type PromoActivator interface {
	GetPromo(models.PromoID) (*models.Promo, error)
	GetPromoCodeByCode(models.Code) (*models.PromoCode, error)
	RegisterUser(models.RegisteredUser) (*models.RegisteredUser, error)
	UpdateRegistration(u models.RegisteredUser) error
	GetRegisteredCodesAmount(models.PromoID) (int, error)
	GetRegistrationsByEmailOrPhone(string, string) (*models.RegisteredUser, error)
}
