package tests

import (
	"encoding/json"
	"go_auth/consts"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	consts.InitializeDB()
	user = buildUser()
}

func TestNotAValidEmailLogin(t *testing.T) {
	user.Email = "example"
	jsonStr, _ := json.Marshal(user)
	data, r := handleLogin(jsonStr)
	notAValidEmail(t, data, r)
	cleanUp(t)
}

func TestEmailCantBeBlankLogin(t *testing.T) {
	user.Email = ""
	jsonStr, _ := json.Marshal(user)
	data, r := handleLogin(jsonStr)
	emailCantBeBlankLogin(t, data, r)
	cleanUp(t)
}

func TestPasswordCantBeBlankLogin(t *testing.T) {
	user.Email = "example@mail.com"
	user.Password = ""
	jsonStr, _ := json.Marshal(user)
	data, r := handleLogin(jsonStr)
	passwordCantBeBlankLogin(t, data, r)
	cleanUp(t)
}

func TestLogin(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	_user := &user
	consts.DB.Create(_user)
	var jsonStr = []byte(`{
		"email":"example@mail.com",
		"password":"12345678"
	}`)
	data, r := handleLogin(jsonStr)
	validResponse(t, data, r)
	cleanUp(t)
}

func TestInvalidEmailLogin(t *testing.T) {
	consts.DB.Create(&user)
	var jsonStr = []byte(`{
		"email":"example2@mail.com",
		"password":"12345678"
	}`)
	invalidCredentials(t, jsonStr)
	cleanUp(t)
}

func TestInvalidPasswordLogin(t *testing.T) {
	consts.DB.Create(&user)
	var jsonStr = []byte(`{
		"email":"example@mail.com",
		"password":"1234567"
	}`)
	invalidCredentials(t, jsonStr)
	cleanUp(t)
}
