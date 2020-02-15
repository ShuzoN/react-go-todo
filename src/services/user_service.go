package services

import (
	"headphonista/src/dto"
	"log"
)

type UserService struct {
	UserRepository UserRepository
}

func (ds *UserService) GetById(id int) (*dto.User, error) {
	user, err := ds.UserRepository.GetByID(id)

	if err != nil {
		log.Println("unable to execute search query", err)
		return user, err
	}

	return user, nil

}
