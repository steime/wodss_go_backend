package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"

	_ "github.com/go-sql-driver/mysql"
)
/*
type handler struct {
	repo persistence.Repository
}

func NewHandler(repo persistence.Repository) *handler {
	return &handler{
		repo: repo,
	}
}
*/
func GetAllUsersHand(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User")
		users, err := repository.GetAllUsers()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		if json, err := json.Marshal(users); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

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
}