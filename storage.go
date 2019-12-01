package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

type File struct {
	filename string
	contents []byte
}

type Storage struct {
	Directory string
}

func NewStorage(dir string) (s Storage, err error) {
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}

	s = Storage{
		Directory: dir,
	}

	return
}

func (s Storage) join(filename string) string {
	return filepath.Join(s.Directory, filename)
}

func (s Storage) Save(f File) (err error) {
	return ioutil.WriteFile(s.join(f.filename), f.contents, 0644)
}

func deriveFilename(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}

	return filepath.Base(u.Path)
}
