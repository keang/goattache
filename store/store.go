package store

import (
	"io"
)

type Store interface {
	Exists(string) bool
	Save(io.Reader, string) (SavedFile, error)
	Open(string) (SavedFile, error)
}
