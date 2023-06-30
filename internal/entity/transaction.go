/*
Author: Satria Bagus(satria.bagus18@gmail.com)
transaction.go (c) 2023
Desc: description
Created:  2023-06-28T13:44:08.689Z
Modified: !date!
*/

package entity

type Transaction struct {
	ID           string  `json:"id"`
	SenderId     string  `json:"sender_id"`
	ReceiverId   string  `json:"receiver_id"`
	MerchantName string  `json:"merchant_name"`
	Amount       float64 `json:"amount"`
	Timestamp    string  `json:"timestamp"`
}
