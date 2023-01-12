package requests

import (
	"go_auth/consts"
	"go_auth/helpers"
	"go_auth/models"
	"net/http"
	"net/mail"
)

type SessionsRequest struct {
	// required: true
	// example: example@mail.com
	Email string `json:"email"`
	// required: true
	Password string `json:"password"`
}

func (s SessionsRequest) HandleFindUserByEmail(w http.ResponseWriter) models.User {
	email := s.Email
	_user := models.User{}
	consts.DB.Find(&_user, models.User{
		Email: email,
	})

	return _user
}

func (s *SessionsRequest) HandleValidateEmailAndPassword() map[string][]interface{} {
	errors := map[string][]interface{}{}
	email := s.Email
	password := s.Password
	if email == "" {
		errors["email"] = append(errors["email"], "Email can't be blank")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		errors["email"] = append(errors["email"], "Email must be valid")
	}
	if password == "" {
		errors["password"] = append(errors["password"], "Password can't be blank")
	}
	return errors
}

func (s *SessionsRequest) SignToken(w http.ResponseWriter, user models.User) {
	tokenString, err := helpers.SignToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Authorization", "Bearer "+tokenString)
}
