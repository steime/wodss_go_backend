package handler

import (
	"encoding/json"
	"github.com/steime/wodss_go_backend/persistence"
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
	})

}
