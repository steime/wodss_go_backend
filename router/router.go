package router

import (

	"github.com/steime/wodss_go_backend/persistence"

	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/handler"
)

func NewRouter(repository persistence.Repository) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user", handler.GetAllUsersHand(repository)).Methods("GET")
	//r.HandleFunc("/user", handler.CreateUserHandler(repository)).Methods("POST")
	return r
}
