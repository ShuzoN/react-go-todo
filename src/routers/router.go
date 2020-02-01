package routers

import (
	"headphonista/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID int `uri:"id" binding:"required"`
}

func Router(r *gin.Engine) {
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

		ds := services.CreateUserService()
		name, err := ds.GetUserName(user.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})
}
