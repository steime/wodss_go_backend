package handler

import (
	"encoding/json"
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"
)

func GetAllModuleGroups(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		moduleGroups := repository.GetAllModuleGroups()

		if json, err := json.Marshal(moduleGroups); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
