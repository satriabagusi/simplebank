/*
Author: Satria Bagus(satria.bagus18@gmail.com)
transaction_usecase.go (c) 2023
Desc: description
Created:  2023-06-30T11:41:39.828Z
Modified: !date!
*/

package usecase

import (
	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/internal/repository"
)

type TransactionUsecase interface {
	TransferMoney(string, string, float64) (*entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{transactionRepository}
}

func (u *transactionUsecase) TransferMoney(senderId string, recipientId string, amount float64) (*entity.Transaction, error) {
	return u.transactionRepository.TransferMoney(senderId, recipientId, amount)
}
