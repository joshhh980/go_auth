package helpers

import (
	"errors"
	"go_auth/consts"
	"go_auth/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

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

func CurrentUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	authorization := r.Header.Get("Authorization")
	authorization = strings.Split(authorization, " ")[1]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return models.User{}, errors.New("")
		}
		w.WriteHeader(http.StatusBadRequest)
		return models.User{}, errors.New("")
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return models.User{}, errors.New("")
	}

	var user models.User

	consts.DB.Find(&user, models.User{
		ID: claims.Id,
	})

	return user, nil
}
