package services

import (
	"headphonista/src/bootstrap"
	"headphonista/src/dto"
	"headphonista/src/infrastructures"
	"log"
)

type userService struct {
	userRepository infrastructures.Repository
}

func (ds *userService) GetUserById(id int) (*dto.User, error) {
	var user dto.User
	err := ds.userRepository.GetByID(id).Scan(&user.ID, &user.Name)

	if err != nil {
		log.Println("unable to execute search query", err)
	}

	return &user, err
}

func CreateUserService() *userService {
	return &userService{
		userRepository: infrastructures.CreateUserRepositoryOnMysql(bootstrap.Pool.GetDB()),
	}
}
