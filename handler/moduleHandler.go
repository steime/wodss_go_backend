package handler

import (
	"encoding/json"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"log"
	"net/http"
	"strconv"
)

func GetAllModules(repository persistence.Repository)func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		degreeID := r.FormValue("degree")
		canVisit := r.FormValue("canVisit")
		emptyString := ""
		var resp []persistence.ModuleResponse
		if degreeID == emptyString &&  canVisit == emptyString{
			if modules, err := repository.GetAllModules(); err != nil {
				util.PrintErrorAndSendBadRequest(w,err)
			} else {
				for _, module := range modules {
					resp = append(resp, ModuleResponseBuilder(module))
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(resp)
			}
		} else if canVisit == "true" && degreeID == emptyString {
			if header,err := jwtmiddleware.FromAuthHeader(r); err != nil {
				util.PrintErrorAndSendBadRequest(w,err)
			} else {
				if token, err := jwt.Parse(header, func(token *jwt.Token) (i interface{}, err error) {
					return []byte("secret"), nil
				}); err != nil {
					util.PrintErrorAndSendBadRequest(w,err)
				} else {
					if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
						log.Print("Claiming problem with Token")
						w.WriteHeader(http.StatusBadRequest)
					} else {
						claimId := int(claims["sub"].(float64))
						studId := strconv.Itoa(claimId)
						if moduleVisits, err := repository.GetAllModuleVisits(studId); err != nil {
							util.PrintErrorAndSendBadRequest(w,err)
						} else {
							var forbiddenModulesId []string
							for _ , moduleVisit := range moduleVisits {
								if moduleVisit.State == "passed" || moduleVisit.State == "failed" {
									forbiddenModulesId = append(forbiddenModulesId, moduleVisit.Module)
								}
							}
							if modules, err := repository.GetAllModules(); err != nil {
								util.PrintErrorAndSendBadRequest(w,err)
							} else {
								var visitableModules []persistence.Module
								addable := false
								for _ , module := range modules {
									for _ , forbiddenModuleId := range forbiddenModulesId {
										if module.ID == forbiddenModuleId {
											addable = false
											break
										} else {
											addable = true
										}
									}
									if addable {
										visitableModules = append(visitableModules,module)
									}
								}
								for _, visitableModule := range visitableModules {
									resp = append(resp, ModuleResponseBuilder(visitableModule))
								}
								w.Header().Set("Content-Type", "application/json")
								json.NewEncoder(w).Encode(resp)
							}
						}
					}
				}
			}
		} else {
			if degree, err := repository.GetDegreeById(degreeID); err != nil  {
				util.PrintErrorAndSendBadRequest(w,err)
			} else {
				for _, degreeGroup := range degree.Groups {
					if group, err := repository.GetModuleGroupById(degreeGroup.GroupID); err != nil {
						util.PrintErrorAndSendBadRequest(w,err)
					} else {
						for _, moduleList := range group.ModulesList {
							if module, err := repository.GetModuleById(moduleList.ModuleID); err != nil {
								util.PrintErrorAndSendBadRequest(w,err)
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
}

func GetModuleById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if module, err := repository.GetModuleById(id); err != nil {
			util.PrintErrorAndSendBadRequest(w,err)
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
