package main

import (
	"headphonista/src/bootstrap"
	"headphonista/src/infrastructures"
	"headphonista/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	boot := bootstrap.ConnectDatabase()

	userRepository := infrastructures.UserRepositoryOnMysql{Bootstrap: boot}

	ds := services.UserService{Repository: &userRepository}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ds.GetUserName(1),
		})
	})
	r.Run(":80")

}
