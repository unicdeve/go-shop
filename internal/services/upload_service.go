package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/unicdev/go-shop/internal/interfaces"
)

// UploadServiceInterface is the interface for the upload service
var _ UploadServiceInterface = (*UploadService)(nil)

// UploadService is the service for the upload
type UploadService struct {
	provider interfaces.UploadProvider
}

// NewUploadService creates a new upload service
func NewUploadService(provider interfaces.UploadProvider) *UploadService {
	return &UploadService{provider: provider}
}

// UploadProductImage uploads a product image
func (s *UploadService) UploadProductImage(productID uint, file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImageExt(ext) {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	path := fmt.Sprintf("products/%d/%s%s", productID, uuid.New().String(), ext)

	return s.provider.UploadFile(file, path)
}

// isValidImageExt checks if the file extension is valid
func isValidImageExt(ext string) bool {
	validExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	return slices.Contains(validExts, ext)
}
