package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(name string, password string, db string) *sql.DB {
	Db, err := sql.Open("mysql", "steime:steime@/user?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return Db
}

func GetUserHandler(Db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("User")
		results, err := Db.Query("SELECT id, name FROM users")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var user persistence.User
		for results.Next() {

			// for each row, scan the result into our tag composite object
			err = results.Scan(&user.ID, &user.Name)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.Printf(user.Name)
		}

		if json, err := json.Marshal(user); err == nil {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
