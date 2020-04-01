package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
)

type DegreeResponse struct {
	ID string
	Name string
	Groups []string
}

func GetAllDegrees(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if degrees , error := repository.GetAllDegrees(); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp []DegreeResponse
			var deg DegreeResponse
			for _ , degree := range degrees {
				deg.ID = degree.ID
				deg.Name = degree.Name
				for _ , g := range degree.Groups {
					deg.Groups = append(deg.Groups,g.GroupID)
				}
				resp = append(resp,deg)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	})
}

func GetDegreeById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if degree, error := repository.GetDegreeById(id); error != nil  {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp DegreeResponse
			resp.ID = degree.ID
			resp.Name = degree.Name
			for _ , g := range degree.Groups {
				resp.Groups = append(resp.Groups,g.GroupID)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	})
}
