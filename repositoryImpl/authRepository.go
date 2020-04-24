package mySQL

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/steime/wodss_go_backend/persistence"
	"golang.org/x/crypto/bcrypt"
	"time"
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
		return err
	} else {
		student := &persistence.Student{}
		if err := r.db.Where("Email = ?", mail).First(student).Error; err != nil {
			return err
		} else {
			if result := r.db.Model(&student).Update("password",string(pass)); result.Error !=nil {
				return result.Error
			} else {
				return nil
			}
		}
	}
}

func (r *MySqlRepository) FindOne(email, password string) (persistence.TokenPair,error) {
	student := &persistence.Student{}
	tokenPair := persistence.TokenPair{}
	if err := r.db.Where("Email = ?", email).First(student).Error; err != nil {
		return tokenPair,err
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return tokenPair,err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = student.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, error := token.SignedString([]byte("secret"))
	if error != nil {
		return tokenPair,err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = student.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return tokenPair,err
	}
	tokenPair.Token = t
	tokenPair.RefreshToken = rt
	return tokenPair,nil
}
