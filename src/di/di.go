package di

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"headphonista/src/services"

	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

var Container = Init()

func Init() *dig.Container {
	c := dig.New()

	c.Provide(bootstrap.GetDB)
	c.Provide(CreateUserRepositoryOnMysql)
	c.Provide(CreateUserService)
	c.Provide(CreateTodoRepositoryOnMysql)
	c.Provide(CreateTodoService)

	return c
}

func CreateUserService() *services.UserService {
	return &services.UserService{
		UserRepository: CreateUserRepositoryOnMysql(bootstrap.GetDB()),
	}
}

func CreateUserRepositoryOnMysql(connection *gorm.DB) services.UserRepository {
	return &infrastructures.UserRepositoryOnMysql{
		DbConnection: connection,
	}
}

func CreateTodoService() *services.TodoService {
	return &services.TodoService{
		TodoRepository: CreateTodoRepositoryOnMysql(bootstrap.GetDB()),
	}
}

func CreateTodoRepositoryOnMysql(connection *gorm.DB) services.TodoRepository {
	return &infrastructures.TodoRepositoryOnMysql{
		DbConnection: connection,
	}
}
