package router

import (
	"github.com/steime/wodss_go_backend/handler"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(repository persistence.Repository) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	r.HandleFunc("/login",handler.Login(repository)).Methods("POST")
	r.HandleFunc("/students", handler.AddStudent(repository)).Methods("POST")
	r.HandleFunc("/modules",handler.GetAllModules(repository)).Methods("GET")
	r.HandleFunc("/modules/{id}",handler.GetModuleById(repository)).Methods("GET")
	//r.HandleFunc("/student", handler.GetAllStudents(repository)).Methods("GET")
	r.HandleFunc("/student/{id}",JwtVerify(handler.GetStudentById(repository))).Methods("GET")

	// Auth Routes
	/*
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(JwtVerify)

	 */

	return r
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
