// Profile Handler functions for /profile routes
package handler

import (
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
				util.EncodeJSONandSendResponse(w,r,resp)
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
				util.EncodeJSONandSendResponse(w,r,resp)
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
			util.EncodeJSONandSendResponse(w,r,ProfileResponseBuilder(profile))
		}
	}
}

func ProfileResponseBuilder(profile persistence.Profile) persistence.ProfileResponse {
	var resp persistence.ProfileResponse
	emptyList := make([]string, 0)
	resp.ID = profile.ID
	resp.Name = profile.Name
	if len(profile.ListOfModules) > 0 {
		for _, p := range profile.ListOfModules {
			resp.ListOfModules = append(resp.ListOfModules, p.ModuleID)
		}
	} else {
		resp.ListOfModules = emptyList
	}
	resp.Minima = profile.Minima
	return resp
}
