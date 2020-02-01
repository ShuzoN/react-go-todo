package services

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"log"
)

type UserService struct {
	Repository *infrastructures.UserRepositoryOnMysql
}

func (ds *UserService) GetUserName(id int) string {
	var name string
	err := ds.Repository.GetUserByID(id).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	return name
}

func CreateUserService() *UserService {
	userService := new(UserService)
	userService.Repository = &infrastructures.UserRepositoryOnMysql{Bootstrap: bootstrap.Connection}

	return userService
}
