package handlers

import (
	"encoding/json"
	"go_auth/requests"
	"net/http"
)

// swagger:route POST /login Auth idLogin
// Logs in a user.
// responses:
//   201: successResponse
//   422: invalidResponse

// swagger:parameters idLogin
type loginParamsWrapper struct {
	// Login a user.
	// in:body
	Body requests.LoginRequest
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest requests.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := loginRequest.HandleValidateEmailAndPassword(w)
	if err != nil {
		return
	}
	user, err := loginRequest.HandleLogin(w)
	if err != nil {
		return
	}
	loginRequest.SignToken(w, user)
	json.NewEncoder(w).Encode(user.BuildUser())
}
