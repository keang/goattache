package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/keang/goattache/middlewares"
	"github.com/keang/goattache/store"
)

func main() {
	portFlag := flag.String("port", "9292", "port number to bind to")
	secretFlag := flag.String("secret", "", "secret key for hmac signature")
	dataDirFlag := flag.String("data_dir", "public", "where your files is stored on local disk")
	flag.Parse()

	dataDir := os.Getenv("ATTACHE_LOCAL_DIR")
	if dataDir == "" {
		dataDir = *dataDirFlag
	}
	secret := os.Getenv("ATTACHE_SECRET_KEY")
	if secret == "" {
		secret = *secretFlag
	}

	g := Goattache{Store: store.Disk{dataDir}}
	uploadHandler := http.HandlerFunc(g.UploadHandler)
	http.Handle("/upload", middlewares.Authorize(secret, uploadHandler))
	// trailling slash /view/ to match subtrees
	http.Handle("/view/", http.HandlerFunc(g.DownloadHandler))
	log.Printf("Listening to %v", *portFlag)
	log.Fatal(http.ListenAndServe(":"+*portFlag, nil))
}
