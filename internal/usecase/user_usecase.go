/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_usecase.go (c) 2023
Desc: description
Created:  2023-06-30T02:52:35.823Z
Modified: !date!
*/

package usecase

import (
	"github.com/satriabagusi/simplebank/internal/entity"
	"github.com/satriabagusi/simplebank/internal/entity/dto/response"
	"github.com/satriabagusi/simplebank/internal/repository"
)

type UserUsecase interface {
	FindByUsername(string) (*response.User, error)
	Login(string) (*entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (u *userUsecase) FindByUsername(username string) (*response.User, error) {
	return u.userRepository.FindByUsername(username)
}

func (u *userUsecase) Login(username string) (*entity.User, error) {
	return u.userRepository.FindUserByUsernameLogin(username)
}
