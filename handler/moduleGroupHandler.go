package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"
)

func GetAllModuleGroups(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if moduleGroups , error := repository.GetAllModuleGroups(); error !=nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(moduleGroups)
		}
	})
}

func GetModuleGroupById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if moduleGroup, error := repository.GetModuleGroupById(id); error != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(moduleGroup)
		}
	})
}
