package di

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"headphonista/src/services"

	"go.uber.org/dig"
)

var Container = di.init()

func init() *dig.Container {
	c := dig.New()

	c.Provide(bootstrap.GetDB())
	c.Provide(infrastructures.CreateUserRepositoryOnMysql)
	c.Provide(services.CreateUserService())

	return c
}
