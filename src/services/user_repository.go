package services

import "headphonista/src/dto"

type UserRepository interface {
	GetByID(id int) (*dto.User, error)
}
