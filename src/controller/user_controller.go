package controller

import (
	"headphonista/src/di"
	"headphonista/src/dto"
	"headphonista/src/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type UserQueryParams struct {
	ID int `uri:"id" binding:"required"`
}

// :id
func (c *UserController) GetUserById(ctx *gin.Context) {
	var queryParams UserQueryParams
	if err := ctx.ShouldBindUri(&queryParams); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	var user *dto.User
	var errGetUser error
	err := di.Container.Invoke(func(ds *services.UserService) {
		user, errGetUser = ds.GetUserById(queryParams.ID)
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		log.Println(err)
		return
	}
	if errGetUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}
