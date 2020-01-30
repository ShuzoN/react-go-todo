package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	ds := databaseService()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ds.getUserName(1),
		})
	})
	r.Run(":80")

}

type DatabaseService struct {
	databaseRepository *DatabaseRepository
}

func databaseService() *DatabaseService {
	ds := new(DatabaseService)
	ds.databaseRepository = databaseRepository()
	return ds
}

func (ds *DatabaseService) getUserName(id int) string {
	var name string
	err := ds.databaseRepository.getUserById(id).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	return name
}

type DatabaseRepository struct {
	db *sql.DB
}

func databaseRepository() *DatabaseRepository {
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

	databaseRepository := new(DatabaseRepository)
	databaseRepository.db = db

	return databaseRepository
}

func (databaseRepository *DatabaseRepository) getUserById(id int) *sql.Row {
	return databaseRepository.db.QueryRow("select p.name from users as p where p.id = ?;", id)
}
