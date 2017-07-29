package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	port := flag.String("port", "8080", "port number to bind to")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
