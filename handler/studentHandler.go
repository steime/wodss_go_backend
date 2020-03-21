package handler

import (
	"encoding/json"
	"fmt"
	"github.com/steime/wodss_go_backend/util"
	"log"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/go-sql-driver/mysql"
)


func CreateStudent(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		student := &persistence.Student{}
		if err := json.NewDecoder(r.Body).Decode(student); err != nil {
			log.Fatal(err)
		}
		if _,error :=repository.CreateStudent(student); error != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, r.Header.Get("Referer"), 201)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func Login(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		student := &persistence.Student{}
		if err := json.NewDecoder(r.Body).Decode(student); err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
			json.NewEncoder(w).Encode(resp)
		}
		resp := repository.FindOne(student.Email, student.Password)
		json.NewEncoder(w).Encode(resp)
	}
}

func GetStudentById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if !checked {
			w.WriteHeader(http.StatusBadRequest)
		}
		if student, err := repository.GetStudentById(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			json.NewEncoder(w).Encode(student)
		}
	})
}

func UpdateStudent(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if !checked {
			w.WriteHeader(http.StatusBadRequest)
		}
		student := &persistence.Student{}
		err := json.NewDecoder(r.Body).Decode(student)
		if err != nil {
			log.Fatal(err)
		}
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
	})
}

func DeleteStudent(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if !checked {
			w.WriteHeader(http.StatusBadRequest)
		}
		if error := repository.DeleteStudent(id); error != nil {
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