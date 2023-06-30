/*
Author: Satria Bagus(satria.bagus18@gmail.com)
transaction_router.go (c) 2023
Desc: description
Created:  2023-06-30T12:11:15.779Z
Modified: !date!
*/

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/simplebank/internal/delivery/handler"
	"github.com/satriabagusi/simplebank/internal/usecase"
	"github.com/satriabagusi/simplebank/middleware"
)

type TransactionRouter struct {
	transactionHander handler.TransactionHander
	publicRoute       *gin.RouterGroup
}

func (t *TransactionRouter) SetupRouter() {
	t.publicRoute.Use(middleware.AuthMiddleware())
	t.publicRoute.POST("transaction/create", t.transactionHander.TransferMoney)
}

func NewTransactionRouter(publicRoute *gin.RouterGroup, transactionUsecase usecase.TransactionUsecase, userUsecase usecase.UserUsecase) {
	transactionHandler := handler.NewTransactionHandler(transactionUsecase, userUsecase)
	rt := TransactionRouter{
		transactionHandler,
		publicRoute,
	}

	rt.SetupRouter()
}
