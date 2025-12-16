package utils

// ContextKey is the key for the context
type ContextKey string

const (
	// UserIDKey is the key for the user id
	UserIDKey ContextKey = "user_id"
	// UserEmailKey is the key for the user email
	UserEmailKey ContextKey = "user_email"
	// UserRoleKey is the key for the user role
	UserRoleKey ContextKey = "user_role"
	// GinContextKey is the key for the gin context
	GinContextKey ContextKey = "gin_context"
)
