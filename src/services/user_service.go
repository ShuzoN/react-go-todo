package services

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"log"
)

type userService struct {
	userRepository infrastructures.Repository
}

func (ds *userService) GetUserName(id int) (string, error) {
	var name string
	err := ds.userRepository.GetByID(id).Scan(&name)
	if err != nil {
		log.Println("unable to execute search query", err)
	}

	return name, err
}

func CreateUserService() *userService {
	return &userService{
		userRepository: infrastructures.CreateUserRepositoryOnMysql(bootstrap.Pool.GetDB()),
	}
}
