package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/steime/wodss_go_backend/router"
)

func main() {
	fmt.Printf("Server started on port 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
