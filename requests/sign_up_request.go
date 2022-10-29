package requests

import (
	"errors"
	"go_auth/consts"
	"go_auth/helpers"
	"go_auth/models"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	//	required: true
	//	example: User
	Name string `json:"name"`
	SessionsRequest
	//	required: true
	C_Password string `json:"c_password"`
}

func (s *SessionsRequest) HandleValidateEmailAndPassword(w http.ResponseWriter) error {
	email := s.Email
	password := s.Password
	if email == "" {
		helpers.HandleError(w, "Email can't be blank")
		return errors.New("")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		helpers.HandleError(w, "Email must be valid")
		return errors.New("")
	}
	if password == "" {
		helpers.HandleError(w, "Password can't be blank")
		return errors.New("")
	}
	return nil
}

func (s *SessionsRequest) SignToken(w http.ResponseWriter, user models.User) {
	tokenString, err := helpers.SignToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Authorization", "Bearer "+tokenString)
}

func (s SignUpRequest) HandleCreateUser(w http.ResponseWriter) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return models.User{}, errors.New("")
	}
	user := models.User{
		Email:    s.Email,
		Password: string(hashedPassword),
	}
	consts.DB.Create(&user)
	return user, nil
}

func (s SignUpRequest) HandleCheckExists(w http.ResponseWriter) (models.User, error) {
	email := s.Email
	_user := models.User{}
	consts.DB.Find(&_user, models.User{
		Email: email,
	})
	if _user.Email != "" {
		helpers.HandleError(w, "Email already exists")
		return _user, errors.New("")
	}
	return _user, nil
}
