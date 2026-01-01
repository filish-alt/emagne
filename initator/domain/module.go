package domain

import (
    "github.com/filagot/emagne/internal/config"
    "github.com/filagot/emagne/internal/module"
    "github.com/filagot/emagne/internal/module/auth"
    catmod "github.com/filagot/emagne/internal/module/category"
    txmod "github.com/filagot/emagne/internal/module/transaction"
)

type Module struct {
    UserAuth module.AuthModule
    Category module.Catagory
    Transaction txmod.Module
}

func InitModule(persistance *PersistanceLayer, cfg *config.Config) *Module {
    return &Module{
        UserAuth: auth.NewAuthModule(persistance.AuthStorage, cfg),
        Category: catmod.New(persistance.ItemCategory, persistance.CategoryAttribute),
        Transaction: txmod.New(persistance.Transaction, persistance.TransactionAttribute),
    }
}
