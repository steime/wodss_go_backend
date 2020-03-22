package router

import (
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/handler"
	"github.com/steime/wodss_go_backend/persistence"
)

func NewRouter(repository persistence.Repository) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	r.HandleFunc("/",handler.IndexHandler()).Methods("GET")
	r.HandleFunc("/auth/login",handler.Login(repository)).Methods("POST")
	r.HandleFunc("/students", handler.CreateStudent(repository)).Methods("POST")
	r.HandleFunc("/modules",handler.GetAllModules(repository)).Methods("GET")
	r.HandleFunc("/modules/{id}",handler.GetModuleById(repository)).Methods("GET")
	r.HandleFunc("/student/{id}",JwtVerify(handler.GetStudentById(repository))).Methods("GET")
	r.HandleFunc("/student/{id}",JwtVerify(handler.UpdateStudent(repository))).Methods("PUT")
	r.HandleFunc("/student/{id}",JwtVerify(handler.DeleteStudent(repository))).Methods("DELETE")
	r.HandleFunc("/auth/refresh",JwtVerify(handler.RefreshToken(repository))).Methods("POST")

	return r
}