package activator

import "github.com/AlexSugak/skycoin-promo/db/models"

// PromoActivator a type of a service that register users with theirs promo
type PromoActivator interface {
	GetPromoCodeByCode(string) (models.PromoCode, error)
	RegisterPromo(models.RegisteredUser) error
}
