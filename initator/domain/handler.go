package domain

import (
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/handler/rest/auth"
)

type Handler struct {
    AuthHandler rest.AuthHandler
}

func InitHandler(modules *Module) *Handler {
	return &Handler{
        AuthHandler: auth.Init(modules.UserAuth),
	}
}