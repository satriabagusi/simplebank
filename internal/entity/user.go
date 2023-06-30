/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user.go (c) 2023
Desc: description
Created:  2023-06-28T02:52:05.597Z
Modified: !date!
*/

package entity

type Users struct {
	User []User `json:"user"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
