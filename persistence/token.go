package persistence

import jwt "github.com/dgrijalva/jwt-go"

type Token struct {
	StudentID uint
	Email  string
	*jwt.StandardClaims
}
