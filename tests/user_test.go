package tests

import (
	"bytes"
	"encoding/json"
	"go_auth/consts"
	"go_auth/handlers"
	"go_auth/helpers"
	"go_auth/models"
	"go_auth/responses"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

var validParams = []byte(`{
	"name":"Update",
	"email":"example@mail.com"
}`)

type UserTestSuite struct {
	suite.Suite
}

func handleUser(token string, method string, jsonStr []byte, handler func(w http.ResponseWriter, r *http.Request)) ([]byte, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, "/user", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	handler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return data, w
}

func (suite *UserTestSuite) BeforeTest(suiteName, testName string) {
	consts.InitializeDB()
	user = buildUser()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	_user := &user
	consts.DB.Create(_user)
}

func (suite *UserTestSuite) TestValidUser() {
	token, _ := helpers.SignToken(user)
	data, r := handleUser(token, http.MethodGet, []byte{}, handlers.ShowUserHandler)
	_user := &responses.UserResponse{}
	json.Unmarshal(data, _user)
	suite.Equal("example@mail.com", _user.Email)
	suite.Equal("Example", _user.Name)
	suite.Equal(http.StatusOK, r.Code)
}

func (suite *UserTestSuite) TestInvalidValidUser() {
	token := "token"
	_, r := handleUser(token, http.MethodGet, []byte{}, handlers.ShowUserHandler)
	suite.Equal(http.StatusBadRequest, r.Code)
}

func (suite *UserTestSuite) TestUpdateUser() {
	token, _ := helpers.SignToken(user)
	data, r := handleUser(token, http.MethodPut, validParams, handlers.UpdateUserHandler)
	_user := &responses.UserResponse{}
	suite.Equal(http.StatusOK, r.Code)
	json.Unmarshal(data, _user)
	suite.Equal("example@mail.com", _user.Email)
	suite.Equal("Update", _user.Name)
}

func (suite *UserTestSuite) TestDeleteUser() {
	token, _ := helpers.SignToken(user)
	_, r := handleUser(token, http.MethodDelete, []byte{}, handlers.DeleteUserHandler)
	suite.Equal(http.StatusOK, r.Code)
	var _user models.User
	consts.DB.Find(&_user, models.User{
		ID: user.ID,
	})
	suite.Equal("", _user.Name)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
