package store

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
)

type Disk struct {
	Dir string
}

func (d Disk) finalPath(relPath string) string {
	return filepath.Join(d.Dir, relPath)
}

func (d Disk) Exists(relPath string) bool {
	_, err := os.Stat(d.finalPath(relPath))
	return !os.IsNotExist(err)
}

func (d Disk) Save(reader io.Reader, relPath string) (f SavedFile, e error) {
	if d.Exists(relPath) {
		e = errors.New(relPath + " already exists")
		return
	}

	fileName := d.finalPath(relPath)
	os.MkdirAll(filepath.Dir(fileName), 0744)
	file, e := os.Create(fileName)
	if e != nil {
		return
	}
	_, e = io.Copy(file, reader)
	if e != nil {
		return
	}
	file.Close()
	f, e = d.Get(relPath, "")
	defer f.Close()
	return
}

func (d Disk) Get(relPath string, geometry string) (f SavedFile, e error) {
	file, e := os.Open(d.finalPath(relPath))
	if e != nil {
		return
	}
	conf, format, e := image.DecodeConfig(file)
	if e != nil {
		file.Close()
		return
	}
	_, e = file.Seek(0, 0)
	if e != nil {
		file.Close()
		return
	}
	stat, e := file.Stat()
	if e != nil {
		file.Close()
		return
	}

	f = SavedFile{
		File:        file,
		Path:        relPath,
		Geometry:    fmt.Sprintf("%vx%v", conf.Width, conf.Height),
		Size:        stat.Size(),
		ContentType: "image/" + format,
	}

	return
}
