package main

import (
	"headphonista/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ds := services.CreateUserService()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ds.GetUserName(1),
		})
	})
	r.Run(":80")

}
