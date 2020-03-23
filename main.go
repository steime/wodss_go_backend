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
  
  go func() {
    err_http := http.ListenAndServe(":8080", loggedRouter)
    if err_http != nil {
        log.Fatal("Web server (HTTP): ", err_http)
    }
   }()

  //  Start HTTPS
  err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", loggedRouter)
  if err != nil {
      log.Fatal("Web server (HTTPS): ", err)
  }

}
