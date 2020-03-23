package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"

	"github.com/steime/wodss_go_backend/repositoryImpl"
	"github.com/steime/wodss_go_backend/router"
	"github.com/steime/wodss_go_backend/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Printf("Server started on port 8080...\n")
	repository := mySQL.NewMySqlRepository()
	util.GetAllModules(repository)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router.NewRouter(repository))
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))

}
