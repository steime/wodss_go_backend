package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/steime/wodss_go_backend/mySQL"
	"github.com/steime/wodss_go_backend/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Printf("Server started on port 8080...\n")
	repository := mySQL.NewMySqlRepository()
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter(repository)))
}
