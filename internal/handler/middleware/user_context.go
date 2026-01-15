package middleware

import "github.com/gin-gonic/gin"

const (
	userIDContextKey    = "user_id"
	userEmailContextKey = "user_email"
	userRoleContextKey  = "user_role"
)

// SetUserContext stores user information in the Gin context.
func SetUserContext(c *gin.Context, userID, email, role string) {
	if userID != "" {
		c.Set(userIDContextKey, userID)
	}
	if email != "" {
		c.Set(userEmailContextKey, email)
	}
	if role != "" {
		c.Set(userRoleContextKey, role)
	}
}

// GetUserID retrieves the user ID from the Gin context.
func GetUserID(c *gin.Context) string {
	if value, exists := c.Get(userIDContextKey); exists {
		if userID, ok := value.(string); ok {
			return userID
		}
	}
	return ""
}

// GetUserEmail retrieves the user email from the Gin context.
func GetUserEmail(c *gin.Context) string {
	if value, exists := c.Get(userEmailContextKey); exists {
		if email, ok := value.(string); ok {
			return email
		}
	}
	return ""
}

// GetUserRole retrieves the user Role from the Gin context.
func GetUserRole(c *gin.Context) string {
	if value, exists := c.Get(userRoleContextKey); exists {
		if role, ok := value.(string); ok {
			return role
		}
	}
	return ""
}

// UserIDFromHeader sets the user ID on the context if provided through the X-User or X-User-ID header.
func UserIDFromHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetUserID(c) == "" {
			if userID := c.GetHeader("X-User-ID"); userID != "" {
				SetUserContext(c, userID, "", "")
			} else if userID := c.GetHeader("X-User"); userID != "" {
				SetUserContext(c, userID, "", "")
			}
		}
		c.Next()
	}
}
