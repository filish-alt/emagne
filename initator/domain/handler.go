package domain

import (
    "github.com/filagot/emagne/internal/handler/rest"
    "github.com/filagot/emagne/internal/handler/rest/auth"
    catrest "github.com/filagot/emagne/internal/handler/rest/category"
    txrest "github.com/filagot/emagne/internal/handler/rest/transaction"
)

type Handler struct {
    AuthHandler rest.AuthHandler
    CategoryHandler rest.CategoryHandler
    TransactionHandler rest.TransactionHandler
}

func InitHandler(modules *Module) *Handler {
    return &Handler{
        AuthHandler: auth.Init(modules.UserAuth),
        CategoryHandler: catrest.Init(modules.Category),
        TransactionHandler: txrest.Init(modules.Transaction),
    }
}
