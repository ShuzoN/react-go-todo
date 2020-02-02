package services

import (
	"headphonista/src/dto"
	"log"
)

type UserService struct {
	UserRepository Repository
}

func (ds *UserService) GetUserById(id int) (*dto.User, error) {
	user, err := ds.UserRepository.GetByID(id)

	if err != nil {
		log.Println("unable to execute search query", err)
		return user, err
	}

	return user, nil

}
