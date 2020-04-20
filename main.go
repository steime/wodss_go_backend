package main

import (
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/steime/wodss_go_backend/util"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/steime/wodss_go_backend/repositoryImpl"
	"github.com/steime/wodss_go_backend/router"
)

func main() {
	//Create new mySQL repository
	repository := mySQL.NewMySqlRepository()
	production := os.Getenv("PRODUCTION")
	if production == "true" {
		//TODO implement data update
	} else {
		util.FetchAllData(repository)
		log.Print("Data Loaded")
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET","DELETE", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	corsRouter := c.Handler(router.NewRouter(repository))
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsRouter)

  // Start HTTP
  go func() {
    err_http := http.ListenAndServe(":8080", loggedRouter)
    if err_http != nil {
		log.Printf("Server started on port 8080...\n")
        log.Fatal("Web server (HTTP): ", err_http)
    }
   }()

  //  Start HTTPS
  err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", loggedRouter)
  if err != nil {
	  log.Printf("Server started on port 8081...\n")
      log.Fatal("Web server (HTTPS): ", err)
  }

}
