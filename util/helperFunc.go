package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CheckID(r *http.Request) (bool,string) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := r.Context()
	tk := ctx.Value("student")
	token, _ := jwt.Parse(tk.(string), func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimId := int(claims["sub"].(float64))
		studId := strconv.Itoa(claimId)
		return studId == id ,studId
	} else {
		return false,"-1"
	}
}

func CheckBodyID(r *http.Request, id uint) bool {
	ctx := r.Context()
	tk := ctx.Value("student")
	token, _ := jwt.Parse(tk.(string), func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimId := uint(claims["sub"].(float64))
		return claimId == id
	} else {
		return false
	}
}
