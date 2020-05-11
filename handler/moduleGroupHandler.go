// ModuleGroup Handler functions for /modulegroup routes
package handler

import (
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"net/http"
)
// /moduleGroup endpoint with degree query param
func GetAllModuleGroups(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		degreeID := r.FormValue("degree")
		emptyString := ""
		var resp []persistence.ModuleGroupsResponse
		if degreeID == emptyString {
			if moduleGroups , err := repository.GetAllModuleGroups(); err !=nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				for _,group := range moduleGroups {
					resp = append(resp, ModuleGroupResponseBuilder(group))
				}
				util.EncodeJSONandSendResponse(w,r,resp)
			}
		} else {
			if degree, err := repository.GetDegreeById(degreeID); err != nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				for _, degreeGroup := range degree.Groups {
					if group, err := repository.GetModuleGroupById(degreeGroup.GroupID); err != nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						resp = append(resp, ModuleGroupResponseBuilder(group))
					}
				}
				util.EncodeJSONandSendResponse(w,r,resp)
			}
		}
	}
}

func GetModuleGroupById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if moduleGroup, err := repository.GetModuleGroupById(id); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			util.EncodeJSONandSendResponse(w,r,ModuleGroupResponseBuilder(moduleGroup))
		}
	}
}

func ModuleGroupResponseBuilder(group persistence.ModuleGroup) persistence.ModuleGroupsResponse{
	var moduleGroupResponse persistence.ModuleGroupsResponse
	emptyString := ""
	emptyList := make([]string, 0)
	moduleGroupResponse.ID = group.ID
	moduleGroupResponse.Name = group.Name
	moduleGroupResponse.Minima = group.Minima
	if group.Parent.Parent == &emptyString {
		moduleGroupResponse.Parent = nil
	} else {
		moduleGroupResponse.Parent = group.Parent.Parent
	}
	if len(group.ModulesList) > 0 {
		for _, m := range group.ModulesList {
			moduleGroupResponse.ModulesList = append(moduleGroupResponse.ModulesList, m.ModuleID)
		}
	} else {
		moduleGroupResponse.ModulesList = emptyList
	}
	return moduleGroupResponse
}
