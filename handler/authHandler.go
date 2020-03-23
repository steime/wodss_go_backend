package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/steime/wodss_go_backend/persistence"
	"github.com/steime/wodss_go_backend/util"
	"log"
	"net/http"
	"time"
)

type LoginBody struct {
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func Login(repository persistence.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		student := &LoginBody{}
		err := json.NewDecoder(r.Body).Decode(student)
		if err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := repository.FindOne(student.Email, student.Password)
			if resp["status"] == false {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				json.NewEncoder(w).Encode(resp)
			}
		}
	}
}

func RefreshToken(repository persistence.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type tokenReqBody struct {
			RefreshToken string `json:"refreshToken"`
		}
		tokenReq := &tokenReqBody{}
		if err := json.NewDecoder(r.Body).Decode(tokenReq); err !=nil {
			log.Fatal(err)
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
				newTokenPair, err := generateTokenPair(student.ID)
				if err != nil {
					log.Fatal(err)
				}
				json.NewEncoder(w).Encode(newTokenPair)
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
	})
}

func generateTokenPair(studentID uint) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = studentID
	//claims["mail"] = mail
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = studentID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token":  t,
		"refreshToken": rt,
	}, nil
}