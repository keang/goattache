package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/keang/goattache/middlewares"
)

func main() {
	port := flag.String("port", "9292", "port number to bind to")
	secret := flag.String("secret", os.Getenv("ATTACHE_SECRET_KEY"), "secret key for hmac signature")
	flag.Parse()
	g := Goattache{}
	upload := http.HandlerFunc(g.UploadHandler)
	http.Handle("/upload", middlewares.Authorize(*secret, upload))
	log.Printf("Listening to %v", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
