package domain

import (
	"github.com/filagot/emagne/internal/database/persistancedb"
	"github.com/filagot/emagne/internal/storage"
	"github.com/filagot/emagne/internal/storage/auth"
)

type PersistanceLayer struct {
	AuthStorage storage.AuthStorage
}

func InitPersistance(db persistancedb.PersistenceDB) *PersistanceLayer {
	return &PersistanceLayer{
		AuthStorage: auth.NewAuthStorage(db.Queries),
	}
}