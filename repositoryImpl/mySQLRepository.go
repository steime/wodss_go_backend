package mySQL

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"os"
)

type MySqlRepository struct {
	db *gorm.DB
	Connected bool
}

func NewMySqlRepository() *MySqlRepository {
	r := MySqlRepository{}
	r.Connect()
	/*
	if !r.db.HasTable(&persistence.Student{}) {
		r.db.Debug().AutoMigrate(&persistence.Student{})
	}

	 */

	//For Development
	r.db.Debug().DropTableIfExists(&persistence.Module{},&persistence.Requirements{},&persistence.Student{},&persistence.ModuleGroup{},persistence.ModulesList{},persistence.Parent{},persistence.Degree{},persistence.Groups{},persistence.ModuleVisit{},persistence.Profile{},persistence.ListOfModules{})
	r.db.Debug().AutoMigrate(&persistence.Module{},&persistence.Requirements{},&persistence.Student{},&persistence.ModuleGroup{},persistence.ModulesList{},persistence.Parent{},persistence.Degree{},persistence.Groups{},persistence.ModuleVisit{},persistence.Profile{},persistence.ListOfModules{})

	return &r
}

func (r *MySqlRepository) Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	name := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", name, password, database)
	//var err error
	if r.db, err = gorm.Open("mysql", dataSourceName); err != nil {
		panic(err)
	}
	r.Connected = true
}
