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
	r.HandleFunc("/",						handler.IndexHandler())								.Methods("GET")
	r.HandleFunc("/test",					handler.StringHandler())							.Methods("GET")
	r.HandleFunc("/auth/login",			handler.Login(repository))							.Methods("POST")
	r.HandleFunc("/auth/forgot",			handler.ForgotPassword(repository))					.Methods("POST").Queries("mail","{mail}")
	r.HandleFunc("/auth/reset",			handler.ResetPassword(repository))					.Methods("POST")
	r.HandleFunc("/students", 			handler.CreateStudent(repository))					.Methods("POST")
	r.HandleFunc("/degree",				handler.GetAllDegrees(repository))					.Methods("GET")
	r.HandleFunc("/degree/{id}",			handler.GetDegreeById(repository))					.Methods("GET")
	r.HandleFunc("/modules",				handler.GetAllModules(repository))					.Methods("GET")
	r.HandleFunc("/modules/{id}",			handler.GetModuleById(repository))					.Methods("GET")
	r.HandleFunc("/modulegroups",			handler.GetAllModuleGroups(repository))				.Methods("GET")
	r.HandleFunc("/modulegroups/{id}",	handler.GetModuleGroupById(repository))				.Methods("GET")
	r.HandleFunc("/profiles",				handler.GetAllProfiles(repository))					.Methods("GET")
	r.HandleFunc("/profiles/{id}",		handler.GetProfilesById(repository))				.Methods("GET")
	// Protected Routes with JWT Middleware
	r.HandleFunc("/students/{id}",		JwtVerify(handler.GetStudentById(repository)))		.Methods("GET")
	r.HandleFunc("/students/{id}",		JwtVerify(handler.UpdateStudent(repository)))		.Methods("PUT")
	r.HandleFunc("/students/{id}",		JwtVerify(handler.DeleteStudent(repository)))		.Methods("DELETE")
	r.HandleFunc("/auth/refresh",			JwtVerify(handler.RefreshToken(repository)))		.Methods("POST")
	r.HandleFunc("/modulevisits",			JwtVerify(handler.CreateModuleVisit(repository)))	.Methods("POST")
	r.HandleFunc("/modulevisits",			JwtVerify(handler.GetAllModuleVisits(repository)))	.Methods("GET")
	r.HandleFunc("/modulevisits/{id}",	JwtVerify(handler.GetModuleVisitById(repository)))	.Methods("GET")
	r.HandleFunc("/modulevisits/{id}",	JwtVerify(handler.UpdateModuleVisit(repository)))	.Methods("PUT")
	r.HandleFunc("/modulevisits/{id}",	JwtVerify(handler.DeleteModuleVisit(repository)))	.Methods("DELETE")

	return r
}
