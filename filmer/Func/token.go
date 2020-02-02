package Func

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/filmer/model"
	"log"
)

func VerifyToken(strToken string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(model.Secret), nil
	})
	log.Println(err)
	if err != nil {
		return nil, errors.New("wrong1")
	}
	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, errors.New("wrong2")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("wrong3")
	}
	fmt.Println("verify")
	return claims, nil
}
