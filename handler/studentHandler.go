package handler

import (
	"encoding/json"
	"fmt"
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
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Semester string `json:"semester,omitempty"`
	Degree string `json:"degree,omitempty"`
}


func CreateStudent(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createBody := &CreateBody{}
		student := &persistence.Student{}
		if err := json.NewDecoder(r.Body).Decode(createBody); err != nil {
			log.Fatal(err)
		}
		student.Password = createBody.Password
		student.Email = createBody.Email
		student.Degree = createBody.Degree
		student.Semester = createBody.Semester
		if _,error :=repository.CreateStudent(student); error != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, r.Header.Get("Referer"), 201)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func GetStudentById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if student, err := repository.GetStudentById(id); err != nil || !checked {
			w.WriteHeader(http.StatusBadRequest)
		} else {
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
			log.Fatal(err)
		}
		vars := mux.Vars(r)
		paramId := vars["id"]
		bodyId := strconv.Itoa(int(student.ID))
		if !checked || paramId != bodyId {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			validate := validator.New()
			if err = validate.Struct(student); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				if updStudent, err := repository.UpdateStudent(id,student); err !=nil {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					json.NewEncoder(w).Encode(updStudent)
				}
			}
		}
	})
}

func DeleteStudent(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if error := repository.DeleteStudent(id); error != nil || !checked {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})
}

func GetAllStudents(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		students := repository.GetAllStudents()

		if json, err := json.Marshal(students); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}