package mySQL

import (
	"errors"
	"github.com/steime/wodss_go_backend/persistence"
	"golang.org/x/crypto/bcrypt"
)

func (r *MySqlRepository) ForgotPassword(mail string) error {
	if !r.CheckIfEmailExists(mail) {
		return nil
	} else {
		return errors.New("email doesn't exist")
	}
}

func (r *MySqlRepository) ResetPassword(mail string, password string) error {
	if pass, err := bcrypt.GenerateFromPassword([]byte(password), 14) ; err != nil {
		return errors.New("password couldn't be hashed")
	} else {
		student := &persistence.Student{}
		if err := r.db.Where("Email = ?", mail).First(student).Error; err != nil {
			return errors.New("email doesn't exist")
		} else {
			if result := r.db.Model(&student).Update("password",string(pass)); result.Error !=nil {
				return errors.New("couldn't update student")
			} else {
				return nil
			}
		}
	}
}
