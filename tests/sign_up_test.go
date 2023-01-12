package tests

import (
	"encoding/json"
	"go_auth/consts"
	"go_auth/handlers"
	"go_auth/models"
	"go_auth/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignUpTestSuite struct {
	suite.Suite
}

func (suite *SignUpTestSuite) BeforeTest(suiteName, testName string) {
	consts.InitializeDB()
	consts.DB.Exec(`DROP TABLE users;`)
	consts.InitializeDB()
	consts.DB.AutoMigrate(&models.User{})
	user = buildUser()
}

func handleSignUp(jsonStr []byte) ([]byte, *httptest.ResponseRecorder) {
	req, w := makeRequest("/sign_up", jsonStr)
	handlers.SignUpHandler(w, req)
	return handle(jsonStr, w)
}

func (suite *SignUpTestSuite) TestValidSignUp() {
	data, r := handleSignUp(loginParams)
	_user := &responses.UserResponse{}
	err := json.Unmarshal(data, _user)
	suite.Nil(err)
	token := r.Result().Header.Get("Authorization")
	suite.Equal("example@mail.com", _user.Email)
	suite.NotNil(token)
	suite.Equal(http.StatusOK, r.Code)
}

func (suite *SignUpTestSuite) TestUserAlreadyExists() {
	consts.DB.Create(&user)
	suite.NotNil(user.Email)
	data, r := handleSignUp(loginParams)
	response := map[string]interface{}{}
	json.Unmarshal(data, &response)
	suite.Contains(response["errors"], "Email already exists")
	suite.Equal(http.StatusBadRequest, r.Code)
}

func (suite *LoginTestSuite) TestInvalidSignUp() {

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
		data, r := handleSignUp(test.input)
		body := map[string]interface{}{}
		json.Unmarshal(data, &body)
		suite.Contains(body[test.key], test.expected)
		suite.Equal(http.StatusBadRequest, r.Code)
	}
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpTestSuite))
}
