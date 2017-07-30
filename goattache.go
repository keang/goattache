package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/keang/goattache/store"
	"github.com/keang/goattache/utils"
)

type Goattache struct {
	Store store.Store
}

func (g Goattache) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 << 24) // 32 MB
	relPath := generatePath(r.Form.Get("file"))
	for g.Store.Exists(relPath) {
		relPath = generatePath(r.Form.Get("file"))
	}
	saved, err := g.Store.Save(r.Body, relPath)
	if err != nil {
		log.Printf("%v: %v", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonStr, err := json.Marshal(saved)
	if err != nil {
		log.Printf("%v: %v", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonStr))
}

func generatePath(name string) string {
	paths := []string{}
	str := utils.RandString(32)
	for i := 0; i < len(str)-2; i += 2 {
		paths = append(paths, str[i:i+2])
	}
	paths = append(paths, name)
	return filepath.Join(paths...)
}
