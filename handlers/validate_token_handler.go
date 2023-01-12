package handlers

import (
	"encoding/json"
	"go_auth/helpers"
	"net/http"
)

// swagger:route GET /validate_token
// Validates token and returns user.
// responses:
//   201: successResponse
//   401: invalidResponse
//   400: invalidResponse

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.CurrentUser(w, r)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(user.BuildUser())
}
