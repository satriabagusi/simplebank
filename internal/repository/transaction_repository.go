/*
Author: Satria Bagus(satria.bagus18@gmail.com)
transaction_repository.go (c) 2023
Desc: description
Created:  2023-06-30T11:13:10.362Z
Modified: !date!
*/

package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/satriabagusi/simplebank/internal/entity"
)

type TransactionRepository interface {
	TransferMoney(string, string, float64) (*entity.Transaction, error)
}

type transactionRepository struct {
}

func NewTransactionRepository() *transactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) TransferMoney(senderId string, recipientId string, amount float64) (*entity.Transaction, error) {
	randNumber := rand.Intn(9999)
	transactionNumber := fmt.Sprintf("trx-", randNumber)

	transaction := entity.Transaction{
		ID:         transactionNumber,
		SenderId:   senderId,
		ReceiverId: recipientId,
		Amount:     amount,
		Timestamp:  time.Now(),
	}

	putTransactionFile, err := json.Marshal(transaction)
	if err != nil {
		log.Println("Could not marshal json: ", err)
		return nil, err
	}

	err = ioutil.WriteFile("data/transaction.json", putTransactionFile, 0644)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
