package models

import (
	"time"

	"gorm.io/gorm"
)

// Order is the model for the order
type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	Status      OrderStatus    `json:"status" gorm:"default:pending"`
	TotalAmount float64        `json:"total_amount" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items"`
}

// OrderStatus is the status of the order
type OrderStatus string

// OrderStatus enums
const (
	OrderStatusPending   OrderStatus = "pending"   // The order is pending
	OrderStatusConfirmed OrderStatus = "confirmed" // The order is confirmed
	OrderStatusShipped   OrderStatus = "shipped"   // The order is shipped
	OrderStatusDelivered OrderStatus = "delivered" // The order is delivered
	OrderStatusCancelled OrderStatus = "cancelled" // The order is cancelled
)

// OrderItem is the model for the order item
type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	OrderID   uint           `json:"order_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Order   Order   `json:"-"`
	Product Product `json:"product"`
}

// Cart is the model for the cart
type Cart struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	CartItems []CartItem `json:"cart_items"`
}

// CartItem is the model for the cart item
type CartItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CartID    uint           `json:"cart_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Cart    Cart    `json:"-"`
	Product Product `json:"product"`
}
