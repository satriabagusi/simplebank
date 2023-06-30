/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user.go (c) 2023
Desc: description
Created:  2023-06-28T02:52:05.597Z
Modified: !date!
*/

package entity

type Data struct {
	Users []User `json:"user"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
