/*
Author: Satria Bagus(satria.bagus18@gmail.com)
server.go (c) 2023
Desc: description
Created:  2023-06-29T16:53:15.180Z
Modified: !date!
*/

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/simplebank/internal/delivery/router"
	"github.com/satriabagusi/simplebank/internal/repository"
	"github.com/satriabagusi/simplebank/internal/usecase"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Initialize(connstr string) error {

	// validator := validator.New()

	userRepo := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)

	r := gin.Default()
	api := r.Group("api/v1")
	router.NewUserRouter(api, userUsecase)

	s.router = r
	return nil
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
