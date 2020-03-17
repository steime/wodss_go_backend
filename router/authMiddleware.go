package router

import (
	"context"
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"
	"strings"
)

//Exception struct
type Exception persistence.Exception

// JwtVerify Middleware function
func JwtVerify(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var header,err = jwtmiddleware.FromAuthHeader(r) //Grab the token from the header
		if err != nil {
			panic(err.Error())
		}
		header = strings.TrimSpace(header)
		if header == "" {
			//Token is missing, returns with error code 401
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &persistence.Token{}


		_, err = jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}
		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
