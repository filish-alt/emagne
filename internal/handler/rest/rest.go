package rest

import (
	"github.com/gin-gonic/gin"
)

// AuthHandler defines the interface for authentication HTTP handlers
type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
}

