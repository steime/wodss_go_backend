package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

func CreateModuleVisit(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createVisit := &persistence.ModuleVisitCreateBody{}
		visit := &persistence.ModuleVisit{}
		if err := json.NewDecoder(r.Body).Decode(createVisit); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			validate := validator.New()
			if err := validate.Struct(createVisit); err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				if !util.CheckBodyID(r,createVisit.Student) {
					log.Print("Token studentId doesn't match body studentId")
					w.WriteHeader(http.StatusBadRequest)
				} else {
					visit.Semester 	= createVisit.Semester
					visit.Student 	= createVisit.Student
					visit.Module	= createVisit.Module
					visit.Grade 	= createVisit.Grade
					visit.State 	= createVisit.State
					visit.Weekday 	= createVisit.Weekday
					visit.TimeEnd 	= createVisit.TimeEnd
					visit.TimeStart = createVisit.TimeStart
					if _, error := repository.CreateModuleVisit(visit); error != nil {
						log.Print(error)
						w.WriteHeader(http.StatusBadRequest)
					} else {
						w.Header().Set("Content-Type", "application/json")
						http.Redirect(w, r, r.Header.Get("Referer"), 201)
						json.NewEncoder(w).Encode(visit)
					}
				}
			}
		}
	})
}

func GetAllModuleVisits(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		studentID := params["student"]
		if !util.CheckQueryID(r,studentID){
			log.Print("Token studentId doesn't match query studentId")
			w.WriteHeader(http.StatusBadRequest)
		} else if visits, error := repository.GetAllModuleVisits(studentID); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(visits)
		}
	})
}
