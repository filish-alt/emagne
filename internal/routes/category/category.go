package category

import (
	"github.com/filagot/emagne/internal/handler/middleware"
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(grp *gin.RouterGroup, handler rest.CategoryHandler, authMw middleware.AuthMiddleware) {
	rs := []routes.Router{
		{
			Method:      "POST",
			Path:        "/categories",
			Handler:     handler.CreateCategory,
			Middlewares: []gin.HandlerFunc{authMw.Authorize(), authMw.RequireRole("super_admin")},
		},
		{
			Method:      "GET",
			Path:        "/categories",
			Handler:     handler.ListCategories,
			Middlewares: []gin.HandlerFunc{authMw.Authorize()},
		},
		{
			Method:      "GET",
			Path:        "/categories/:id",
			Handler:     handler.GetCategory,
			Middlewares: []gin.HandlerFunc{authMw.Authorize()},
		},
		{
			Method:      "PUT",
			Path:        "/categories/:id",
			Handler:     handler.UpdateCategory,
			Middlewares: []gin.HandlerFunc{authMw.Authorize(), authMw.RequireRole("super_admin")},
		},
		{
			Method:      "DELETE",
			Path:        "/categories/:id",
			Handler:     handler.DeleteCategory,
			Middlewares: []gin.HandlerFunc{authMw.Authorize(), authMw.RequireRole("super_admin")},
		},
		{
			Method:      "POST",
			Path:        "/categories/:id/attributes",
			Handler:     handler.CreateAttribute,
			Middlewares: []gin.HandlerFunc{authMw.Authorize(), authMw.RequireRole("super_admin")},
		},
		{
			Method:      "GET",
			Path:        "/categories/:id/attributes",
			Handler:     handler.ListAttributes,
			Middlewares: []gin.HandlerFunc{authMw.Authorize()},
		},
		{
			Method:      "GET",
			Path:        "/categories/:id/attributes/:name",
			Handler:     handler.GetAttribute,
			Middlewares: []gin.HandlerFunc{authMw.Authorize()},
		},
		{
			Method:      "DELETE",
			Path:        "/categories/attributes/:attrID",
			Handler:     handler.DeleteAttribute,
			Middlewares: []gin.HandlerFunc{authMw.Authorize(), authMw.RequireRole("super_admin")},
		},
	}
	routes.RegisterRoute(grp, rs)
}
