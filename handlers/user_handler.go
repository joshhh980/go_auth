package handlers

import (
	"encoding/json"
	"go_auth/consts"
	"go_auth/helpers"
	"go_auth/models"
	"go_auth/requests"
	"net/http"
)

// swagger:route GET /user
// Get current user.
// responses:
//   201: successResponse
//   401: invalidResponse
//   400: invalidResponse

func ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.CurrentUser(w, r)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(user.BuildUser())
}

// swagger:route PUT /user idUpdateUser
// Update user details.
// responses:
//   200: successResponse
//   422: invalidResponse

// swagger:parameters idUpdateUser
type userParamsWrapper struct {
	// update user.
	// in:body
	Body requests.UserRequest
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.CurrentUser(w, r)
	if err != nil {
		return
	}
	var signUpRequest requests.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&signUpRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	consts.DB.Model(&user).Updates(models.User{
		Email: signUpRequest.Email,
		Name:  signUpRequest.Name,
	})
	json.NewEncoder(w).Encode(user.BuildUser())
}