package domain

import (
	"github.com/filagot/emagne/internal/handler/rest"
)

type Handler struct {
    AuthHandler rest.AuthHandler
}

func InitHandler(modules *Module) *Handler {
	return &Handler{
        AuthHandler: rest.NewHandler(modules.UserAuth),
	}
}