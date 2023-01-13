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
	r.HandleFunc("/user", handlers.ShowUserHandler).Methods("GET")
	r.HandleFunc("/user", handlers.UpdateUserHandler).Methods("PUT")
	http.ListenAndServe(":3000", muxHandler.CORS(
		muxHandler.AllowedOrigins([]string{"*"}),
		muxHandler.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		muxHandler.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
		muxHandler.ExposedHeaders([]string{"Authorization"}),
	)(r))
}
