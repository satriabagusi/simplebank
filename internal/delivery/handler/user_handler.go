/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_handler.go (c) 2023
Desc: description
Created:  2023-06-29T17:14:03.521Z
Modified: !date!
*/

package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		return
	}

	if err := utils.VerifyPassword(dataUser.Password, login.Password); err != nil {
		log.Println(dataUser.Password, login.Password)
		log.Println("Username: ", login.Username, "Password: ", login.Password)
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Wrong Username or Password",
		})
		return
	}

	tokenString, err := token.CreateToken(dataUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error. Please try again",
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": "Internal server error. Please try again",
		"data":    tokenString,
	})
	return
}

func (u *userHandler) Logout(ctx *gin.Context) {
	authPayload := ctx.MustGet("payload").(*token.MyCustomClaims)

	authPayload.ExpiresAt = &jwt.NumericDate{time.Now()}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Logged out.",
	})
}
