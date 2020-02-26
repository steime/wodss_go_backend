package router

import (
	"github.com/steime/wodss_go_backend/handler"
	"github.com/steime/wodss_go_backend/persistence"

	"github.com/gorilla/mux"
)

func NewRouter(repository persistence.Repository) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user", handler.AddUserHand(repository)).Methods("POST")
	r.HandleFunc("/user", handler.GetAllUsersHand(repository)).Methods("GET")
	//r.HandleFunc("/user", handler.CreateUserHandler(repository)).Methods("POST")
	return r
}
