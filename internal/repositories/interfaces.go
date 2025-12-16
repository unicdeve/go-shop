package repositories

import "github.com/unicdev/go-shop/internal/models"

// UserRepositoryInterface is the interface for the user repository
type UserRepositoryInterface interface {
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByEmailAndActive(email string, isActive bool) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error

	CreateRefreshToken(token *models.RefreshToken) error
	GetValidRefreshToken(token string) (*models.RefreshToken, error)
	DeleteRefreshToken(token string) error
	DeleteRefreshTokenByID(id uint) error
}

// CartRepositoryInterface is the interface for the cart repository
type CartRepositoryInterface interface {
	GetByUserID(userID uint) (*models.Cart, error)
	Create(cart *models.Cart) error
	Update(cart *models.Cart) error
	Delete(id uint) error
}
