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
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/pkg/utils"
)

var (
	mySigningKey     = []byte(utils.GetEnv("SECRET_KEY"))
	expireTimeInt, _ = strconv.Atoi(utils.GetEnv("TOKEN_EXPIRATION_TIME_IN_MINUTES"))
)

type MyCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *entity.User) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Username,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireTimeInt) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claims, nil
}

func Valid() error {

}
