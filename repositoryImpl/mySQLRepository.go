// Creates repository, either in production or testing/development mode
package mySQL

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
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
		// Production Mode, Tables are Created, if they don't exist
		// Nested Tables are in the same if clause with the parent table
		// Existing Tables are updated
		log.Print("prod")
		if !r.db.HasTable(&persistence.Student{}) {
			r.db.AutoMigrate(&persistence.Student{})
		}
		if !r.db.HasTable(&persistence.Module{}) {
			r.db.DropTableIfExists(&persistence.Module{},&persistence.Requirements{})
			r.db.AutoMigrate(&persistence.Module{},&persistence.Requirements{})
			util.FetchAllModules(&r)
		}
		util.UpdateAllModules(&r)
		if !r.db.HasTable(&persistence.ModuleGroup{}) {
			r.db.DropTableIfExists(&persistence.ModuleGroup{},&persistence.ModulesList{},&persistence.Parent{})
			r.db.AutoMigrate(&persistence.ModuleGroup{},&persistence.ModulesList{},&persistence.Parent{})
			util.FetchAllModuleGroups(&r)
		}
		util.UpdateAllModuleGroups(&r)
		if !r.db.HasTable(&persistence.Degree{}) {
			r.db.DropTableIfExists(&persistence.Degree{},&persistence.Groups{},&persistence.ProfilesByDegree{})
			r.db.AutoMigrate(&persistence.Degree{},&persistence.Groups{},&persistence.ProfilesByDegree{})
			util.FetchAllDegrees(&r)
		}
		util.UpdateAllDegrees(&r)
		if !r.db.HasTable(&persistence.ModuleVisit{}) {
			r.db.AutoMigrate(&persistence.ModuleVisit{})
		}
		if !r.db.HasTable(&persistence.Profile{}) {
			r.db.DropTableIfExists(&persistence.Profile{},&persistence.ListOfModules{})
			r.db.AutoMigrate(&persistence.Profile{},&persistence.ListOfModules{})
			util.FetchAllProfiles(&r)
		}
		util.UpdateAllProfiles(&r)
		// Create Cron Job
		util.CronJob(&r)
	} else {
		//For Development and Testing, Drops all Tables and Recreates them
		r.db.DropTableIfExists(
			&persistence.Module{},
			&persistence.Requirements{},
			&persistence.Student{},
			&persistence.ModuleGroup{},
			&persistence.ModulesList{},
			&persistence.Parent{},
			&persistence.Degree{},
			&persistence.Groups{},
			&persistence.ProfilesByDegree{},
			&persistence.ModuleVisit{},
			&persistence.Profile{},
			&persistence.ListOfModules{})
		r.db.AutoMigrate(
			&persistence.Module{},
			&persistence.Requirements{},
			&persistence.Student{},
			&persistence.ModuleGroup{},
			&persistence.ModulesList{},
			&persistence.Parent{},
			&persistence.Degree{},
			&persistence.Groups{},
			&persistence.ProfilesByDegree{},
			&persistence.ModuleVisit{},
			&persistence.Profile{},
			&persistence.ListOfModules{})
		util.FetchAllData(&r)
		log.Print("Data Loaded")
	}
	return &r
}

// Connect to DB
func (r *MySqlRepository) Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	name := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", name, password, database)
	// Chose DB Dialect for GORM
	if r.db, err = gorm.Open("mysql", dataSourceName); err != nil {
		panic(err)
	}
	r.Connected = true
}
