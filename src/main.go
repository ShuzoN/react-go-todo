package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	dbConnection := MysqlConnect()

	userRepository := UserRepository{connection: dbConnection}

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
	connection *MysqlConnection
}

func (userRepository *UserRepository) getUserById(id int) *sql.Row {
	return userRepository.connection.db.QueryRow("select p.name from users as p where p.id = ?;", id)
}

type MysqlConnection struct {
	db *sql.DB
}

func MysqlConnect() *MysqlConnection {
	db, err := sql.Open("mysql", "root:mysqlrootpassword@tcp(mysqld:3306)/headphonista")
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	c := new(MysqlConnection)
	c.db = db

	return c
}
