package Func

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
)

func VerifyToken(strToken string) (*model.JWTClaims, err) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New("wrong")
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("wrong")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("wrong")
	}
	fmt.Println("verify")
	return claims, nil
}
