package promo

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/AlexSugak/skycoin-promo/db/models"
	e "github.com/AlexSugak/skycoin-promo/src/errors"
	"github.com/AlexSugak/skycoin-promo/src/util"
	"github.com/AlexSugak/skycoin-promo/src/util/httputil"
)

// ActivationRequest represents a data for activation a promo code
type ActivationRequest struct {
	FirstName     string `json:"firstName" validate:"required"`
	LastName      string `json:"lastName" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Mobile        string `json:"mobile" validate:"required"`
	AddressLine1  string `json:"addressLine1" validate:"required"`
	AddressLine2  string `json:"addressLine2" validate:"required"`
	City          string `json:"city" validate:"required"`
	State         string `json:"state" validate:"required"`
	Postcode      string `json:"postcode" validate:"required"`
	CountryCode   string `json:"countryCode" validate:"required"`
	Recaptcha     string `json:"recaptcha" validate:"required"`
	PromotionCode string `json:"promotionCode" validate:"required"`
}

// ActivationResponse represents a data for activation a promo code
type ActivationResponse struct {
	Seed string `json:"seed"`
}

// ActivationHandler activates a promo
// Method: POST
// Content-type: application/json
// URI: /promo/activate
func ActivationHandler(s *HTTPServer) httputil.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		activationRequest := &ActivationRequest{}

		if err := json.NewDecoder(r.Body).Decode(activationRequest); err != nil {
			return httputil.StatusError{
				Err:  fmt.Errorf("Invalid json body of the request: %v", err),
				Code: http.StatusBadRequest,
			}
		}

		if err := s.validate.Struct(activationRequest); err != nil {
			return e.ValidatorErrorsResponse(err.(validator.ValidationErrors))
		}

		// TODO: Uncomment when recaptcha will be ready
		// res, err := s.checkRecaptcha(activationRequest.Recaptcha)
		// if err != nil {
		// 	return err
		// } else if !res {
		// 	return e.CreateSingleValidationError("recaptcha", "is not valid")
		// }

		promoCode, err := s.activator.GetPromoCodeByCode(activationRequest.PromotionCode)
		if err != nil {
			return err
		}

		if promoCode == nil {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code doesn't exist", activationRequest.PromotionCode),
				Code: http.StatusNotFound,
			}
		} else if !promoCode.Activated {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code has been already activated", promoCode.Code),
				Code: http.StatusBadRequest,
			}
		}

		u := models.RegisteredUser{
			Code:         promoCode.Code,
			PromoCodeID:  promoCode.ID,
			FirstName:    activationRequest.FirstName,
			LastName:     activationRequest.LastName,
			Email:        activationRequest.Email,
			Mobile:       activationRequest.Mobile,
			AddressLine1: activationRequest.AddressLine1,
			AddressLine2: activationRequest.AddressLine2,
			City:         activationRequest.City,
			State:        activationRequest.State,
			Postcode:     activationRequest.Postcode,
			IP:           r.RemoteAddr,
			CountryCode:  activationRequest.CountryCode,
			UserAgent:    util.TrimLong(r.Header.Get("User-Agent"), 1000),
			Amount:       promoCode.Amount,
			PublicKey:    "PublicKey",
		}

		err = s.activator.RegisterPromo(u)
		if err != nil {
			return err
		}

		seed, err := s.skyNode.GetSeed()
		if err != nil {
			return err
		}

		csrf, err := s.skyNode.GetCsrfToken()
		if err != nil {
			return err
		}

		_, err = s.skyNode.CreateWallet(fmt.Sprintf("%s_%s_wallet_%s", u.FirstName, u.LastName, u.Code), seed, csrf)
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(ActivationResponse{Seed: seed})
	}
}
