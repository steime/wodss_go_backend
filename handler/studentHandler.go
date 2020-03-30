package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/util"
	"log"
	"net/http"
	"strconv"

	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/go-sql-driver/mysql"
)

type CreateBody struct {
	Email string `json:"email,omitempty" validate:"required,email,min=6,max=320"`
	Password string `json:"password,omitempty" validate:"required,min=10"`
	Semester string `json:"semester,omitempty" validate:"required"`
	Degree string `json:"degree,omitempty" validate:"required,alpha"`
}

func CreateStudent(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createBody := &CreateBody{}
		student := &persistence.Student{}
		if err := json.NewDecoder(r.Body).Decode(createBody); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			validate := validator.New()
			if err := validate.Struct(createBody); err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				student.Password = createBody.Password
				student.Email = createBody.Email
				student.Degree = createBody.Degree
				student.Semester = createBody.Semester
				if _, error := repository.CreateStudent(student); error != nil {
					log.Print(error)
					w.WriteHeader(http.StatusBadRequest)
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
		checked,id := util.CheckID(r)
		if student, err := repository.GetStudentById(id); err != nil || !checked {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
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
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			vars := mux.Vars(r)
			paramId := vars["id"]
			bodyId := strconv.Itoa(int(student.ID))
			if !checked || paramId != bodyId {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				validate := validator.New()
				if err = validate.Struct(student); err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusBadRequest)
				} else {
					if updStudent, err := repository.UpdateStudent(id,student); err !=nil {
						log.Print(err)
						w.WriteHeader(http.StatusBadRequest)
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
		if error := repository.DeleteStudent(id); error != nil || !checked {
			log.Print(error)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})
}
