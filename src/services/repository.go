package services

import "headphonista/src/dto"

type Repository interface {
	GetByID(id int) (*dto.User, error)
}
