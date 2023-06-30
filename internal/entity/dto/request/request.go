/*
Author: Satria Bagus(satria.bagus18@gmail.com)
request.go (c) 2023
Desc: description
Created:  2023-06-29T16:58:01.197Z
Modified: !date!
*/

package request

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
}

type Payment struct {
	SenderID    string  `json:"sender_id" validate:"required"`
	RecipientID string  `json:"recipient_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}
