package bootstrap

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Pool = boot()

type bootstrap struct {
	db *sql.DB
}

func (bs *bootstrap) GetDB() *sql.DB {
	return bs.db
}

func boot() *bootstrap {
	db, err := sql.Open("mysql", "root:mysqlrootpassword@tcp(mysqld:3306)/headphonista")
	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &bootstrap{
		db: db,
	}
}
