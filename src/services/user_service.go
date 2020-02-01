package services

import (
	"headphonista/src/infrastructures"
	"log"
)

// UserService はユーザのドメインロジックを持つ
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