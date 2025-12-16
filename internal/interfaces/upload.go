package interfaces

import "mime/multipart"

// UploadProvider is the interface for the upload provider
type UploadProvider interface {
	UploadFile(file *multipart.FileHeader, path string) (string, error)
	DeleteFile(path string) error
}
