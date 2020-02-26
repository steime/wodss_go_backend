package persistence

import jwt "github.com/dgrijalva/jwt-go"

type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}
