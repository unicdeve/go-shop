package resolver

import (
	"context"
	"errors"

	"github.com/unicdev/go-shop/internal/models"
	"github.com/unicdev/go-shop/internal/utils"
)

var (
	// ErrUnauthorized is the error for unauthorized access
	ErrUnauthorized = errors.New("unauthorized")
)

const (
	// adminRole is the role for the admin
	adminRole = "admin"
)

// GetUserIDFromContext functions to extract user info from GraphQL context
func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID := ctx.Value(utils.UserIDKey)
	if userID == nil {
		return 0, ErrUnauthorized
	}

	if id, ok := userID.(uint); ok {
		return id, nil
	}

	return 0, ErrUnauthorized
}

// GetUserRoleFromContext functions to extract user role from GraphQL context
func GetUserRoleFromContext(ctx context.Context) (string, error) {
	userRole := ctx.Value(utils.UserRoleKey)
	if userRole == nil {
		return "", ErrUnauthorized
	}

	if role, ok := userRole.(models.UserRole); ok {
		return string(role), nil
	}

	return "", ErrUnauthorized
}

// IsAdminFromContext functions to check if the user is an admin
func IsAdminFromContext(ctx context.Context) bool {
	role, err := GetUserRoleFromContext(ctx)
	if err != nil {
		return false
	}

	return role == adminRole
}

// GetPageAndLimit functions to get the page and limit from the context
func GetPageAndLimit(page, limit *int) (pageNum, limitNum int) {
	if page != nil {
		pageNum = *page
	}

	if limit != nil {
		limitNum = *limit
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	if limitNum <= 0 {
		limitNum = 10
	}

	return pageNum, limitNum
}
