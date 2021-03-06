// Module Handler functions for /module routes
package handler

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"net/http"
	"os"
)
// /module Endpoint with support for degree and canVisit support
func GetAllModules(repository persistence.Repository)func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		degreeID := r.FormValue("degree")
		canVisit := r.FormValue("canVisit")
		emptyString := ""
		var resp []persistence.ModuleResponse
		if degreeID == emptyString &&  canVisit == emptyString{
			if resp, err := BuildModuleResponse(repository); err != nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				util.EncodeJSONandSendResponse(w,r,resp)
			}
		} else if canVisit == "true" && degreeID == emptyString {
			if studId, err := CheckIfTokenIsInHeader(r); err != nil {
				util.LogErrorAndSendBadRequest(w,r, err)
			} else {
				if forbiddenModulesId, err := GetForbiddenModules(repository,studId); err != nil {
					util.LogErrorAndSendBadRequest(w,r, err)
				} else if forbiddenModulesId == nil {
					if resp, err := BuildModuleResponse(repository); err != nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						util.EncodeJSONandSendResponse(w,r,resp)
					}
				} else {
					if modules, err := repository.GetAllModules(); err != nil {
						util.LogErrorAndSendBadRequest(w,r, err)
					} else {
						util.EncodeJSONandSendResponse(w,r,BuildVisitableModulesResponse(forbiddenModulesId,modules))
					}
				}
			}
		} else if canVisit == "true" && degreeID != emptyString {
			if degree, err := repository.GetDegreeById(degreeID); err != nil  {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				if studId, err := CheckIfTokenIsInHeader(r); err != nil {
					util.LogErrorAndSendBadRequest(w,r, err)
				} else {
					if forbiddenModulesId, err := GetForbiddenModules(repository,studId); err != nil {
						util.LogErrorAndSendBadRequest(w,r, err)
					} else if forbiddenModulesId == nil{
						if resp, err = GetModulesResponseFromDegree(repository,degree); err != nil {
							util.LogErrorAndSendBadRequest(w,r,err)
						} else {
							util.EncodeJSONandSendResponse(w,r,resp)
						}
					} else {
						if modules, err := GetModulesFromDegree(repository,degree); err != nil {
							util.LogErrorAndSendBadRequest(w,r,err)
						} else {
							util.EncodeJSONandSendResponse(w,r,BuildVisitableModulesResponse(forbiddenModulesId,modules))
						}
					}
				}
			}
		} else {
			if degree, err := repository.GetDegreeById(degreeID); err != nil  {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				if resp, err = GetModulesResponseFromDegree(repository,degree); err != nil {
					util.LogErrorAndSendBadRequest(w,r,err)
				} else {
					util.EncodeJSONandSendResponse(w,r,resp)
				}
			}
		}
	}
}

func GetModuleById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if module, err := repository.GetModuleById(id); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			util.EncodeJSONandSendResponse(w,r,ModuleResponseBuilder(module))
		}
	}
}
// Checks if the Header has a token field
func CheckIfTokenIsInHeader(r *http.Request) (string,error) {
	if header, err := jwtmiddleware.FromAuthHeader(r); err != nil {
		return "-1",err
	} else {
		if token, err := jwt.Parse(header, func(token *jwt.Token) (i interface{}, err error) {
			return []byte(os.Getenv("SECRET")), nil
		}); err != nil {
			return "-1",err
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
				return "-1",err
			} else {
				return claims["sub"].(string), nil
			}
		}
	}
}
// Gets all Modules from the visits which are either passed or failed
func GetForbiddenModules(repository persistence.Repository, studId string) ([]string,error) {
	if moduleVisits, err := repository.GetAllModuleVisits(studId); err != nil {
		return nil,err
	} else {
		var forbiddenModulesId []string
		for _, moduleVisit := range moduleVisits {
			if moduleVisit.State == "passed" || moduleVisit.State == "failed" {
				forbiddenModulesId = append(forbiddenModulesId, moduleVisit.Module)
			}
		}
		encountered := map[string]bool{}
		for v:= range forbiddenModulesId {
			encountered[forbiddenModulesId[v]] = true
		}

		var result []string
		for key := range encountered {
			result = append(result, key)
		}
		return result, nil
	}
}

func GetModulesFromDegree(repository persistence.Repository,degree persistence.Degree) ([]persistence.Module,error){
	var modules []persistence.Module
	for _, degreeGroup := range degree.Groups {
		if group, err := repository.GetModuleGroupById(degreeGroup.GroupID); err != nil {
			return modules,err
		} else {
			if group.Parent.Parent == nil {
				for _, moduleList := range group.ModulesList {
					if module, err := repository.GetModuleById(moduleList.ModuleID); err != nil {
						return modules, err
					} else {
						modules = append(modules, module)
					}
				}
			}
		}
	}
	return modules, nil
}

func GetModulesResponseFromDegree(repository persistence.Repository,degree persistence.Degree) ([]persistence.ModuleResponse,error){
	var resp []persistence.ModuleResponse
	for _, degreeGroup := range degree.Groups {
		if group, err := repository.GetModuleGroupById(degreeGroup.GroupID); err != nil {
			return resp,err
		} else {
			if group.Parent.Parent == nil {
				for _, moduleList := range group.ModulesList {
					if module, err := repository.GetModuleById(moduleList.ModuleID); err != nil {
						return resp, err
					} else {
						resp = append(resp, ModuleResponseBuilder(module))
					}
				}
			}
		}
	}
	return resp, nil
}

func BuildVisitableModulesResponse(forbiddenModulesId []string,modules []persistence.Module) ([]persistence.ModuleResponse) {
	var visitableModules []persistence.Module
	var resp []persistence.ModuleResponse
	addable := false
	for _, module := range modules {
		for _, forbiddenModuleId := range forbiddenModulesId {
			if module.ID == forbiddenModuleId {
				addable = false
				break
			} else {
				addable = true
			}
		}
		if addable {
			visitableModules = append(visitableModules, module)
		}
	}
	for _, visitableModule := range visitableModules {
		resp = append(resp, ModuleResponseBuilder(visitableModule))
	}
	return resp
}

func BuildModuleResponse(repository persistence.Repository) ([]persistence.ModuleResponse,error){
	var resp []persistence.ModuleResponse
	if modules, err := repository.GetAllModules(); err != nil {
		return resp,err
	} else {
		for _, module := range modules {
			resp = append(resp, ModuleResponseBuilder(module))
		}
		return resp,nil
	}
}

func ModuleResponseBuilder(module persistence.Module) persistence.ModuleResponse {
	var resp persistence.ModuleResponse
	emptyList := make([]string, 0)
	resp.ID = module.ID
	resp.Name = module.Name
	resp.Credits = module.Credits
	resp.Code = module.Code
	resp.Fs = module.Fs
	resp.Hs = module.Hs
	resp.Msp = module.Msp
	if len(module.Requirements) > 0 {
		for _, m := range module.Requirements {
			resp.Requirements = append(resp.Requirements, m.ReqID)
		}
	} else {
		resp.Requirements = emptyList
	}
	return resp
}
