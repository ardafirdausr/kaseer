package internal

import "mime/multipart"

type Storage interface {
	Save(file *multipart.FileHeader, dir string, filename string) (string, error)
}
