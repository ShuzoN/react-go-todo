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
func (c *UserController) GetUser(ctx *gin.Context) {
	var queryParams UserQueryParams
	if err := ctx.ShouldBindUri(&queryParams); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	ds := services.CreateUserService()
	name, err := ds.GetUserName(queryParams.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": name,
	})

}
