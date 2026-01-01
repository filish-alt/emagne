package transaction

import (
	"github.com/filagot/emagne/internal/handler/middleware"
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(grp *gin.RouterGroup, handler rest.TransactionHandler, authMw middleware.AuthMiddleware) {
	rs := []routes.Router{
		{Method: "POST", Path: "/transactions", Handler: handler.CreateTransaction, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
		{Method: "GET", Path: "/transactions/:id", Handler: handler.GetTransaction, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
		{Method: "GET", Path: "/transactions/categories/:categoryID", Handler: handler.ListTransactionsByCategory, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
		{Method: "PATCH", Path: "/transactions/:id/status", Handler: handler.UpdateTransactionStatus, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},

		{Method: "POST", Path: "/transactions/:id/attributes", Handler: handler.AddTransactionAttribute, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
		{Method: "GET", Path: "/transactions/:id/attributes", Handler: handler.ListTransactionAttributes, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
		{Method: "GET", Path: "/transactions/:id/attributes/:attrID", Handler: handler.GetTransactionAttribute, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
		{Method: "DELETE", Path: "/transactions/:id/attributes", Handler: handler.DeleteTransactionAttributes, Middlewares: []gin.HandlerFunc{authMw.Authorize()}},
	}
	routes.RegisterRoute(grp, rs)
}
