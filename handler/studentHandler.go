package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/util"
	"net/http"
	"strconv"

	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/go-sql-driver/mysql"
)

func CreateStudent(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createBody := &persistence.CreateStudentBody{}
		student := &persistence.Student{}
		if err := json.NewDecoder(r.Body).Decode(createBody); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			validate := validator.New()
			if err := validate.Struct(createBody); err != nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				student.Password = createBody.Password
				student.Email = createBody.Email
				student.Degree = createBody.Degree
				student.Semester = createBody.Semester
				if _, err := repository.CreateStudent(student); err != nil {
					util.LogErrorAndSendBadRequest(w,r,err)
				} else {
					http.Redirect(w, r, r.Header.Get("Referer"), 201)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(student)
				}
			}
		}
	}
}

func GetStudentById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if checked,id := util.CheckID(r); !checked {
			util.LogErrorAndSendBadRequest(w,r,errors.New("param id doesn't match token id"))
		} else if student, err := repository.GetStudentById(id); err != nil || !checked {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
		}
	})
}

func UpdateStudent(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		student := &persistence.Student{}
		err := json.NewDecoder(r.Body).Decode(student)
		if err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			vars := mux.Vars(r)
			paramId := vars["id"]
			bodyId := strconv.Itoa(int(student.ID))
			if !checked || paramId != bodyId {
				util.LogErrorAndSendBadRequest(w,r,errors.New("param and body id mismatch"))
			} else {
				validate := validator.New()
				if err = validate.Struct(student); err != nil {
					util.LogErrorAndSendBadRequest(w,r,err)
				} else {
					if updStudent, err := repository.UpdateStudent(id,student); err !=nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(updStudent)
					}
				}
			}
		}
	})
}

func DeleteStudent(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if !checked {
			util.LogErrorAndSendBadRequest(w,r,errors.New("student's can only delete their own account"))
		} else if err := repository.DeleteStudent(id); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})
}
