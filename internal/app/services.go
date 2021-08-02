package app

import (
	"path/filepath"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/pkg/storage"
)

type services struct {
	Storage internal.Storage
}

func NewServices() *services {
	storageDir := filepath.Join("web", "storage")
	fileSystemStorage := storage.NewFileSystemStorage(storageDir)

	services := new(services)
	services.Storage = fileSystemStorage
	return services
}
