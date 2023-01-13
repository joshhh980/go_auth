package main

import (
	"go_auth/consts"
	"go_auth/handlers"
	"net/http"
	muxHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	consts.InitializeDB()
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", handlers.SignUpHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/user", handlers.ValidateTokenHandler).Methods("GET")
	http.ListenAndServe(":3000", muxHandler.CORS(
		muxHandler.AllowedOrigins([]string{"*"}),
		muxHandler.AllowedMethods([]string{"POST", "GET"}),
		muxHandler.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
		muxHandler.ExposedHeaders([]string{"Authorization"}),
	)(r))
}
