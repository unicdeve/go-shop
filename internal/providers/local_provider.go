package providers

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
)

// LocalUploadProvider is the provider for the local upload
type LocalUploadProvider struct {
	basePath string
}

// NewLocalUploadProvider creates a new local upload provider
func NewLocalUploadProvider(basePath string) *LocalUploadProvider {
	return &LocalUploadProvider{basePath: basePath}
}

// UploadFile uploads a file to the local filesystem
func (p *LocalUploadProvider) UploadFile(file *multipart.FileHeader, path string) (string, error) {
	fullPath := filepath.Join(p.basePath, path)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return "", err
	}

	// Open source
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func() {
		if err := src.Close(); err != nil {
			fmt.Printf("Failed to close source: %v", err)
		}
	}()

	// create destination
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := dst.Close(); err != nil {
			fmt.Printf("Failed to close destination: %v", err)
		}
	}()

	// read from source to destination
	if _, err := dst.ReadFrom(src); err != nil {
		return "", err
	}

	return fmt.Sprintf("/uploads/%s", path), nil

}

// DeleteFile deletes a file from the local filesystem
func (p *LocalUploadProvider) DeleteFile(path string) error {
	fullPath := filepath.Join(p.basePath, path)
	return os.Remove(fullPath)
}
