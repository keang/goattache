package store

import (
	"io"
)

type Store interface {
	Exists(string) bool
	Save(io.Reader, string) (SavedFile, error)
	Get(string, string) (SavedFile, error)
}
