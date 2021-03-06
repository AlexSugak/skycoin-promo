package generator

import (
	"strings"

	activator "github.com/AlexSugak/skycoin-promo/src/promo_activator"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// PromoGenerator represents a service that generates promocodes
type PromoGenerator interface {
	Generate([]activator.Promo) error
	GetEmptyPromos() ([]activator.Promo, error)
}

// Generator represents a service that generates promocodes
type Generator struct {
	DB *sqlx.DB
}

// NewGenerator creates a new instance of the Generator
func NewGenerator(DB *sqlx.DB) *Generator {
	return &Generator{DB: DB}
}

// Generate generates promo codes for specified promo campaign
func (g Generator) Generate(promos []activator.Promo) error {
	for i := 0; i < len(promos); i++ {
		promoCodes := make([]activator.PromoCode, promos[i].MaxAccounts)
		for j := 0; j < promos[i].MaxAccounts; j++ {
			u := uuid.NewV4()

			pc := activator.PromoCode{
				PromoID: promos[i].ID,
				Code:    activator.Code(strings.Replace(u.String(), "-", "", -1)),
			}
			_, err := g.insertPromoCode(&pc)
			if err != nil {
				return err
			}
			promoCodes[i] = pc
		}
	}
	return nil
}

func (g Generator) insertPromoCode(code *activator.PromoCode) (*activator.PromoCode, error) {
	cmd := "INSERT INTO PromoCode (PromoId, Code) VALUES(:PromoId, :Code)"

	res, err := g.DB.NamedExec(cmd, code)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	code.ID = activator.PromoCodeID(id)
	return code, nil
}

// GetEmptyPromos returns a list of promo campaigns that don't have any codes
func (g Generator) GetEmptyPromos() ([]activator.Promo, error) {
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
		`LEFT JOIN PromoCode pc ON pc.PromoId = p.Id ` +
		`WHERE pc.Id IS NULL`

	promos := []activator.Promo{}
	err := g.DB.Select(&promos, cmd)

	return promos, err
}
