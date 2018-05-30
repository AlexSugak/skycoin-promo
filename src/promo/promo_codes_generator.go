package promo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexSugak/skycoin-promo/src/util/httputil"
)

const promoAmount = 20

// GenerationHandler generates a bunch of promo codes
// Method: POST
// Content-type: application/json
// URI: /promo/codes/generate
func GenerationHandler(s *HTTPServer) httputil.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		promoCampaigns, err := s.generator.GetEmptyPromos()
		if err != nil {
			return err
		} else if len(promoCampaigns) == 0 {
			return httputil.StatusError{
				Err:  fmt.Errorf("there are no promo campaigns without promo codes. Add promo an empty campaign first"),
				Code: http.StatusBadGateway,
			}
		}

		err = s.generator.Generate(promoCampaigns)
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode("OK")
	}
}
