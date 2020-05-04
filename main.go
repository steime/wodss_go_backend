// Main entry point to server
package main

import (
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
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

	// Setup CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET","DELETE", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	// Create or open requestLog.txt file
	f,err := os.OpenFile("requestLog.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		log.Print(err)
	}
	// Install CORS and Logging Middleware
	corsRouter := c.Handler(router.NewRouter(repository))
	loggedRouter := handlers.LoggingHandler(f, corsRouter)

  // Start HTTP
  go func() {
    err_http := http.ListenAndServe(":3000", loggedRouter)
    if err_http != nil {
		log.Printf("Server started on port 8080...\n")
        log.Fatal("Web server (HTTP): ", err_http)
    }
   }()

  //  Start HTTPS
  err = http.ListenAndServeTLS(":3001", "/etc/letsencrypt/live/wodss.xyz/fullchain.pem", "/etc/letsencrypt/live/wodss.xyz/privkey.pem", loggedRouter)
  if err != nil {
	  log.Printf("Server started on port 8081...\n")
      log.Fatal("Web server (HTTPS): ", err)
  }

}

