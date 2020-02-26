package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"


	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	db *gorm.DB
	Connected bool
}

/*
func (r *MySqlRepository) GetAllUsers() ([]persistence.User) {
	var users []persistence.User
	res := r.db.Find(&users)

	return res
}
*/

func (r *MySqlRepository) Connect(name string, password string, database string) {
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", name, password, database)
	var err error
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
	r.Connect("steime", "steime", "user")
	if !r.db.HasTable(&persistence.User{}) {
		r.db.Debug().AutoMigrate(&persistence.User{})
	}
	return &r
}




