package services

import (
	"github.com/unicdev/go-shop/internal/dto"
	"github.com/unicdev/go-shop/internal/models"
	"gorm.io/gorm"
)

// UserServiceInterface is the interface for the user service
var _ UserServiceInterface = (*UserService)(nil)

// UserService is the service for the user
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new user service
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetProfile gets a user profile
func (s *UserService) GetProfile(userID uint) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Role:      string(user.Role),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// UpdateProfile updates a user profile
func (s *UserService) UpdateProfile(userID uint, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return s.GetProfile(userID)
}
