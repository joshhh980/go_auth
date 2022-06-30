package main

import (
	"go_auth/consts"
	"go_auth/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	consts.InitializeDB()
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.SignUpHandler).Methods("POST")
	http.ListenAndServe(":8090", r)
}

// swagger:route POST /sign_up SignUp idSignUp
// Signs up a new user.
// responses:
//   201: signUpSuccessResponse

// swagger:parameters idSignUp
type signUpParamsWrapper struct {
	// Signs up a new user.
	// in:body
	Body handlers.SignUpRequest
}

// Successful sign up response.
// swagger:response signUpSuccessResponse
type signUpResponseWrapper struct {
	// Authorization header
	Authorization string

	// in:body
	Body handlers.UserResponse
}
