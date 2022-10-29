package helpers

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Id uint
	jwt.StandardClaims
}
