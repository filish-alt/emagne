package rest

import (
	"net/http"

	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/filagot/emagne/internal/module"
	"github.com/gin-gonic/gin"
)

// AuthHandler defines the interface for authentication HTTP handlers
type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
}

// Handler represents the HTTP handler
type Handler struct {
	authModule module.AuthModule
}

// NewHandler creates a new handler instance
func NewHandler(authModule module.AuthModule) *Handler {
	return &Handler{
		authModule: authModule,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password, first name, last name and optional phone
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration request"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authModule.Register(c.Request.Context(), &req)
	if err != nil {
		if err.Error() == "user with this email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary Login user
// @Description Login user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authModule.Login(c.Request.Context(), &req)
	if err != nil {
		if err.Error() == "invalid email or password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the current user's profile information
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	userEmail := c.GetString("user_email")
	userRole := c.GetString("user_role")
	
	c.JSON(http.StatusOK, gin.H{
		"user_id":    userID,
		"user_email": userEmail,
		"user_role":  userRole,
	})
}

