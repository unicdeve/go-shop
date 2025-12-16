package resolver

import (
	"strconv"

	"github.com/unicdev/go-shop/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver for the graphql schema
type Resolver struct {
	authService    *services.AuthService
	userService    *services.UserService
	productService *services.ProductService
	cartService    *services.CartService
	orderService   *services.OrderService
}

// NewResolver creates a new resolver
func NewResolver(
	authService *services.AuthService,
	userService *services.UserService,
	productService *services.ProductService,
	cartService *services.CartService,
	orderService *services.OrderService,
) *Resolver {
	return &Resolver{
		authService:    authService,
		userService:    userService,
		productService: productService,
		cartService:    cartService,
		orderService:   orderService,
	}

}

// parseID parses a string to a uint
func (r *Resolver) parseID(id string) (uint, error) {
	parsed, err := strconv.ParseUint(id, 10, 32)
	return uint(parsed), err
}
