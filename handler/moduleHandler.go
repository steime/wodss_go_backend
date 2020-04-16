package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
)

func GetAllModules(repository persistence.Repository)func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if modules, error := repository.GetAllModules(); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp []persistence.ModuleResponse
			for _ , module := range modules {
				resp = append(resp, ModuleResponseBuilder(module))
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func GetAllModulesByDegree(repository persistence.Repository)func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		degreeID := params["degree"]
		var resp []persistence.ModuleResponse
		if degree, error := repository.GetDegreeById(degreeID); error != nil  {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			for _, degreeGroup := range degree.Groups {
				if group, error := repository.GetModuleGroupById(degreeGroup.GroupID); error != nil {
					log.Print(error)
					w.WriteHeader(http.StatusBadRequest)
				} else {
					for _, moduleList := range group.ModulesList {
						if module, error := repository.GetModuleById(moduleList.ModuleID); error != nil {
							log.Print(error)
							w.WriteHeader(http.StatusBadRequest)
						} else {
							resp = append(resp, ModuleResponseBuilder(module))
						}
					}

				}
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func GetModuleById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if module, error := repository.GetModuleById(id); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ModuleResponseBuilder(module))
		}
	}
}

func ModuleResponseBuilder(module persistence.Module) persistence.ModuleResponse {
	var resp persistence.ModuleResponse
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
	return resp
}
