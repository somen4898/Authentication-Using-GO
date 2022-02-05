package jwtHelper

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-auth/models/user"
	"time"
)

var mySignedKey = []byte("mysupersecretkey")

func Sign(u user.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = u
	claims["expire"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func Parse(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return mySignedKey, nil
	})
}
