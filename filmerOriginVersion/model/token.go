package model

import (
	"github.com/dgrijalva/jwt-go"
)

//token的结构体
type JWTClaims struct {
	jwt.StandardClaims
	UserID string
}

var Secret = "filmer"
