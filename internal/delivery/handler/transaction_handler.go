/*
Author: Satria Bagus(satria.bagus18@gmail.com)
transaction_handler.go (c) 2023
Desc: description
Created:  2023-06-30T11:45:26.541Z
Modified: !date!
*/

package handler

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/internal/entity/dto/request"
	"github.com/satriabagusi/simplebank/internal/usecase"
	"github.com/satriabagusi/simplebank/pkg/utils"
)

type TransactionHander interface {
	TransferMoney(*gin.Context)
}

type transactionHandler struct {
	transcationUsecase usecase.TransactionUsecase
	userUsecase        usecase.UserUsecase
}

func NewTransactionHandler(transactionUsecase usecase.TransactionUsecase, userUsecase usecase.UserUsecase) TransactionHander {
	return &transactionHandler{
		transactionUsecase,
		userUsecase,
	}
}

func (t *transactionHandler) TransferMoney(ctx *gin.Context) {
	var transaction request.Payment

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request. Check your request parameters",
		})
		return
	}

	_, err := t.userUsecase.FindById(transaction.RecipientID)
	if err != nil {
		log.Println("Sender ID: ", transaction.SenderID, "Recipient ID: ", transaction.RecipientID, "Amount: ", transaction.Amount)
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Failed to create transaction. Check sender ID",
		})
	}

	createTransaction, err := t.transcationUsecase.TransferMoney(transaction.SenderID, transaction.RecipientID, transaction.Amount)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to create transaction causing by Internal Server error.",
		})
		return
	}

	logId := fmt.Sprintf("log-", rand.Intn(99999))
	logData := entity.Log{
		ID:         logId,
		LogName:    "Transaction",
		LogStatus:  "success",
		LogMessage: "User with ID:" + transaction.SenderID + " trying to commit transaction and successful",
		Timestamp:  time.Now(),
	}

	putLog, err := utils.PutToLog(logData)
	if putLog && err != nil {
		log.Println("Failed to create log transaction: ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully create transaction",
		"data":    createTransaction,
	})
}
