// Auth Handler functions for /auth routes
package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/smtp"
	"time"
)

var tok persistence.ForgotTokenContext

func Login(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		student := &persistence.LoginBody{}
		if err := json.NewDecoder(r.Body).Decode(student); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			if resp, err := repository.FindOne(student.Email, student.Password); err!= nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				util.EncodeJSONandSendResponse(w,r,resp)
			}
		}
	}
}

func RefreshToken(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenReq := &persistence.TokenRequestBody{}
		if err := json.NewDecoder(r.Body).Decode(tokenReq); err !=nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		}
		token, _ := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			checked,id := util.CheckID(r)
			if student , err := repository.GetStudentById(id); err == nil || checked {
				newTokenPair, err := util.GenerateTokenPair(student.ID)
				if err != nil {
					util.LogErrorAndSendBadRequest(w,r,err)
				}
				util.EncodeJSONandSendResponse(w,r,newTokenPair)
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
	})
}

func ForgotPassword(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mail := params["mail"]
		if !util.ValidateMail(mail) {
			util.LogErrorAndSendBadRequest(w,r,errors.New("mail Address invalid"))
		} else {
			if err := repository.ForgotPassword(mail); err != nil {
				util.LogErrorAndSendBadRequest(w,r,errors.New("mail Address not existing"))
			} else {
				token := jwt.New(jwt.SigningMethodHS256)
				claims := token.Claims.(jwt.MapClaims)
				claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
				if ft, err := token.SignedString([]byte("secret")); err != nil {
					util.LogErrorAndSendBadRequest(w,r,errors.New("token Creation failed"))
				} else {
					tok.ForgotToken = ft
					auth := smtp.PlainAuth("", "wodssgoserver@gmail.com", "", "smtp.gmail.com")
					to := []string{mail}
					text := "Sie können diesen Token hier eingeben und ihr Passwort zurücksetzen \n" + ft
					msg := []byte(text)
					if err := smtp.SendMail("smtp.gmail.com:587", auth, "wodssgoserver@gmail.com", to, msg); err != nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						w.WriteHeader(http.StatusNoContent)
					}
				}
			}
		}
	}
}

func ResetPassword(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resetBody := &persistence.PasswordResetBody{}
		if err := json.NewDecoder(r.Body).Decode(resetBody); err != nil {
			util.LogErrorAndSendBadRequest(w,r,err)
		} else {
			validate := validator.New()
			if err = validate.Struct(resetBody) ; err != nil {
				util.LogErrorAndSendBadRequest(w,r,err)
			} else {
				ft := tok.ForgotToken
				if ft != resetBody.ForgotToken {
					util.LogErrorAndSendBadRequest(w,r,errors.New("forgot Token mismatch"))
				} else {
					if err := repository.ResetPassword(resetBody.Email,resetBody.Password); err != nil {
						util.LogErrorAndSendBadRequest(w,r,err)
					} else {
						w.WriteHeader(http.StatusNoContent)}}
			}
		}
	}
}
