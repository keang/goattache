package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/keang/goattache/utils"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 << 24) // 32 MB

	if !authorized(r) {
		w.Header().Set("X-Exception", "Authorization failed")
		http.Error(w, "Authorization failed", 401)
		log.Printf("%+v", 401)
		return
	}

	fmt.Fprintf(w, "{}")
}

// TODO: add docs
func authorized(r *http.Request) bool {
	expectedHMAC := utils.SignHMAC(os.Getenv("ATTACHE_SECRET_KEY"),
		r.Form.Get("uuid")+r.Form.Get("expiration"),
	)
	return expectedHMAC == r.Form.Get("hmac")
}
