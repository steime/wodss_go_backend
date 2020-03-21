package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"
)

func GetAllModules(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		modules := repository.GetAllModules()

		if json, err := json.Marshal(modules); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func GetModuleById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		module := repository.GetModuleById(id)
		json.NewEncoder(w).Encode(module)
	}
}
