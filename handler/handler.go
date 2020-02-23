package handler

import (
	"fmt"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("User")
}
