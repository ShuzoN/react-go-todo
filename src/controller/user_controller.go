package controller

import (
	"headphonista/src/services"
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

	ds := services.CreateUserService()
	user, err := ds.GetUserById(queryParams.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})

}
