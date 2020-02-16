package services

import "headphonista/src/dto"

type TodoRepository interface {
	GetByID(id int) (dto.Todo, error)
	GetAll() ([]dto.Todo, error)
	Create(todo dto.Todo) error
	Update(todo dto.Todo) error
}
