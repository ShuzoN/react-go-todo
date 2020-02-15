package services

import (
	"headphonista/src/dto"
	"log"
)

type TodoService struct {
	TodoRepository TodoRepository
}

func (s *TodoService) GetAll() (*[]dto.Todo, error) {
	todos, err := s.TodoRepository.GetAll()

	if err != nil {
		log.Println("unable to execute search query", err)
		return todos, err
	}

	return todos, nil

}

func (s *TodoService) GetById(id int) (*dto.Todo, error) {
	todo, err := s.TodoRepository.GetByID(id)

	if err != nil {
		log.Println("unable to execute search query", err)
		return todo, err
	}

	return todo, nil

}
