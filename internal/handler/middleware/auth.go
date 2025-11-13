package middleware

import (
	"net/http"
	"strings"

	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Authorize() gin.HandlerFunc

}

type authMiddleware struct{
	cfg *config.Config
}	

func InitAuthMiddleware(cfg *config.Config) AuthMiddleware{
	return &authMiddleware{
		cfg: cfg,
	}
}
func (a *authMiddleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token
		claims, err := utils.VerifyToken(tokenString, a.cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in context as strings for easy retrieval
		SetUserContext(c, claims.UserID.String(), claims.Email, claims.Role)

		c.Next()
	}
}
