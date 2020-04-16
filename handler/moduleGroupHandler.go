package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
)

func GetAllModuleGroups(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if moduleGroups , error := repository.GetAllModuleGroups(); error !=nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp []persistence.ModuleGroupsResponse
			for _,group := range moduleGroups {
				resp = append(resp,ModuleResponseBuilder(group))
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func GetAllModuleGroupsByDegree(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		degreeID := params["degree"]
		if degree, error := repository.GetDegreeById(degreeID); error != nil  {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp []persistence.ModuleGroupsResponse
			for _,degreeGroup := range degree.Groups {
				if group, error := repository.GetModuleGroupById(degreeGroup.GroupID); error != nil {
					log.Print(error)
					w.WriteHeader(http.StatusBadRequest)
				} else {
					resp = append(resp, ModuleResponseBuilder(group))
				}
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func GetModuleGroupById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if moduleGroup, error := repository.GetModuleGroupById(id); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp persistence.ModuleGroupsResponse
			epmtyString := ""
			// Parse DB Data to response format
			resp.ID = moduleGroup.ID
			resp.Name = moduleGroup.Name
			resp.Minima = moduleGroup.Minima
			if moduleGroup.Parent.Parent == &epmtyString{
				resp.Parent = nil
			} else {
				resp.Parent = moduleGroup.Parent.Parent
			}
			for _, m := range moduleGroup.ModulesList {
				resp.ModulesList = append(resp.ModulesList, m.ModuleID)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func ModuleResponseBuilder(group persistence.ModuleGroup) persistence.ModuleGroupsResponse{
	var moduleGroupResponse persistence.ModuleGroupsResponse
	emptyString := ""
	moduleGroupResponse.ID = group.ID
	moduleGroupResponse.Name = group.Name
	moduleGroupResponse.Minima = group.Minima
	if group.Parent.Parent == &emptyString {
		moduleGroupResponse.Parent = nil
	} else {
		moduleGroupResponse.Parent = group.Parent.Parent
	}
	for _, m := range group.ModulesList {
		moduleGroupResponse.ModulesList = append(moduleGroupResponse.ModulesList, m.ModuleID)
	}
	return moduleGroupResponse
}
