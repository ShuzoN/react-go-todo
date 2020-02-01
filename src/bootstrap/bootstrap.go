package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Pool = boot()

type bootstrap struct {
	db *gorm.DB
}

func (bs *bootstrap) GetDB() *gorm.DB {
	return bs.db
}

func boot() *bootstrap {
	db, err := gorm.Open("mysql", "root:mysqlrootpassword@tcp(mysqld:3306)/headphonista")
	if err != nil {
		panic(err)
	}

	return &bootstrap{
		db: db,
	}
}
