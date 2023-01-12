package handlers

import (
	"encoding/json"
	"go_auth/requests"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
	errors := loginRequest.HandleValidateEmailAndPassword()
	body, _ := json.Marshal(errors)
	if len(errors) > 0 {
		http.Error(w, string(body), http.StatusBadRequest)
		return
	}

	user := loginRequest.HandleFindUserByEmail(w)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		_errors := map[string][]string{
			"errors": {"Invalid Email or Password"},
		}
		body, _ := json.Marshal(_errors)
		http.Error(w, string(body), http.StatusUnauthorized)
		return
	}
	loginRequest.SignToken(w, user)
	json.NewEncoder(w).Encode(user.BuildUser())
}
