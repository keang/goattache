package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/keang/goattache/handlers"
)

func main() {
	port := flag.String("port", "9292", "port number to bind to")
	flag.Parse()

	http.HandleFunc("/upload", handlers.UploadHandler)
	log.Printf("Listening to %v", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
