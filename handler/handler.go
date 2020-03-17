package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"

	_ "github.com/go-sql-driver/mysql"
)

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

func GetStudentById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		student := repository.FindById(id)
		json.NewEncoder(w).Encode(student)
	}
}

func AddStudent(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var student persistence.Student
		json.Unmarshal(reqBody, &student)
		_,error :=repository.CreateStudent(&student)
		if error != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
		}
		http.Redirect(w, r, r.Header.Get("Referer"), 200)
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

func GetAllModules(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		modules := repository.GetAllModules()

		if json, err := json.Marshal(modules); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func GetModuleById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		module := repository.GetModuleById(id)
		json.NewEncoder(w).Encode(module)
	}
}