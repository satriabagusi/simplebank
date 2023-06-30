/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_handler.go (c) 2023
Desc: description
Created:  2023-06-29T17:14:03.521Z
Modified: !date!
*/

package handler

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/internal/entity/dto/request"
	"github.com/satriabagusi/simplebank/internal/usecase"
	"github.com/satriabagusi/simplebank/pkg/token"
	"github.com/satriabagusi/simplebank/pkg/utils"
)

type UserHandler interface {
	Login(*gin.Context)
	Logout(*gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}

func (u *userHandler) Login(ctx *gin.Context) {
	var login request.Login
	randInteger := rand.Intn(99999)
	logId := fmt.Sprintf("log-", string(randInteger))

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request. Data required.",
		})
		return
	}

	dataUser, err := u.userUsecase.Login(login.Username)
	if err != nil {
		log.Println("Username: ", login.Username, "Password: ", login.Password)
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Wrong Username or Password",
		})
		fmt.Println("find user failure")
		return
	}

	if err := utils.VerifyPassword(dataUser.Password, login.Password); err != nil {
		log.Println("Username: ", login.Username, "Password: ", login.Password)
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Wrong Username or Password",
		})
		fmt.Println("matches user failure")
		return
	}

	tokenString, err := token.CreateToken(dataUser, 2)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error. Please try again",
		})

		logData := entity.Log{
			ID:         logId,
			LogName:    "Auth",
			LogStatus:  "success",
			LogMessage: "User with ID:" + dataUser.Id + " trying to logged in and was unsuccessful logged in. caused by Internal Server Error",
			Timestamp:  time.Now(),
		}

		putLog, err := utils.PutToLog(logData)
		if putLog && err != nil {
			log.Println("Failed to create log transaction: ", err)
		}

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code":    http.StatusAccepted,
		"message": "Login successfully",
		"data":    tokenString,
	})

	logData := entity.Log{
		ID:         logId,
		LogName:    "Auth",
		LogStatus:  "success",
		LogMessage: "User with ID:" + dataUser.Id + " trying to logged in and was successfully logged in",
		Timestamp:  time.Now(),
	}

	putLog, err := utils.PutToLog(logData)
	if putLog && err != nil {
		log.Println("Failed to create log auth: ", err)
	}

	return
}

func (u *userHandler) Logout(ctx *gin.Context) {
	token, ok := ctx.MustGet("payload").(*jwt.Token)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Invalid Token.",
		})
		return
	}

	claims, ok := token.Claims.(jwt.StandardClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Invalid Token Claims.",
		})
		return
	}

	claims.ExpiresAt = time.Now().Unix() - 1

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Logged out.",
	})
}
