package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/handler"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user", handler.GetAllUsersHandler(db)).Methods("GET")
	r.HandleFunc("/user", handler.CreateUserHandler(db)).Methods("POST")
	return r
}
