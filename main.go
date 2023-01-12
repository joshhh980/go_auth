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
	r.HandleFunc("/sign_up", handlers.SignUpHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/validate_token", handlers.ValidateTokenHandler).Methods("GET")
	http.ListenAndServe(":8090", r)
}
