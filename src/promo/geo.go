package promo

import (
	"encoding/json"
	"net/http"

	"github.com/AlexSugak/skycoin-promo/src/util/httputil"
)

// CountriesHandler returns available countries
// Method: GET
// Content-type: application/json
// URI: /api/countries
func CountriesHandler(s *HTTPServer) httputil.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		countries, err := s.geo.Countries()
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(countries)
	}
}
