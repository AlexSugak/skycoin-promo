package promo

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	e "github.com/AlexSugak/skycoin-promo/src/errors"
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
	Region        string `json:"region" validate:"required"`
	Postcode      string `json:"postcode" validate:"required"`
	Country       string `json:"country" validate:"required"`
	Recaptcha     string `json:"recaptcha" validate:"required"`
	PromotionCode string `json:"promotionCode" validate:"required"`
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
		res, err := s.checkRecaptcha(activationRequest.Recaptcha)
		if err != nil {
			return err
		} else if !res {
			return e.CreateSingleValidationError("recaptcha", "is not valid")
		}

		promoCode, err := s.activator.GetPromoCodeByCode(activationRequest.PromotionCode)
		if err != nil {
			return err
		}

		if promoCode.PromoID == 0 {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code doesn't exist", promoCode.Code),
				Code: http.StatusBadRequest,
			}
		} else if !promoCode.Activated {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code has been already activated", promoCode.Code),
				Code: http.StatusBadRequest,
			}
		}

		return json.NewEncoder(w).Encode(activationRequest)
	}
}
