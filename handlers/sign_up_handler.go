package handlers

import (
	"encoding/json"
	"go_auth/requests"
	"go_auth/responses"
	"net/http"
)

// swagger:route POST /sign_up Auth idSignUp
// Signs up a new user.
// responses:
//   201: successResponse
//   422: invalidResponse

// swagger:parameters idSignUp
type signUpParamsWrapper struct {
	// Signs up a new user.
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
	err := signUpRequest.HandleValidateEmailAndPassword(w)
	if err != nil {
		return
	}
	_, err = signUpRequest.HandleCheckExists(w)
	if err != nil {
		return
	}
	user, err := signUpRequest.HandleCreateUser(w)
	if err != nil {
		return
	}
	signUpRequest.SignToken(w, user)
	json.NewEncoder(w).Encode(user.BuildUser())
}
