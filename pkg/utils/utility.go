/*
Author: Satria Bagus(satria.bagus18@gmail.com)
utility.go (c) 2023
Desc: description
Created:  2023-06-29T16:45:55.512Z
Modified: !date!
*/

package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key string, v ...any) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
	}

	if key != "" {
		return os.Getenv(key)
	}

	return v[0].(string)
}

func Encrypt(str string) string {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(encryptedPassword)
}

func VerifyPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
