/*
Author: Satria Bagus(satria.bagus18@gmail.com)
token.go (c) 2023
Desc: description
Created:  2023-06-30T03:16:06.845Z
Modified: !date!
*/

package token

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/pkg/utils"
)

var (
	mySigningKey     = []byte(utils.GetEnv("SECRET_KEY"))
	expireTimeInt, _ = strconv.Atoi(utils.GetEnv("TOKEN_EXPIRATION_TIME_IN_MINUTES"))
)

type MyCustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(user *entity.User, duration time.Duration) (string, error) {
	claims := MyCustomClaims{
		user.Id,
		user.Username,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(tokenString string) (any, error) {
	// Remove "Bearer " prefix from tokenString
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		log.Println("Failure to parse token:", err)
		return nil, fmt.Errorf("Unauthorized")
	}

	if !token.Valid {
		log.Println("Token is not valid")
		return nil, fmt.Errorf("Unauthorized")
	}

	return token, nil
}

// func Valid() error {

// }
