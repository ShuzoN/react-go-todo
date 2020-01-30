package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	ds := databaseService()

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ds.getUserName(1),
		})
	})
	r.Run(":80")

}

type DatabaseServeice struct {
	db *sql.DB
}

func databaseService() *DatabaseServeice {
	db, _ := sql.Open("mysql", "root:mysqlrootpassword@tcp(172.22.0.3:3306)/headphonista")

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	ds := new(DatabaseServeice)
	ds.db = db

	return ds
}

func (ds *DatabaseServeice) getUserName(id int) string {
	var name string
	err := ds.db.QueryRow("select p.name from users as p where p.id = ?;", id).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	return name
}
