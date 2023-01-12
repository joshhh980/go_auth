package tests

import (
	"encoding/json"
	"go_auth/consts"
	"go_auth/handlers"
	"go_auth/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

var loginParams = []byte(`{
	"email":"example@mail.com",
	"password":"12345678"
}`)

var loginParamsBlankEmail = []byte(`{
	"email":"",
	"password":"12345678"
}`)

var loginParamsInvalidEmail = []byte(`{
	"email":"example",
	"password":"12345678"
}`)

var loginParamsBlankPassword = []byte(`{
	"email":"example@mail.com",
	"password":""
}`)

var invalidloginParamsEmail = []byte(`{
	"email":"email@mail.com",
	"password":"12345678"
}`)

var invalidloginParamsPassword = []byte(`{
	"email":"example@mail.com",
	"password":"123456782"
}`)

type LoginTestSuite struct {
	suite.Suite
}

func handleLogin(jsonStr []byte) ([]byte, *httptest.ResponseRecorder) {
	req, w := makeRequest("/login", jsonStr)
	handlers.LoginHandler(w, req)
	return handle(jsonStr, w)
}

func (suite *LoginTestSuite) BeforeTest(suiteName, testName string) {
	consts.InitializeDB()
	user = buildUser()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	_user := &user
	consts.DB.Create(_user)
}

func (suite *LoginTestSuite) TestValidLogin() {
	data, r := handleLogin(loginParams)
	_user := &responses.UserResponse{}
	json.Unmarshal(data, _user)
	token := r.Result().Header.Get("Authorization")
	suite.Equal("example@mail.com", _user.Email)
	suite.NotNil(token)
	suite.Equal(http.StatusOK, r.Code)
}

func (suite *LoginTestSuite) TestInvalidLogin() {

	var tests = []struct {
		input    []byte
		key      string
		expected string
	}{
		{loginParamsBlankEmail, "email", "Email can't be blank"},
		{loginParamsInvalidEmail, "email", "Email must be valid"},
		{loginParamsBlankPassword, "password", "Password can't be blank"},
	}

	for _, test := range tests {
		data, r := handleLogin(test.input)
		body := map[string]interface{}{}
		json.Unmarshal(data, &body)
		suite.Contains(body[test.key], test.expected)
		suite.Equal(http.StatusBadRequest, r.Code)
	}
}

func (suite *LoginTestSuite) TestInvalidLoginCredentails() {

	var tests = []struct {
		input    []byte
		expected string
	}{
		{invalidloginParamsEmail, "Invalid Email or Password"},
		{invalidloginParamsPassword, "Invalid Email or Password"},
	}

	for _, test := range tests {
		data, r := handleLogin(test.input)
		body := map[string]interface{}{}
		json.Unmarshal(data, &body)
		suite.Contains(body["errors"], test.expected)
		suite.Equal(http.StatusUnauthorized, r.Code)
	}
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
