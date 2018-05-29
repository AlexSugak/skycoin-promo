package promo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

		pir, err := strconv.ParseInt(vars.Get("pid"), 10, 64)
		if err != nil {
			return e.CreateSingleValidationError("pid", "is not valid. pid should be a number")
		}
		pID := models.PromoID(pir)

		pCode := models.Code(vars.Get("code"))
		if pCode == "" {
			return e.CreateSingleValidationError("code", "is not required")
		}

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

		// cp, err := s.checkRecaptcha(activationRequest.Recaptcha)
		// if err != nil {
		// 	return err
		// } else if !cp {
		// 	return e.CreateSingleValidationError("recaptcha", "is not valid")
		// }

		u := &models.RegisteredUser{
			PromoID:       pID,
			FirstName:     activationRequest.FirstName,
			LastName:      activationRequest.LastName,
			Email:         models.Email(activationRequest.Email),
			Mobile:        models.Mobile(activationRequest.Mobile),
			AddressLine1:  activationRequest.AddressLine1,
			AddressLine2:  activationRequest.AddressLine2,
			CountryCode:   activationRequest.CountryCode,
			City:          activationRequest.City,
			State:         activationRequest.State,
			Postcode:      activationRequest.Postcode,
			IP:            r.RemoteAddr,
			UserAgent:     util.TrimLong(r.Header.Get("User-Agent"), 1000),
			Status:        models.Rejected,
			RejectionCode: models.Aborted,
		}

		var publicKey *string
		defer func() {
			if u.Status != models.Completed {
				u.Status = models.Rejected
			} else {
				u.PublicKey = *publicKey
			}
			s.activator.RegisterUser(*u)
		}()

		promo, err := s.activator.GetPromo(pID)
		if err != nil {
			return err
		} else if promo == nil {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%d' promo campaign doesn't exist", pID),
				Code: http.StatusNotFound,
			}
		} else if promo.StartAt != nil && time.Now().Before(*promo.StartAt) {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%d' promo campaign hasn't started yet", pID),
				Code: http.StatusBadRequest,
			}
		} else if promo.EndAt != nil && time.Now().Before(*promo.EndAt) {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%d' promo campaign has already finished", pID),
				Code: http.StatusBadRequest,
			}
		}

		promoCode, err := s.activator.GetPromoCodeByCode(pCode)
		if err != nil {
			return err
		} else if promoCode == nil {
			return httputil.StatusError{
				Err:  fmt.Errorf("'%s' promo code doesn't exist", pCode),
				Code: http.StatusNotFound,
			}
		}
		u.PromoCodeID = promoCode.ID
		if promo.ID != promoCode.PromoID {
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

		registeredCodesAmount, err := s.activator.GetRegisteredCodesAmount(promo.ID)
		if err != nil {
			return err
		} else if registeredCodesAmount == promo.MaxAccounts {
			u.RejectionCode = models.MaxAccountsReached
			return httputil.StatusError{
				Err:  fmt.Errorf("'%d' promo campaign has already reached max amount of registered codes", pID),
				Code: http.StatusBadRequest,
			}
		}

		eu, err := s.activator.GetRegistrationByEmailOrPhone(u.Email, u.Mobile)
		if err != nil {
			return err
		} else if eu != nil {
			u.RejectionCode = models.Duplicate
			return httputil.StatusError{
				Err:  fmt.Errorf("A user with such email or mobile has already registered"),
				Code: http.StatusBadRequest,
			}
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
		publicKey = &wll.Entries[0].PublicKey

		return json.NewEncoder(w).Encode(ActivationResponse{Seed: seed})
	}
}
