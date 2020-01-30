package main

import (
	"bytes"
	"context"
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

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
		tmpl := template.Must(template.ParseFiles("./static/index.tpl"))
		buff := new(bytes.Buffer)
		wbf := io.Writer(buff)
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		var name string
		err := db.QueryRowContext(ctx, "select p.name from users as p where p.id = ?;", 1).Scan(&name)
		if err != nil {
			log.Fatal("unable to execute search query", err)
		}
		log.Println("name=", name)

		data := struct {
			Name string
		}{
			Name: name,
		}

		err = tmpl.ExecuteTemplate(wbf, "index", data)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(buff.Bytes())
		return
	}
}
