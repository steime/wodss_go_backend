package handler

import (
	"encoding/json"
	"fmt"
	"github.com/steime/wodss_go_backend/util"
	"io/ioutil"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/go-sql-driver/mysql"
)


func CreateStudent(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var student persistence.Student
		json.Unmarshal(reqBody, &student)
		_,error :=repository.CreateStudent(&student)
		if error != nil {
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
		err := json.NewDecoder(r.Body).Decode(student)
		if err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := repository.FindOne(student.Email, student.Password)
		json.NewEncoder(w).Encode(resp)
	}
}

func GetStudentById(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if checked {
			student, err := repository.GetStudentById(id)
			if err == nil {
				json.NewEncoder(w).Encode(student)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

func UpdateStudent(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checked,id := util.CheckID(r)
		if !checked {
			w.WriteHeader(http.StatusBadRequest)
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		student := &persistence.Student{}
		err := json.Unmarshal(reqBody, &student)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		validate := validator.New()
		err = validate.Struct(student)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			updStudent, err := repository.UpdateStudent(id,student)
			if err !=nil {
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
		if checked {
			error := repository.DeleteStudent(id)
			if error != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusAccepted)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
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