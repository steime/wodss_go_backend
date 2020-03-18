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
	if !r.db.HasTable(&persistence.Student{}) {
		r.db.Debug().AutoMigrate(&persistence.Student{})
	}

	//For Development
	r.db.Debug().DropTableIfExists(&persistence.Module{},&persistence.Requirements{},&persistence.Group{})
	r.db.Debug().AutoMigrate(&persistence.Module{},&persistence.Requirements{},&persistence.Group{})

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

func (r *MySqlRepository) CreateStudent(student *persistence.Student) (*persistence.Student,error){
	if r.CheckIfEmailExists(student.Email) {
		pass, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err.Error())
		}
		student.Password = string(pass)
		r.db.Create(&student)
		return student,nil
	} else {
		return student,errors.New("email already used")
	}
}

func (r *MySqlRepository) UpdateStudent(id string, student *persistence.Student) (*persistence.Student,error){
	oldStudent := r.GetStudentById(id)
	oldStudent.ID = student.ID
	oldStudent.Email = student.Email
	oldStudent.Degree = student.Degree
	oldStudent.Semester = student.Semester
	r.db.Save(oldStudent)
	return &oldStudent,nil
}

func (r *MySqlRepository) GetAllStudents() []persistence.Student {
	var students []persistence.Student
	r.db.Find(&students).Rows()
	return students
}

func (r *MySqlRepository) FindOne(email, password string) map[string]interface{} {
	student := &persistence.Student{}

	if err := r.db.Where("Email = ?", email).First(student).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &persistence.Token{
		StudentID: student.ID,
		Email:  student.Email,
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
	resp["student"] = student
	return resp
}

func (r *MySqlRepository) GetStudentById(id string) persistence.Student {
	var student persistence.Student
	i, err := strconv.Atoi(id)
	r.db.First(&student,i).Scan(&student)
	if err != nil {
		panic(err.Error())
	}

	return student
}

func (r *MySqlRepository) CheckIfEmailExists(email string) bool {
	student := &persistence.Student{}
	if err := r.db.Where("Email = ?", email).First(student).Error; err != nil {
		return true
	}
	return false
}

func (r *MySqlRepository) SaveAllModules(modules []persistence.Module) {
	for _, m := range modules {
		fmt.Println(m)
		r.db.Create(&m)
	}
	//r.db.Create(&modules)
}

func (r *MySqlRepository) GetAllModules() []persistence.Module{
	var modules []persistence.Module
	r.db.Preload("Requirements").Find(&modules)
	return modules
}

func (r *MySqlRepository) GetModuleById(id string) persistence.Module{
	var module persistence.Module
	i, err := strconv.Atoi(id)
	r.db.Preload("Requirements").Find(&module,i)
	if err != nil {
		panic(err)
	}

	return module
}
