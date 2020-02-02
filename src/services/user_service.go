package services

import (
	"headphonista/src/bootstrap"
	"headphonista/src/dto"
	"headphonista/src/infrastructures"
	"log"
)

type UserService struct {
	userRepository infrastructures.Repository
}

func (ds *UserService) GetUserById(id int) (*dto.User, error) {
	user, err := ds.userRepository.GetByID(id)

	if err != nil {
		log.Println("unable to execute search query", err)
		return user, err
	}

	return user, nil

}

func CreateUserService() *UserService {
	return &UserService{
		userRepository: infrastructures.CreateUserRepositoryOnMysql(bootstrap.GetDB()),
	}
}
