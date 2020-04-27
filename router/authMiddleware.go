//JWT Verification Middleware
package router

import (
	"context"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

// JwtVerify Middleware function
func JwtVerify(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Grab the token from the header
		var header,err = jwtmiddleware.FromAuthHeader(r)
		if err != nil {
			w.WriteHeader(401)
			return
		}
		header = strings.TrimSpace(header)
		if header == "" {
			//Token is missing, returns with error code 401
			w.WriteHeader(401)
			return
		}

		_, err = jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(401)
			return
		}
		//Store JWT Token in the Request Context
		ctx := context.WithValue(r.Context(), "student", header)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
