package store

import "os"

type SavedFile struct {
	*os.File
	Path        string `json:"path,omitempty"`
	Geometry    string `json:"geometry,omitempty"`
	Size        int64  `json:"bytes,omitempty"`
	ContentType string `json:"content_type,omitempty"`
}
