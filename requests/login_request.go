package requests

import (
	"errors"
	"go_auth/consts"
	"go_auth/helpers"
	"go_auth/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SessionsRequest struct {
	// required: true
	// example: example@mail.com
	Email string `json:"email"`
	// required: true
	Password string `json:"password"`
}

type LoginRequest struct {
	SessionsRequest
}

func (l LoginRequest) HandleLogin(w http.ResponseWriter) (models.User, error) {
	email := l.Email
	user := models.User{}
	// check if email exists
	consts.DB.Find(&user, models.User{
		Email: email,
	})
	if user.Email == "" {
		helpers.HandleError(w, "Invalid Email or Password")
		return user, errors.New("")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		helpers.HandleError(w, "Invalid Email or Password")
		return user, errors.New("")
	}
	return user, nil
}
