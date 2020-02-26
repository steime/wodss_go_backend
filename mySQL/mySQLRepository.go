package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	db *gorm.DB
	Connected bool
}

func (r *MySqlRepository) GetAllUsers() []persistence.User {
	var users []persistence.User
	rows, err := r.db.Find(&users).Rows()

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var user persistence.User
		r.db.ScanRows(rows, &user)
		users = append(users, user)
	}
	return users
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
	/*
		if err = r.db.Ping(); err != nil {
			panic(err)
		}
	*/
	r.Connected = true
}

func (r *MySqlRepository) AddUser(user *persistence.User) {
	r.db.Create(&user)
}

func NewMySqlRepository() *MySqlRepository {
	r := MySqlRepository{}
	r.Connect()
	if !r.db.HasTable(&persistence.User{}) {
		r.db.Debug().AutoMigrate(&persistence.User{})
	}
	return &r
}




