package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsersHand(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
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


func AddUserHand(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var user persistence.User
		json.Unmarshal(reqBody, &user)
		repository.AddUser(&user)
		http.Redirect(w, r, r.Header.Get("Referer"), 200)
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