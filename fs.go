package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func SaveUploadedFile(r *http.Request, fieldName string, directory string, filename string) (string, error) {
	// Get uploadad file
	uploadedFile, handler, err := r.FormFile("photo")
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename = filename + filepath.Ext(handler.Filename)
	fileLocation := filepath.Join(dir, directory, filename)

	// Open file with Write action or Create File if not exists
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	// create file
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return "", err
	}

	return filename, nil
}
