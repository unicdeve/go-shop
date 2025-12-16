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
	authService    services.AuthServiceInterface
	userService    services.UserServiceInterface
	productService services.ProductServiceInterface
	cartService    services.CartServiceInterface
	orderService   services.OrderServiceInterface
}

// NewResolver creates a new resolver
func NewResolver(
	authService services.AuthServiceInterface,
	userService services.UserServiceInterface,
	productService services.ProductServiceInterface,
	cartService services.CartServiceInterface,
	orderService services.OrderServiceInterface,
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
