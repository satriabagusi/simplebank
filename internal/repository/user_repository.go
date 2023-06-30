/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_repository.go (c) 2023
Desc: description
Created:  2023-06-29T17:15:40.089Z
Modified: !date!
*/

package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/internal/entity/dto/response"
)

type UserRepository interface {
	FindUserByUsernameLogin(string) (*entity.User, error)
	FindByUsername(string) (*response.User, error)
}

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) FindByUsername(username string) (*response.User, error) {
	var user entity.Users

	getJson, err := os.Open("data/user.json")
	if err != nil {
		return &response.User{}, nil
	}

	defer getJson.Close()

	byteVal, _ := ioutil.ReadAll(getJson)

	json.Unmarshal(byteVal, &user)

	res := response.User{}

	for i := 0; i < len(user.User); i++ {
		if user.User[i].Username == username {
			res.Username = user.User[i].Username
			res.Email = user.User[i].Email
		}
	}

	return &res, nil
}

func (r *userRepository) FindUserByUsernameLogin(username string) (*entity.User, error) {
	var user entity.Users

	getJson, err := os.Open("data/user.json")
	if err != nil {
		return &entity.User{}, nil
	}

	defer getJson.Close()

	byteVal, _ := ioutil.ReadAll(getJson)

	json.Unmarshal(byteVal, &user)

	res := entity.User{}

	for i := 0; i < len(user.User); i++ {
		if user.User[i].Username == username {
			res.Username = user.User[i].Username
			res.Password = user.User[i].Password
			res.Email = user.User[i].Email
		}
	}

	return &res, nil

}
