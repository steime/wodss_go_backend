package router

import (
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/handler"
	"github.com/steime/wodss_go_backend/persistence"
)

func NewRouter(repository persistence.Repository) *mux.Router {
	s := mux.NewRouter().StrictSlash(true)
	s.Use(CommonMiddleware)
	r := s.PathPrefix("/api").Subrouter()
	r.HandleFunc("/",handler.IndexHandler()).Methods("GET","OPTIONS")
	r.HandleFunc("/auth/login",handler.Login(repository)).Methods("POST","OPTIONS")
	r.HandleFunc("/students", handler.CreateStudent(repository)).Methods("POST","OPTIONS")
	r.HandleFunc("/modules",handler.GetAllModules(repository)).Methods("GET","OPTIONS")
	r.HandleFunc("/modules/{id}",handler.GetModuleById(repository)).Methods("GET","OPTIONS")
	r.HandleFunc("/student/{id}",JwtVerify(handler.GetStudentById(repository))).Methods("GET","OPTIONS")
	r.HandleFunc("/student/{id}",JwtVerify(handler.UpdateStudent(repository))).Methods("PUT","OPTIONS")
	r.HandleFunc("/student/{id}",JwtVerify(handler.DeleteStudent(repository))).Methods("DELETE","OPTIONS")
	r.HandleFunc("/auth/refresh",JwtVerify(handler.RefreshToken(repository))).Methods("POST","OPTIONS")

	return r
}
