package mySQL

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/steime/wodss_go_backend/persistence"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

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

func (r *MySqlRepository) UpdateStudent(id string, student *persistence.Student) *persistence.Student{
	r.db.Model(&student).Updates(persistence.Student{Email: student.Email, Semester: student.Semester, Degree: student.Degree})
	return student
}

func (r *MySqlRepository) DeleteStudent(id string) error {
	var student persistence.Student
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	result := r.db.First(&student,i).Scan(&student)
	if  result.Error != nil {
		return result.Error
	}
	if result = r.db.Delete(student); result.Error != nil {
		return result.Error
	}
	return result.Error

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

func (r *MySqlRepository) GetStudentById(id string) (persistence.Student,error) {
	var student persistence.Student
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	if result := r.db.First(&student,i).Scan(&student); result.Error != nil {
		return student, result.Error
	}
	return student,nil
}

func (r *MySqlRepository) CheckIfEmailExists(email string) bool {
	student := &persistence.Student{}
	if err := r.db.Where("Email = ?", email).First(student).Error; err != nil {
		return true
	}
	return false
}