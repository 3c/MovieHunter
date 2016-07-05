package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("webui: http://127.0.0.1:4321/")

	// static content
	fs := http.FileServer(http.Dir("webui"))
	http.Handle("/", fs)
	http.ListenAndServe(":4321", nil)
}
