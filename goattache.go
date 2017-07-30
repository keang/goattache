package main

import (
	"fmt"
	"net/http"
)

type Goattache struct {
}

func (g Goattache) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 << 24) // 32 MB
	//g.Store.save(image)
	fmt.Fprintf(w, `{"url":""}`, "dummy/image/path")
}
