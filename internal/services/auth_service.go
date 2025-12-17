package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/unicdev/go-shop/internal/config"
	"github.com/unicdev/go-shop/internal/dto"
	"github.com/unicdev/go-shop/internal/events"
	"github.com/unicdev/go-shop/internal/models"
	"github.com/unicdev/go-shop/internal/notifications"
	"github.com/unicdev/go-shop/internal/repositories"
	"github.com/unicdev/go-shop/internal/utils"
)

// AuthServiceInterface is the interface for the auth service
var _ AuthServiceInterface = (*AuthService)(nil)

// AuthService is the service for the auth
type AuthService struct {
	userRepo       repositories.UserRepositoryInterface
	cartRepo       repositories.CartRepositoryInterface
	config         *config.Config
	eventPublisher events.Publisher
}

// NewAuthService creates a new auth service
func NewAuthService(
	config *config.Config,
	userRepo repositories.UserRepositoryInterface,
	cartRepo repositories.CartRepositoryInterface,
	eventPublisher events.Publisher,
) *AuthService {
	return &AuthService{
		config:         config,
		eventPublisher: eventPublisher,
		userRepo:       userRepo,
		cartRepo:       cartRepo,
	}
}

// Register registers a new user
func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if user exists
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("user already exists")
	}
	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      models.UserRoleCustomer,
	}
	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}
	// create a cart
	cart := models.Cart{UserID: user.ID}
	if err := s.cartRepo.Create(&cart); err != nil {
		fmt.Println("Unable to create cart")
		_ = err // ignore error
	}

	// generate token
	return s.generateAuthResponse(&user)

}

// Login logs in a user
func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByEmailAndActive(req.Email, true)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return s.generateAuthResponse(user)
}

// RefreshToken refreshes a token
func (s *AuthService) RefreshToken(req *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	claims, err := utils.ValidateToken(req.RefreshToken, s.config.JWT.Secret)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	refreshToken, err := s.userRepo.GetValidRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("refresh token not found or expired")
	}

	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := s.userRepo.DeleteRefreshToken(refreshToken.Token); err != nil {
		log.Println(err)
		_ = err // ignore error
	}

	return s.generateAuthResponse(user)
}

// Logout logs out a user
func (s *AuthService) Logout(refreshToken string) error {
	return s.userRepo.DeleteRefreshToken(refreshToken)
}

// generateAuthResponse generates an auth response
func (s *AuthService) generateAuthResponse(user *models.User) (*dto.AuthResponse, error) {
	accessToken, refreshToken, err := utils.GenerateTokenPair(
		&s.config.JWT,
		user.ID,
		user.Email,
		string(user.Role),
	)
	if err != nil {
		return nil, err
	}

	refreshTokenModel := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.config.JWT.RefreshExpiresIn),
	}

	if err := s.userRepo.CreateRefreshToken(&refreshTokenModel); err != nil {
		log.Println(err)
		_ = err // ignore error
	}

	err = s.eventPublisher.Publish(notifications.UserLoggedIn, user, map[string]string{
		"user_id": fmt.Sprintf("%d", user.ID),
		"email":   user.Email,
		"role":    string(user.Role),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to publish user logged in event: %w", err)
	}

	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Role:      string(user.Role),
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
