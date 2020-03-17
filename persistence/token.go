package persistence

import "github.com/dgrijalva/jwt-go"

type Token struct {
	StudentID uint
	Email  string
	*jwt.StandardClaims
}
