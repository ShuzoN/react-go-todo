package di

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"headphonista/src/services"

	"go.uber.org/dig"
)

var Container = Init()

func Init() *dig.Container {
	c := dig.New()

	c.Provide(bootstrap.GetDB)
	c.Provide(infrastructures.CreateUserRepositoryOnMysql)
	c.Provide(services.CreateUserService)

	return c
}
