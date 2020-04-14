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
	// Unprotected Routes
	r.HandleFunc("/",handler.IndexHandler()).Methods("GET")
	r.HandleFunc("/test",handler.StringHandler()).Methods("GET")
	r.HandleFunc("/auth/login",handler.Login(repository)).Methods("POST")
	r.HandleFunc("/auth/forgot",handler.Forgot(repository)).Queries("mail","{mail}").Methods("POST")
	r.HandleFunc("/students", handler.CreateStudent(repository)).Methods("POST")
	r.HandleFunc("/degree",handler.GetAllDegrees(repository)).Methods("GET")
	r.HandleFunc("/degree/{id}",handler.GetDegreeById(repository)).Methods("GET")
	// Protected Routes with JWT Middleware
	r.HandleFunc("/modules",JwtVerify(handler.GetAllModules(repository))).Methods("GET")
	r.HandleFunc("/modules/{id}",JwtVerify(handler.GetModuleById(repository))).Methods("GET")
	r.HandleFunc("/modulegroups",JwtVerify(handler.GetAllModuleGroups(repository))).Methods("GET")
	r.HandleFunc("/modulegroups/{id}",JwtVerify(handler.GetModuleGroupById(repository))).Methods("GET")
	r.HandleFunc("/students/{id}",JwtVerify(handler.GetStudentById(repository))).Methods("GET")
	r.HandleFunc("/students/{id}",JwtVerify(handler.UpdateStudent(repository))).Methods("PUT")
	r.HandleFunc("/students/{id}",JwtVerify(handler.DeleteStudent(repository))).Methods("DELETE")
	r.HandleFunc("/auth/refresh",JwtVerify(handler.RefreshToken(repository))).Methods("POST")
	r.HandleFunc("/modulevisits",JwtVerify(handler.CreateModuleVisit(repository))).Methods("POST")
	r.HandleFunc("/modulevisits",JwtVerify(handler.GetAllModuleVisits(repository))).Queries("student","{student}").Methods("GET")
	r.HandleFunc("/modulevisits/{id}",JwtVerify(handler.GetModuleVisitById(repository))).Methods("GET")
	r.HandleFunc("/modulevisits/{id}",JwtVerify(handler.UpdateModuleVisit(repository))).Methods("PUT")
	r.HandleFunc("/modulevisits/{id}",JwtVerify(handler.DeleteModuleVisit(repository))).Methods("DELETE")

	return r
}
