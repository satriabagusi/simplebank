/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_router.go (c) 2023
Desc: description
Created:  2023-06-29T17:12:53.303Z
Modified: !date!
*/

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/simplebank/internal/delivery/handler"
	"github.com/satriabagusi/simplebank/internal/usecase"
	"github.com/satriabagusi/simplebank/middleware"
)

type UserRouter struct {
	userHandler handler.UserHandler
	publicRoute *gin.RouterGroup
}

func (u *UserRouter) SetupRouter() {
	u.publicRoute.POST("/login", u.userHandler.Login)

	u.publicRoute.Use(middleware.AuthMiddleware())
	u.publicRoute.POST("/logout", u.userHandler.Logout)
}

func NewUserRouter(publicRoute *gin.RouterGroup, userUsecase usecase.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)
	rt := UserRouter{
		userHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
