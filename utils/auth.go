package utils

import (
	"errors"
	"fmt"
	"jwt-demo/service"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaim struct {
	Userid string `json:"userid"`
	Gruop  int    `json:"group"`
	jwt.StandardClaims
}

type UserInfo struct {
	Userid string
	Group  int
}

func SignJwt(key []byte, user service.User) string {
	customClaim := CustomClaim{
		user.Userid,
		user.Group,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func ValidJwt(tokenString string, key []byte) (user UserInfo, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		user.Userid = claims.Userid
		user.Group = claims.Gruop
	} else {
		err = errors.New("user authentication failed")
	}
	return
}
