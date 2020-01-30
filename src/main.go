package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	db, _ := sql.Open("mysql", "root:mysqlrootpassword@tcp(172.22.0.3:3306)/headphonista")

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &Service{db: db}

	log.Println("Listening...")
	http.ListenAndServe(":80", s)
}

type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/":
		log.Println("/")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		var name string
		err := db.QueryRowContext(ctx, "select p.name from users as p where p.id = ?;", 1).Scan(&name)
		if err != nil {
			log.Fatal("unable to execute search query", err)
		}
		log.Println("name=", name)
		return
	}
}
