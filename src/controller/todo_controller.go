package controller

import (
	"headphonista/src/di"
	"headphonista/src/dto"
	"headphonista/src/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
}

type TodoQueryParams struct {
	ID int `uri:"id" binding:"required"`
}

// get /
func (c *TodoController) All(ctx *gin.Context) {
	var todos *[]dto.Todo
	var errGetTodos error
	err := di.Container.Invoke(func(s *services.TodoService) {
		todos, errGetTodos = s.GetAll()
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		log.Println(err)
		return
	}
	if errGetTodos != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

// get :id
func (c *TodoController) GetById(ctx *gin.Context) {
	var queryParams TodoQueryParams
	if err := ctx.ShouldBindUri(&queryParams); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	var todo *dto.Todo
	var errGetTodo error
	err := di.Container.Invoke(func(s *services.TodoService) {
		todo, errGetTodo = s.GetById(queryParams.ID)
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		log.Println(err)
		return
	}
	if errGetTodo != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       todo.ID,
		"title":    todo.Title,
		"deadline": todo.Deadline,
	})
}
