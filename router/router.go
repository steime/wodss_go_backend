package router

import (
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user", handler.GetUserHandler)
	return r
}
