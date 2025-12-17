package repositories

import (
	"github.com/unicdev/go-shop/internal/models"
	"gorm.io/gorm"
)

// CartRepository is the repository for the cart
type CartRepository struct {
	db *gorm.DB
}

// NewCartRepository creates a new cart repository
func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

// GetByUserID gets a cart by user id
func (r *CartRepository) GetByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("CartItems.Product.Category").
		Where("user_id = ?", userID).First(&cart).Error

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

// Create creates a new cart
func (r *CartRepository) Create(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

// Update updates a cart
func (r *CartRepository) Update(cart *models.Cart) error {
	return r.db.Save(cart).Error
}

// Delete deletes a cart
func (r *CartRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cart{}, id).Error
}
