package main

import (
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	// db, err := sql.Open("mysql", "root:mysqlrootpassword@/headphonista")

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
