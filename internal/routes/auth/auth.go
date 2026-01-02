package auth

import (
	"github.com/filagot/emagne/internal/handler/middleware"
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/routes"
	"github.com/gin-gonic/gin"
)

// InitRoutes registers auth routes on provided group
func InitRoutes(
     grp *gin.RouterGroup,
     handler rest.AuthHandler,
     authMiddleware middleware.AuthMiddleware) {
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
          {
            Method:  "GET",
            Path:    "/profile",
            Handler: handler.GetProfile,
            Middlewares: []gin.HandlerFunc{
                authMiddleware.Authorize(),
            },
        },
          {
            Method:  "PUT",
            Path:    "/profile",
            Handler: handler.UpdateProfile,
            Middlewares: []gin.HandlerFunc{
                authMiddleware.Authorize(),
                middleware.UserIDFromHeader(),
            },
        },
          {
            Method:  "DELETE",
            Path:    "/users/:id",
            Handler: handler.DeleteUser,
            Middlewares: []gin.HandlerFunc{
                authMiddleware.Authorize(),
            },
        },
    }
    routes.RegisterRoute(grp, authRoutes)
}