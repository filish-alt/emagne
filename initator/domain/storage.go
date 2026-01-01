package domain

import (
    "github.com/filagot/emagne/internal/database/persistancedb"
    "github.com/filagot/emagne/internal/storage"
    "github.com/filagot/emagne/internal/storage/auth"
    catattr "github.com/filagot/emagne/internal/storage/categoryattribute"
    itemcat "github.com/filagot/emagne/internal/storage/itemcategory"
    txstore "github.com/filagot/emagne/internal/storage/transaction"
    txattr "github.com/filagot/emagne/internal/storage/transactionattribute"
)

type PersistanceLayer struct {
    AuthStorage storage.AuthStorage
    ItemCategory storage.ItemCategoryStorage
    CategoryAttribute storage.CategoryAttributeStorage
    Transaction storage.TransactionStorage
    TransactionAttribute storage.TransactionItemAttributeStorage
}

func InitPersistance(db persistancedb.PersistenceDB) *PersistanceLayer {
    return &PersistanceLayer{
        AuthStorage: auth.NewAuthStorage(db.Queries),
        ItemCategory: itemcat.New(db.Queries),
        CategoryAttribute: catattr.New(db.Queries),
        Transaction: txstore.New(db.Queries),
        TransactionAttribute: txattr.New(db.Queries),
    }
}
