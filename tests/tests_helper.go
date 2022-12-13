package tests

import (
	"bytes"
	"go_auth/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
