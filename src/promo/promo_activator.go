package promo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/AlexSugak/skycoin-promo/src/db/models"
	e "github.com/AlexSugak/skycoin-promo/src/errors"
	"github.com/AlexSugak/skycoin-promo/src/util"
	"github.com/AlexSugak/skycoin-promo/src/util/httputil"
)

// ActivationRequest represents a data for activation a promo code
type ActivationRequest struct {
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Mobile       string `json:"mobile" validate:"required"`
	AddressLine1 string `json:"addressLine1" validate:"required"`
	AddressLine2 string `json:"addressLine2" validate:"required"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	Postcode     string `json:"postcode" validate:"required"`
	CountryCode  string `json:"countryCode" validate:"required"`
	Recaptcha    string `json:"recaptcha" validate:"required"`
}

// ActivationResponse represents a data for activation a promo code
type ActivationResponse struct {
	Seed string `json:"seed"`
}

// ActivationHandler activates a promo
// Method: POST
// Content-type: application/json
// URI: /promo/activate?pid={promoId}&code={promoCode}
func ActivationHandler(s *HTTPServer) httputil.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vars := r.URL.Query()

		res, err := strconv.ParseInt(vars.Get("pid"), 10, 64)
		if err != nil {
			return e.CreateSingleValidationError("pid", "is not valid. pid should be a number")
		}
		pID := fmt.Sprintf("%d", res)

		res, err = strconv.ParseInt(vars.Get("code"), 10, 64)
		if err != nil {
			return e.CreateSingleValidationError("pCode", "is not valid. pCode should be a number")
		}
		pCode := fmt.Sprintf("%d", res)

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

		cp, err := s.checkRecaptcha(activationRequest.Recaptcha)
		if err != nil {
			return err
		} else if !cp {
			return e.CreateSingleValidationError("recaptcha", "is not valid")
		}

		promo, err := s.activator.GetPromo(pID)
		if err != nil {
			return err
		} else if promo == nil {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo campaign doesn't exist", pID),
				Code: http.StatusNotFound,
			}
		}

		promoCode, err := s.activator.GetPromoCodeByCode(pCode)
		if err != nil {
			return err
		}

		if promoCode == nil {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code doesn't exist", pCode),
				Code: http.StatusNotFound,
			}
		} else if promo.ID != promoCode.PromoID {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code doesn't exist", pCode),
				Code: http.StatusBadRequest,
			}
		} else if promoCode.Activated {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code has been already activated", promoCode.Code),
				Code: http.StatusBadRequest,
			}
		}

		u := models.RegisteredUser{
			PromoID:      promoCode.PromoID,
			PromoCodeID:  promoCode.ID,
			FirstName:    activationRequest.FirstName,
			LastName:     activationRequest.LastName,
			Email:        activationRequest.Email,
			Mobile:       activationRequest.Mobile,
			AddressLine1: activationRequest.AddressLine1,
			AddressLine2: activationRequest.AddressLine2,
			CountryCode:  activationRequest.CountryCode,
			City:         activationRequest.City,
			State:        activationRequest.State,
			Postcode:     activationRequest.Postcode,
			IP:           r.RemoteAddr,
			UserAgent:    util.TrimLong(r.Header.Get("User-Agent"), 1000),
			Amount:       promoCode.Amount,
		}

		seed, err := s.skyNode.GetSeed()
		if err != nil {
			return err
		}

		csrf, err := s.skyNode.GetCsrfToken()
		if err != nil {
			return err
		}

		wll, err := s.skyNode.CreateWallet(fmt.Sprintf("%s_%s_promo_wallet_%s", u.FirstName, u.LastName, promoCode.Code), seed, csrf)
		if err != nil {
			return err
		}

		csrf, err = s.skyNode.GetCsrfToken()
		if err != nil {
			return err
		}

		factor, _ := decimal.NewFromString("1000000")
		coins := promo.AmountPerAccount.Mul(factor)
		err = s.skyNode.TransferMoney(promo.SourceKey, wll.Entries[0].Address, coins, csrf)
		if err != nil {
			return err
		}

		u.PublicKey = wll.Entries[0].PublicKey
		_, err = s.activator.RegisterUser(u)

		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(ActivationResponse{Seed: seed})
	}
}
