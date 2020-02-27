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

func GetAllUsers(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users := repository.GetAllUsers()

		if json, err := json.Marshal(users); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func GetUserById(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user := repository.FindById(id)
		json.NewEncoder(w).Encode(user)
	}
}


func AddUser(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var user persistence.User
		json.Unmarshal(reqBody, &user)
		repository.CreateUser(&user)
		http.Redirect(w, r, r.Header.Get("Referer"), 200)
	}
}

func Login(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	user := &persistence.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := repository.FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
	}
}

/*
func CreateUserHandler(Db *sql.DB) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "%+v", string(reqBody))
		var user persistence.User
		json.Unmarshal(reqBody, &user)
		id := user.ID
		name := user.Name
		insForm, err := Db.Prepare("INSERT INTO users(id, name) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(id, name)
		//log.Println("INSERT: id: " + id + " | name: " + name)
		http.Redirect(w, r, "/", 301)
	}
}*/