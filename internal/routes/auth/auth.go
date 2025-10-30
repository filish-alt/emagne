package auth

import (
    "github.com/filagot/emagne/internal/handler/rest"
    "github.com/filagot/emagne/internal/routes"
    "github.com/gin-gonic/gin"
)

// InitRoutes registers auth routes on provided group
func InitRoutes(grp *gin.RouterGroup, handler rest.AuthHandler) {
    authRoutes := []routes.Router{
        {
            Method:  "POST",
            Path:    "/register",
            Handler: handler.Register,
        },
        {
            Method:  "POST",
            Path:    "/login",
            Handler: handler.Login,
        },
    }
    routes.RegisterRoute(grp, authRoutes)
}