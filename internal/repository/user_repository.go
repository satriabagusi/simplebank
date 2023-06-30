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
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/internal/entity/dto/response"
)

type UserRepository interface {
	FindUserByUsernameLogin(string) (*entity.User, error)
	FindByUsername(string) (*response.User, error)
	FindById(string) (*entity.User, error)
}

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) FindByUsername(username string) (*response.User, error) {
	var user entity.Data

	getJson, err := os.Open("data/user.json")
	if err != nil {
		return &response.User{}, nil
	}

	defer getJson.Close()

	byteVal, _ := ioutil.ReadAll(getJson)

	json.Unmarshal(byteVal, &user)

	res := response.User{}

	for i := 0; i < len(user.Users); i++ {
		if user.Users[i].Username == username {
			res.Username = user.Users[i].Username
			res.Email = user.Users[i].Email
		}
	}

	return &res, nil
}

func (r *userRepository) FindUserByUsernameLogin(username string) (*entity.User, error) {

	getJson, err := os.Open("data/user.json")
	if err != nil {
		log.Println("Error opening json file:", err)
		return nil, err
	}
	defer getJson.Close()

	byteVal, err := io.ReadAll(getJson)
	if err != nil {
		log.Println("Error reading json file:", err)
		return nil, err
	}

	var data struct {
		Users []entity.User `json:"users"`
	}
	if err := json.Unmarshal(byteVal, &data); err != nil {
		log.Println("Error unmarshalling json:", err)
		return nil, err
	}

	log.Println(data)

	res := entity.User{}

	for i := 0; i < len(data.Users); i++ {
		if data.Users[i].Username == username {
			res.Id = data.Users[i].Id
			res.Username = data.Users[i].Username
			res.Email = data.Users[i].Email
			res.Password = data.Users[i].Password
		}
	}

	return &res, nil

}

func (r *userRepository) FindById(id string) (*entity.User, error) {
	var user entity.Data

	getJson, err := os.Open("data/user.json")
	if err != nil {
		return &entity.User{}, nil
	}

	defer getJson.Close()

	byteVal, _ := ioutil.ReadAll(getJson)

	json.Unmarshal(byteVal, &user)

	res := entity.User{}

	for i := 0; i < len(user.Users); i++ {
		if user.Users[i].Id == id {
			res.Id = user.Users[i].Id
			res.Username = user.Users[i].Username
			res.Email = user.Users[i].Email
		}
	}

	return &res, nil
}
