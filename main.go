package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/steime/wodss_go_backend/handler"
	"github.com/steime/wodss_go_backend/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Printf("Server started on port 8080...\n")
	name := "steime"
	pw := "steime"
	database := "user"
	Db := handler.DbConnect(name, pw, database)
	defer handler.DbClose(Db)
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter(Db)))

}
