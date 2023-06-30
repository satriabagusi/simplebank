/*
Author: Satria Bagus(satria.bagus18@gmail.com)
response.go (c) 2023
Desc: description
Created:  2023-06-29T16:58:05.135Z
Modified: !date!
*/

package response

import "time"

type ResponseMessage struct {
	Code   int    `json:"code"`
	Mesage string `json:"status"`
	Data   any    `json:"data"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type TransactionHistory struct {
	TransactionId        string    `json:"transaction_id"`
	SenderId             int       `json:"sender_id"`
	RecipientId          int       `json:"recipient_id"`
	TimestampTransaction time.Time `json:"timestamp_transaction"`
}

type ResponseMessageWithoutData struct {
	Code   int    `json:"code"`
	Mesage string `json:"status"`
}
