package api

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/check-username", CheckUsernameHandler).Methods("GET")
	r.HandleFunc("/create-username", CreateUsernameHandler).Methods("POST")
	return r
}
