package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
)

func GetAllModuleGroups(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if moduleGroups , error := repository.GetAllModuleGroups(); error !=nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp []persistence.ModuleGroupsResponse
			var modresp persistence.ModuleGroupsResponse
			for _,group := range moduleGroups {
				modresp.ID = group.ID
				modresp.Name = group.Name
				modresp.Minima = group.Minima
				modresp.Parent = group.Parent.Parent
				for _, m := range group.ModulesList {
					modresp.ModulesList = append(modresp.ModulesList, m.ModuleID)
				}
				resp = append(resp,modresp)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	})
}

func GetModuleGroupById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if moduleGroup, error := repository.GetModuleGroupById(id); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp persistence.ModuleGroupsResponse
			// Parse DB Data to response format
			resp.ID = moduleGroup.ID
			resp.Name = moduleGroup.Name
			resp.Minima = moduleGroup.Minima
			resp.Parent = moduleGroup.Parent.Parent
			for _, m := range moduleGroup.ModulesList {
				resp.ModulesList = append(resp.ModulesList, m.ModuleID)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	})
}
