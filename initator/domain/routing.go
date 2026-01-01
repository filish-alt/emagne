package domain

import (
	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/handler/middleware"
	authroutes "github.com/filagot/emagne/internal/routes/auth"
	catroutes "github.com/filagot/emagne/internal/routes/category"
	txroutes "github.com/filagot/emagne/internal/routes/transaction"
	"github.com/gin-gonic/gin"
)

// InitiateRouting sets up all application routes
func InitiateRouting(
	group *gin.RouterGroup,
	router *gin.Engine,
	handler *Handler,
	cfg *config.Config) {
	authMw := middleware.InitAuthMiddleware(cfg)
	authroutes.InitRoutes(group, handler.AuthHandler, authMw)
	catroutes.InitRoutes(group, handler.CategoryHandler, authMw)
	txroutes.InitRoutes(group, handler.TransactionHandler, authMw)

}
