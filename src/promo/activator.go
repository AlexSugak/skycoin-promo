package promo

import (
	"encoding/json"
	"net/http"

	"github.com/AlexSugak/skycoin-promo/src/util/httputil"
)

// ActivationHandler activates a promo
// Method: POST
// Content-type: application/json
// URI: /promo/activate
func ActivationHandler(s *HTTPServer) httputil.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		res := ""
		return json.NewEncoder(w).Encode(res)
	}
}
