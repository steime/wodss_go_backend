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
	corsRouter := handlers.CORS()(router.NewRouter(repository))
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsRouter)
 
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", loggedRouter)

	if err != nil {
		log.Fatal(err)
	}
}
