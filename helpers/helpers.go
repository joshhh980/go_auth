package helpers

import (
	"go_auth/models"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func SignToken(user models.User) (string, error) {
	claims := &Claims{
		Id: user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func HandleError(w http.ResponseWriter, s string) {
	http.Error(w, s, http.StatusBadRequest)
}
