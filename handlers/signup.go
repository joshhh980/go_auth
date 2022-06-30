package handlers

import (
	"encoding/json"
	"go_auth/consts"
	"go_auth/models"
	"net/http"
	"net/mail"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Id uint
	jwt.StandardClaims
}

type SignUpRequest struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID    uint
	Email string
}

func buildUser(u *models.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}

func handleError(w http.ResponseWriter, s string) {
	http.Error(w, s, http.StatusBadRequest)
}

// signUpHandler handles incoming sign up requests
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var signUpRequest SignUpRequest
	json.NewDecoder(r.Body).Decode(&signUpRequest)
	if signUpRequest.Email == "" {
		handleError(w, "Email can't be blank")
		return
	}
	_, err := mail.ParseAddress(signUpRequest.Email)
	if err != nil {
		handleError(w, "Email must be valid")
		return
	}
	if signUpRequest.Password == "" {
		handleError(w, "Password can't be blank")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := &models.User{
		Email:    signUpRequest.Email,
		Password: string(hashedPassword),
	}
	consts.InitializeDB()
	consts.DB.Create(&user)
	claims := &Claims{
		Id: user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(consts.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Authorization", "Bearer "+tokenString)
	json.NewEncoder(w).Encode(buildUser(user))
}
