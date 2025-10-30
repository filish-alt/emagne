package routes

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func RegisterRoute(
	grp *gin.RouterGroup,
	routes []Router,
) {
	for _, route := range routes {
		
		handlers := append(route.Middlewares, route.Handler)
		// Register the route with the specified method
		switch route.Method {
		case "GET":
			grp.GET(route.Path, handlers...)
		case "POST":
			grp.POST(route.Path, handlers...)
		case "PUT":
			grp.PUT(route.Path, handlers...)
		case "DELETE":
			grp.DELETE(route.Path, handlers...)
		case "PATCH":
			grp.PATCH(route.Path, handlers...)
		default:
			grp.Any(route.Path, handlers...)
		}
	}
}