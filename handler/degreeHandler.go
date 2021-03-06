// Degree Handler functions for /degree routes
package handler

import (
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
			util.EncodeJSONandSendResponse(w,r,resp)
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
			util.EncodeJSONandSendResponse(w,r,DegreeResponseBuilder(degree))
		}
	}
}

func DegreeResponseBuilder(degree persistence.Degree) persistence.DegreeResponse {
	var resp persistence.DegreeResponse
	emptyList := make([]string, 0)
	resp.ID = degree.ID
	resp.Name = degree.Name
	if len(degree.Groups) > 0 {
		for _, g := range degree.Groups {
			resp.Groups = append(resp.Groups, g.GroupID)
		}
	} else {
		resp.Groups = emptyList
	}
	if len(degree.ProfilesByDegree) > 0 {
		for _, p := range degree.ProfilesByDegree {
			resp.Profiles = append(resp.Profiles, p.ProfileID)
		}
	} else {
		resp.Groups = emptyList
	}
	return resp
}
