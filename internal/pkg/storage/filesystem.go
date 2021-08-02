package storage

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileSystemStorage struct {
	storageDir string
}

func NewFileSystemStorage(storageDir string) *FileSystemStorage {
	fss := new(FileSystemStorage)
	fss.storageDir = storageDir

	return fss
}

func (fss FileSystemStorage) Save(file *multipart.FileHeader, dir string, filename string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(root, fss.storageDir, dir, filename)
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		log.Println(err)
		return "", err
	}

	// Create file
	dstFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	// Copy File
	if _, err = io.Copy(dstFile, src); err != nil {
		return "", err
	}

	// Generate path (local)
	dirToPath := strings.ReplaceAll(dir, "\\", "/")
	paths := []string{"", "storage", dirToPath, filename}
	sotrageFilepath := strings.Join(paths, "/")
	return sotrageFilepath, nil
}
