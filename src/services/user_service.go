package services

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"log"
)

type UserService struct {
	userRepository infrastructures.Repository
}

func (ds *UserService) GetUserName(id int) string {
	var name string
	err := ds.userRepository.GetByID(id).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	return name
}

func CreateUserService() *UserService {
	return &UserService{
		userRepository: infrastructures.CreateUserRepositoryOnMysql(bootstrap.Connection.DbConnection),
	}
}
