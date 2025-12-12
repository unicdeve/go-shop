package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unicdev/go-shop/internal/models"
	"github.com/unicdev/go-shop/internal/utils"
)

// AuthMiddleware is the middleware for the authentication
func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(c, "Unauthorized")
			c.Abort()
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			utils.UnauthorizedResponse(c, "Unauthorized")
			c.Abort()
			return
		}

		token := tokenParts[1]
		claims, err := utils.ValidateToken(token, s.config.JWT.Secret)
		if err != nil {
			utils.UnauthorizedResponse(c, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Set("user_email", claims.Email)
		c.Next()
	}
}

// AdminMiddleware is the middleware for the admin
func (s *Server) adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			utils.ForbiddenResponse(c, "Forbidden")
			c.Abort()
			return
		}

		if role != string(models.UserRoleAdmin) {
			utils.ForbiddenResponse(c, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}
