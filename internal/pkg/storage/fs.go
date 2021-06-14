package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileSystemStorage struct {
	root string
}

func (fss FileSystemStorage) Save(file multipart.FileHeader, dir string, filename string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	fname := filepath.Join(fss.root, dir, filename)
	dstFile, err := os.Create(fname)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	// Copy
	if _, err = io.Copy(dstFile, src); err != nil {
		return "", err
	}

	return fname, nil
}
