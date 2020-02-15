package infrastructures

import (
	"headphonista/src/dto"

	"github.com/jinzhu/gorm"
)

type TodoRepositoryOnMysql struct {
	DbConnection *gorm.DB
}

func (todoRepository *TodoRepositoryOnMysql) GetAll() (*[]dto.Todo, error) {
	todos := []dto.Todo{}
	if err := todoRepository.DbConnection.Find(&todos).Error; err != nil {
		return &todos, err
	}
	return &todos, nil
}

func (todoRepository *TodoRepositoryOnMysql) GetByID(id int) (*dto.Todo, error) {
	todo := dto.Todo{}
	if err := todoRepository.DbConnection.First(&todo, id).Error; err != nil {
		return &todo, err
	}
	return &todo, nil
}
