package tests

import (
	"bytes"
	"encoding/json"
	"go_auth/consts"
	"go_auth/handlers"
	"go_auth/helpers"
	"go_auth/responses"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type ValidateTokenTestSuite struct {
	suite.Suite
}

func handleValidateToken(jsonStr []byte, token string) ([]byte, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/validate_token", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	handlers.ValidateTokenHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return data, w
}

func (suite *ValidateTokenTestSuite) BeforeTest(suiteName, testName string) {
	consts.InitializeDB()
	user = buildUser()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	_user := &user
	consts.DB.Create(_user)
}

func (suite *ValidateTokenTestSuite) TestValidValidateToken() {
	token, _ := helpers.SignToken(user)
	data, r := handleValidateToken([]byte{}, token)
	_user := &responses.UserResponse{}
	json.Unmarshal(data, _user)
	suite.Equal("example@mail.com", _user.Email)
	suite.Equal("Example", _user.Name)
	suite.Equal(http.StatusOK, r.Code)
}

func (suite *ValidateTokenTestSuite) TestInvalidValidValidateToken() {
	token := "token"
	_, r := handleValidateToken([]byte{}, token)
	suite.Equal(http.StatusBadRequest, r.Code)
}

func TestValidateTokenTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateTokenTestSuite))
}
