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
