package lib

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
)

type FS interface {
	Write(name string, content []byte) error
	Read(name string) ([]byte, error)
	FullPath(name string) string
}

type LocalFS struct {
	basePath string
}

func NewLocalFS(basePath string) *LocalFS {
	return &LocalFS{basePath: basePath}
}

func (this *LocalFS) Read(name string) ([]byte, error) {
	return ioutil.ReadFile(this.FullPath(name))
}

func (this *LocalFS) Write(name string, content []byte) error {
	return ioutil.WriteFile(this.FullPath(name), content, 0644)
}

func (this *LocalFS) FullPath(name string) string {
	safeFileName := SafeFileName(name)
	path, err := filepath.Abs(path.Join(this.basePath, safeFileName))
	if err != nil {
		log.Fatal(err)
	}
	return path
}
