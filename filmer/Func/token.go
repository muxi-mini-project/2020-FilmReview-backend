package Func

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/filmer2/model"
	"github.com/filmer2/modelWency"
	"log"
)

func VerifyToken(strToken string) (*modelWency.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &modelWency.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(model.Secret), nil
	})
	log.Println(err)
	if err != nil {
		return nil, errors.New("wrong1")
	}
	claims, ok := token.Claims.(*modelWency.JwtClaims)
	if !ok {
		return nil, errors.New("wrong2")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("wrong3")
	}
	fmt.Println("verify")
	return claims, nil
}
