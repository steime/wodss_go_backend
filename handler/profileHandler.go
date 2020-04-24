package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"net/http"
)

func GetAllProfiles(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		degreeID := r.FormValue("degree")
		emptyString := ""
		var resp []persistence.ProfileResponse
		if degreeID == emptyString {
			if profiles, err := repository.GetAllProfiles(); err != nil {
				util.LogErrorAndSendBadRequest(w,r, err)
			} else {

				for _, profile := range profiles {
					resp = append(resp, ProfileResponseBuilder(profile))
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}
		} else {
			if degree, err := repository.GetDegreeById(degreeID); err != nil  {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				for _, profileList := range degree.ProfilesByDegree {
					if profile, err := repository.GetProfileById(profileList.ProfileID); err != nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						resp = append(resp, ProfileResponseBuilder(profile))
					}
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}
		}
	}
}

func GetProfilesById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if profile, err := repository.GetProfileById(id); err != nil {
			util.LogErrorAndSendBadRequest(w,r, err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ProfileResponseBuilder(profile))
		}
	}
}

func ProfileResponseBuilder(profile persistence.Profile) persistence.ProfileResponse {
	var resp persistence.ProfileResponse
	resp.ID = profile.ID
	resp.Name = profile.Name
	for _, p := range profile.ListOfModules {
		resp.ListOfModules = append(resp.ListOfModules,p.ModuleID)
	}
	resp.Minima = profile.Minima
	return resp
}
