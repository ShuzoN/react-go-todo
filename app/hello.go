package main

import (
	"log"
	"net/http"
)


func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

