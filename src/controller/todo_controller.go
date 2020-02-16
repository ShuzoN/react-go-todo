package controller

import (
	"headphonista/src/di"
	"headphonista/src/dto"
	"headphonista/src/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
}

type TodoQueryParams struct {
	ID int `uri:"id" binding:"required"`
}

type TodoSerialize struct {
	ID       int       `json:"id" form:"id"`
	Title    string    `json:"title" form:"title"`
	Deadline time.Time `json:"deadline" form:"deadline"`
}

func todosSerializer(todos []dto.Todo) []TodoSerialize {
	todoSerializes := make([]TodoSerialize, len(todos))
	for i, todo := range todos {
		todoSerializes[i] = todoSerializer(todo)
	}

	return todoSerializes
}

func todoSerializer(todo dto.Todo) TodoSerialize {
	todoSerializes := TodoSerialize{}
	todoSerializes.ID = todo.ID
	todoSerializes.Title = todo.Title
	todoSerializes.Deadline = todo.Deadline

	return todoSerializes
}

// get /
func (c *TodoController) All(ctx *gin.Context) {
	var todos []dto.Todo
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

	ctx.JSON(http.StatusOK, todosSerializer(todos))
}

// get :id
func (c *TodoController) GetById(ctx *gin.Context) {
	var queryParams TodoQueryParams
	if err := ctx.ShouldBindUri(&queryParams); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	var todo dto.Todo
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

	ctx.JSON(http.StatusOK, todoSerializer(todo))
}

// post :id
// '{"id": 1, "title":"hoge", "deadline": "2020-01-02T00:00:00+09:00"}'
func (c *TodoController) Post(ctx *gin.Context) {
	var reqTodo TodoSerialize
	if err := ctx.BindJSON(&reqTodo); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}

	todo := dto.Todo{
		ID:       reqTodo.ID,
		Title:    reqTodo.Title,
		Deadline: reqTodo.Deadline,
	}

	var errUpdate error
	err := di.Container.Invoke(func(s *services.TodoService) {
		errUpdate = s.Update(todo)
	})

	if err != nil || errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
	return
}
