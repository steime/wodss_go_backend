package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
)

func GetAllModules(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if modules, error := repository.GetAllModules(); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(modules)
		}
	})
}

func GetModuleById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if module, error := repository.GetModuleById(id); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp persistence.ModuleResponse
			//emptyString := ""
			resp.ID = module.ID
			resp.Name = module.Name
			resp.Credits = module.Credits
			resp.Code = module.Code
			resp.Fs = module.Fs
			resp.Hs = module.Hs
			resp.Msp = module.Msp
			for _ , m := range module.Requirements {
				resp.Requirements = append(resp.Requirements,m.ReqID)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	})
}
