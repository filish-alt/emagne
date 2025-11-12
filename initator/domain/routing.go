package domain

import (
	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/handler/middleware"
	"github.com/filagot/emagne/internal/routes/auth"
	"github.com/gin-gonic/gin"
)

// InitiateRouting sets up all application routes
func InitiateRouting(
    group *gin.RouterGroup,
    router *gin.Engine,
    handler *Handler, 
    cfg *config.Config) {
    authMw := middleware.InitAuthMiddleware(cfg)
    apiGroup := router.Group("/api")
    auth.InitRoutes(apiGroup,handler.AuthHandler, authMw)
    // authGroup := router.Group("/auth")
    // {
    //     authGroup.POST("/register", handler.AuthHandler.Register)
    //     authGroup.POST("/login", handler.AuthHandler.Login)
    // }

    // apiGroup := router.Group("/api")
    // apiGroup.Use(authMw.Authorize())
    // {
    //     apiGroup.GET("/profile", handler.AuthHandler.GetProfile)
    // }
}
