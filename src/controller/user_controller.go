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
	err := di.Container.Invoke(func(ds *services.UserService) {
		var err error
		user, err = ds.GetUserById(queryParams.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			log.Println(err)
			return
		}
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}
