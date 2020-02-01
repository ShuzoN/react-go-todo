package main

import (
	"database/sql"
	"headphonista/src/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	boot := bootstrap.ConnectDatabase()

	userRepository := UserRepository{bootstrap: boot}

	ds := UserService{repository: &userRepository}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ds.getUserName(1),
		})
	})
	r.Run(":80")

}

type UserService struct {
	repository *UserRepository
}

func (ds *UserService) getUserName(id int) string {
	var name string
	err := ds.repository.getUserById(id).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	return name
}

type UserRepository struct {
	bootstrap *bootstrap.Bootstrap
}

func (userRepository *UserRepository) getUserById(id int) *sql.Row {
	return userRepository.bootstrap.DbConnection.QueryRow("select p.name from users as p where p.id = ?;", id)
}
