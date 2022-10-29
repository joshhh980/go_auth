package tests

import (
	"encoding/json"
	"go_auth/consts"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	consts.InitializeDB()
	user = buildUser()
}

func TestNotAValidEmail(t *testing.T) {
	user.Email = "example"
	jsonStr, _ := json.Marshal(user)
	data, r := handleSignUp(jsonStr)
	notAValidEmail(t, data, r)
}

func TestEmailCantBeBlank(t *testing.T) {
	user.Email = ""
	jsonStr, _ := json.Marshal(user)
	data, r := handleSignUp(jsonStr)
	emailCantBeBlankLogin(t, data, r)
}

func TestPasswordCantBeBlank(t *testing.T) {
	user.Email = "example@mail.com"
	user.Password = ""
	jsonStr, _ := json.Marshal(user)
	data, r := handleSignUp(jsonStr)
	passwordCantBeBlankLogin(t, data, r)
}

func TestSignUp(t *testing.T) {
	var jsonStr = []byte(`{
		"email":"example@mail.com",
		"password":"12345678"
	}`)
	data, r := handleSignUp(jsonStr)
	validResponse(t, data, r)
	cleanUp(t)
}

func TestUserAlreadyExists(t *testing.T) {
	consts.DB.Create(&user)
	var jsonStr = []byte(`{
		"email":"example@mail.com",
		"password":"12345678"
	}`)
	data, r := handleSignUp(jsonStr)
	s := string(data)
	assert.Equal(t, "Email already exists\n", s)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	cleanUp(t)
}
