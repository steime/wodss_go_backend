package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/steime/wodss_go_backend/persistence"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("User")
	user := persistence.User{Id: 1, Name: "Dominik"}

	if json, err := json.Marshal(user); err == nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(json))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
