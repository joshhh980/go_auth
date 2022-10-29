package tests

import (
	"bytes"
	"encoding/json"
	"go_auth/consts"
	"go_auth/handlers"
	"go_auth/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var user models.User

func buildUser() models.User {
	return models.User{
		Email:    "example@mail.com",
		Password: "12345678",
	}
}

func handle(jsonStr []byte, w *httptest.ResponseRecorder) ([]byte, *httptest.ResponseRecorder) {
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return data, w
}

func makeRequest(path string, jsonStr []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, "/sign_up", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	return req, w
}

func handleSignUp(jsonStr []byte) ([]byte, *httptest.ResponseRecorder) {
	req, w := makeRequest("/sign_up", jsonStr)
	handlers.SignUpHandler(w, req)
	return handle(jsonStr, w)
}

func handleLogin(jsonStr []byte) ([]byte, *httptest.ResponseRecorder) {
	req, w := makeRequest("/login", jsonStr)
	handlers.LoginHandler(w, req)
	return handle(jsonStr, w)
}

func notAValidEmail(t *testing.T, data []byte, r *httptest.ResponseRecorder) {
	s := string(data)
	assert.Equal(t, "Email must be valid\n", s)
	assert.Equal(t, http.StatusBadRequest, r.Code)
}

func emailCantBeBlankLogin(t *testing.T, data []byte, r *httptest.ResponseRecorder) {
	s := string(data)
	assert.Equal(t, "Email can't be blank\n", s)
	assert.Equal(t, http.StatusBadRequest, r.Code)
}

func passwordCantBeBlankLogin(t *testing.T, data []byte, r *httptest.ResponseRecorder) {
	s := string(data)
	assert.Equal(t, "Password can't be blank\n", s)
	assert.Equal(t, http.StatusBadRequest, r.Code)
}

func validResponse(t *testing.T, data []byte, r *httptest.ResponseRecorder) {
	var user map[string]interface{}
	json.Unmarshal(data, &user)
	s := r.Result().Header.Get("Authorization")
	assert.Equal(t, "example@mail.com", user["email"])
	assert.NotNil(t, s)
	assert.Equal(t, http.StatusOK, r.Code)
}

func invalidCredentials(t *testing.T, jsonStr []byte) {
	data, r := handleLogin(jsonStr)
	s := string(data)
	assert.Equal(t, "Invalid Email or Password\n", s)
	assert.Equal(t, http.StatusBadRequest, r.Code)
}

func cleanUp(t *testing.T) {
	t.Cleanup(func() {
		consts.DB.Delete(&user)
	})
}
