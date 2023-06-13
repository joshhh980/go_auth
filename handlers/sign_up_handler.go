package handlers

import (
	"encoding/json"
	"go_auth/consts"
	"go_auth/models"
	"go_auth/requests"
	"go_auth/responses"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// swagger:route POST /sign_up idSignUp
// Sign up user.
// responses:
//   201: successResponse
//   422: invalidResponse

// swagger:parameters idSignUp
type signUpParamsWrapper struct {
	// Signup request params.
	// in:body
	Body requests.SignUpRequest
}

// Successful response.
// swagger:response successResponse
type successResponse struct {
	// Authorization header
	Authorization string

	// in:body
	Body responses.UserResponse
}

// Invalid response.
// swagger:response invalidResponse
type InvalidResponse struct {
	// in:body
	Body struct {
		Err string
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var signUpRequest requests.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&signUpRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errors := signUpRequest.HandleValidateEmailAndPassword()
	body, _ := json.Marshal(errors)
	if len(errors) > 0 {
		http.Error(w, string(body), http.StatusBadRequest)
		return
	}

	_user := signUpRequest.HandleFindUserByEmail(w)

	if _user.Email != "" {
		_errors := map[string][]string{
			"errors": {"Email already exists"},
		}
		body, _ := json.Marshal(_errors)
		http.Error(w, string(body), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := models.User{
		Name:     signUpRequest.Name,
		Email:    signUpRequest.Email,
		Password: string(hashedPassword),
	}

	consts.DB.Create(&user)
	signUpRequest.SignToken(w, user)
	json.NewEncoder(w).Encode(user.BuildUser())
}
