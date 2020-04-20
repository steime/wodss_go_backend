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
	production := os.Getenv("PRODUCTION")
	if production == "true" {
		if !r.db.HasTable(&persistence.Student{}) {
			r.db.AutoMigrate(&persistence.Student{})
		}
		if !r.db.HasTable(&persistence.Module{}) {
			r.db.AutoMigrate(&persistence.Module{})
		}
		if !r.db.HasTable(&persistence.Requirements{}) {
			r.db.AutoMigrate(&persistence.Requirements{})
		}
		if !r.db.HasTable(&persistence.ModuleGroup{}) {
			r.db.AutoMigrate(&persistence.ModuleGroup{})
		}
		if !r.db.HasTable(&persistence.ModulesList{}) {
			r.db.AutoMigrate(&persistence.ModulesList{})
		}
		if !r.db.HasTable(&persistence.Parent{}) {
			r.db.AutoMigrate(&persistence.Parent{})
		}
		if !r.db.HasTable(&persistence.Degree{}) {
			r.db.AutoMigrate(&persistence.Degree{})
		}
		if !r.db.HasTable(&persistence.Groups{}) {
			r.db.AutoMigrate(&persistence.Groups{})
		}
		if !r.db.HasTable(&persistence.ProfilesByDegree{}) {
			r.db.AutoMigrate(&persistence.ProfilesByDegree{})
		}
		if !r.db.HasTable(&persistence.ModuleVisit{}) {
			r.db.AutoMigrate(&persistence.ModuleVisit{})
		}
		if !r.db.HasTable(&persistence.Profile{}) {
			r.db.AutoMigrate(&persistence.Profile{})
		}
		if !r.db.HasTable(&persistence.ListOfModules{}) {
			r.db.AutoMigrate(&persistence.ListOfModules{})
		}
	} else {
		//For Development and Testing
		r.db.DropTableIfExists(&persistence.Module{}, &persistence.Requirements{}, &persistence.Student{}, &persistence.ModuleGroup{}, persistence.ModulesList{}, persistence.Parent{}, persistence.Degree{}, persistence.Groups{}, persistence.ProfilesByDegree{}, persistence.ModuleVisit{}, persistence.Profile{}, persistence.ListOfModules{})
		r.db.AutoMigrate(&persistence.Module{}, &persistence.Requirements{}, &persistence.Student{}, &persistence.ModuleGroup{}, persistence.ModulesList{}, persistence.Parent{}, persistence.Degree{}, persistence.Groups{}, persistence.ProfilesByDegree{}, persistence.ModuleVisit{}, persistence.Profile{}, persistence.ListOfModules{})
	}
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
