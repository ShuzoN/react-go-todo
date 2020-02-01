package bootstrap

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Connection = ConnectDatabase()

type Bootstrap struct {
	DbConnection *sql.DB
}

func ConnectDatabase() *Bootstrap {
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

	bootstrap := new(Bootstrap)
	bootstrap.DbConnection = db

	return bootstrap
}
