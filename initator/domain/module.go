package domain

import (
	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/module"
	"github.com/filagot/emagne/internal/module/auth"
)

type Module struct {
	UserAuth module.AuthModule
}

func InitModule(persistance *PersistanceLayer, cfg *config.Config) *Module {
	return &Module{
		UserAuth: auth.NewAuthModule(persistance.AuthStorage, cfg),
	}
}