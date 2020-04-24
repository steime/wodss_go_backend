package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"net/http"
)

func GetAllDegrees(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if degrees , err := repository.GetAllDegrees(); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			var resp []persistence.DegreeResponse
			for _ , degree := range degrees {
				resp = append(resp,DegreeResponseBuilder(degree))
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func GetDegreeById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if degree, err := repository.GetDegreeById(id); err != nil  {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(DegreeResponseBuilder(degree))
		}
	}
}

func DegreeResponseBuilder(degree persistence.Degree) persistence.DegreeResponse {
	var resp persistence.DegreeResponse
	resp.ID = degree.ID
	resp.Name = degree.Name
	for _ , g := range degree.Groups {
		resp.Groups = append(resp.Groups,g.GroupID)
	}
	for _ , p := range degree.ProfilesByDegree {
		resp.Profiles = append(resp.Profiles,p.ProfileID)
	}
	return resp
}
