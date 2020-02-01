package main

import (
	"headphonista/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `uri:"id" binding:"required"`
}

func main() {
	r := gin.Default()

	ds := services.CreateUserService()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
			return
		}

		name, err := ds.GetUserName(user.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})
	r.Run(":80")

}
