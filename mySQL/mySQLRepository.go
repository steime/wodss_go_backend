package mySQL

import (
	"database/sql"
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	db        *sql.DB
	Connected bool
}

func (r *MySqlRepository) GetAllUsers() ([]persistence.User, error) {
	results, err := r.db.Query("SELECT id, name FROM users")
	var user persistence.User
	var users []persistence.User
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(user.Name)
		users = append(users, user)
	}
	return users,err
}

func NewMySqlRepository() *MySqlRepository {
	r := MySqlRepository{}
	r.Connect("steime", "steime", "user")
	return &r
}

func (r *MySqlRepository) Connect(name string, password string, database string) {
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", name, password, database)
	var err error
	if r.db, err = sql.Open("mysql", dataSourceName); err != nil {
		panic(err)
	}

	if err = r.db.Ping(); err != nil {
		panic(err)
	}

	r.Connected = true
}


