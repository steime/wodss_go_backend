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
				var modResp persistence.ModuleResponse
				modResp.ID = module.ID
				modResp.Name = module.Name
				modResp.Credits = module.Credits
				modResp.Code = module.Code
				modResp.Fs = module.Fs
				modResp.Hs = module.Hs
				modResp.Msp = module.Msp
				for _ , m := range module.Requirements {
					modResp.Requirements = append(modResp.Requirements,m.ReqID)
				}
				resp = append(resp, modResp)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func GetAllModulesByDegree(repository persistence.Repository)func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if modules, error := repository.GetAllModules(); error != nil {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var resp []persistence.ModuleResponse
			for _ , module := range modules {
				var modResp persistence.ModuleResponse
				modResp.ID = module.ID
				modResp.Name = module.Name
				modResp.Credits = module.Credits
				modResp.Code = module.Code
				modResp.Fs = module.Fs
				modResp.Hs = module.Hs
				modResp.Msp = module.Msp
				for _ , m := range module.Requirements {
					modResp.Requirements = append(modResp.Requirements,m.ReqID)
				}
				resp = append(resp, modResp)
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
	}
}
