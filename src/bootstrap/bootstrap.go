package bootstrap

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var pool *bootstrap
var once sync.Once

type bootstrap struct {
	db *gorm.DB
}

func GetDB() *gorm.DB {
	once.Do(func() {
		pool = boot()
	})
	return pool.db
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

func CloseDB() {
	if err := pool.db.Close(); err != nil {
		panic(err)
	}
}
