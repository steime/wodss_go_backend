package mySQL

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type MySqlRepository struct {
	db *gorm.DB
	Connected bool
}

func NewMySqlRepository() *MySqlRepository {
	r := MySqlRepository{}
	r.Connect()
	if !r.db.HasTable(&persistence.User{}) {
		r.db.Debug().AutoMigrate(&persistence.User{})
	}
	if !r.db.HasTable(&persistence.Module{}) {
		r.db.Debug().AutoMigrate(&persistence.Module{})
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

func (r *MySqlRepository) CreateUser(user *persistence.User) (*persistence.User,error){
	if r.CheckIfEmailExists(user.Email) {
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err.Error())
		}
		user.Password = string(pass)
		r.db.Create(&user)
		return user,nil
	} else {
		return user,errors.New("email already used")
	}
}

func (r *MySqlRepository) GetAllUsers() []persistence.User {
	var users []persistence.User
	r.db.Find(&users).Rows()
	return users
}

func (r *MySqlRepository) FindOne(email, password string) map[string]interface{} {
	user := &persistence.User{}

	if err := r.db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &persistence.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

func (r *MySqlRepository) FindById(id string) persistence.User {
	var user persistence.User
	i, err := strconv.Atoi(id)
	r.db.First(&user,i).Scan(&user)
	if err != nil {
		panic(err.Error())
	}

	return user
}

func (r *MySqlRepository) CheckIfEmailExists(email string) bool {
	user := &persistence.User{}
	if err := r.db.Where("Email = ?", email).First(user).Error; err != nil {
		return true
	}
	return false
}


