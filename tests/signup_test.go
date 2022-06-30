package tests

import (
	"bytes"
	"encoding/json"
	"go_auth/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func handleSignUp(jsonStr []byte) ([]byte, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	handlers.SignUpHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return data, w
}

func TestNotAValidEmail(t *testing.T) {
	var jsonStr = []byte(`{
		"email":"example"
	}`)
	data, _ := handleSignUp(jsonStr)
	s := string(data)
	assert.Equal(t, "Email must be valid\n", s)
}

func TestEmailCantBeBlank(t *testing.T) {
	var jsonStr = []byte(`{
		"password": "11.22"
	}`)
	data, _ := handleSignUp(jsonStr)
	s := string(data)
	assert.Equal(t, "Email can't be blank\n", s)
}

func TestPasswordCantBeBlank(t *testing.T) {
	var jsonStr = []byte(`{
		"email":"example@mail.com"
	}`)
	data, _ := handleSignUp(jsonStr)
	s := string(data)
	assert.Equal(t, "Password can't be blank\n", s)
}

func TestReturnsUser(t *testing.T) {
	var jsonStr = []byte(`{
		"email":"example@mail.com",
		"password":"12345678"
	}`)
	data, w := handleSignUp(jsonStr)
	var user map[string]interface{}
	json.Unmarshal(data, &user)
	s := w.Result().Header.Get("Authorization")
	assert.Equal(t, "example@mail.com", user["Email"])
	assert.NotNil(t, s)
}
