package mySQL

import (
	"errors"
	"github.com/steime/wodss_go_backend/persistence"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func (r *MySqlRepository) CreateStudent(student *persistence.Student) (*persistence.Student,error){
	if r.CheckIfEmailExists(student.Email) {
		pass, err := bcrypt.GenerateFromPassword([]byte(student.Password), 14)
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
	result := r.db.Model(&student).Omit("id").Updates(persistence.Student{Email: student.Email, Semester: student.Semester, Degree: student.Degree})
	if result.Error != nil {
		return &persistence.Student{},result.Error
	}
	return student, result.Error
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

func (r *MySqlRepository) GetStudentById(id string) (persistence.Student,error) {
	var student persistence.Student
	if i, err := strconv.Atoi(id); err != nil {
		return student,err
	} else if result := r.db.First(&student,i).Scan(&student); result.Error != nil {
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


