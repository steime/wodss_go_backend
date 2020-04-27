// ModuleVisit Handler functions for /modulevisit routes
package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

func CreateModuleVisit(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createVisit := &persistence.ModuleVisitCreateBody{}
		visit := &persistence.ModuleVisit{}
		if err := json.NewDecoder(r.Body).Decode(createVisit); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			validate := validator.New()
			if err := validate.Struct(createVisit); err != nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				if !util.CheckBodyID(r,createVisit.Student) {
					util.LogErrorAndSendBadRequest(w,r,errors.New("token studentId doesn't match body studentId"))
				} else {
					visit.Semester 	= createVisit.Semester
					visit.Student 	= createVisit.Student
					visit.Module	= createVisit.Module
					visit.Grade 	= createVisit.Grade
					visit.State 	= createVisit.State
					visit.Weekday 	= createVisit.Weekday
					visit.TimeEnd 	= createVisit.TimeEnd
					visit.TimeStart = createVisit.TimeStart
					if _, err := repository.CreateModuleVisit(visit); err != nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						w.Header().Set("Content-Type", "application/json")
						http.Redirect(w, r, r.Header.Get("Referer"), 201)
						json.NewEncoder(w).Encode(visit)
					}
				}
			}
		}
	})
}

func GetAllModuleVisits(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		studentID := r.FormValue("student")
		emptyString := ""
		if studentID == emptyString {
			util.LogErrorAndSendBadRequest(w,r,errors.New("query Param missing"))
		} else if !util.CheckQueryID(r,studentID){
			util.LogErrorAndSendBadRequest(w,r,errors.New("token studentId doesn't match query studentId"))
		} else if visits, err := repository.GetAllModuleVisits(studentID); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(visits)
		}
	})
}

func GetModuleVisitById (repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		visitId := params["id"]
		if studentId, err := util.GetStudentIdFromToken(r); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else if visit , err := repository.GetModuleVisitById(visitId,studentId); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(visit)
		}
	})
}

func UpdateModuleVisit (repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		visit := &persistence.ModuleVisit{}
		vars := mux.Vars(r)
		visitId := vars["id"]
		if err := json.NewDecoder(r.Body).Decode(visit); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else if studentId, err := util.GetStudentIdFromToken(r); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else if bodyId := strconv.Itoa(int(visit.ID)); bodyId != visitId {
			util.LogErrorAndSendBadRequest(w,r,errors.New("bodyId doesn't match pathId"))
		} else if bodyStudentId := strconv.Itoa(int(visit.Student)); bodyStudentId != studentId {
			util.LogErrorAndSendBadRequest(w,r,errors.New("bodyStudentId doesn't match StudentId"))
		} else {
			validate := validator.New()
			if err = validate.Struct(visit); err != nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				if updVisit, err := repository.UpdateModuleVisit(visit); err !=nil {
					util.LogErrorAndSendBadRequest(w,r,err)
				} else {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(updVisit)
				}
			}
		}
	})
}

func DeleteModuleVisit(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		visitId := vars["id"]
		if  studentId, err := util.GetStudentIdFromToken(r); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else if error := repository.DeleteModuleVisit(visitId,studentId); error != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})
}
