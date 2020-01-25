package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	db, _ := sql.Open("mysql", "root:mysqlrootpassword@/headphonista")
	err := db.Ping()

	log.Println(err)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
