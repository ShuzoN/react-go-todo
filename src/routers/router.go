package routers

import (
	"net/http"

	"headphonista/src/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	u := r.Group("/users")
	{
		controller := controller.UserController{}
		u.GET("/:id", controller.GetUserById)
	}
	t := r.Group("/todos")
	{
		controller := controller.TodoController{}
		t.GET("/", controller.All)
		t.GET("/:id", controller.GetById)
	}
}
