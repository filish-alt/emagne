package persistancedb

import (

    database "github.com/filagot/emagne/internal/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PersistenceDB struct {
    *database.Queries
	Pool      *pgxpool.Pool
}

func New(pool *pgxpool.Pool) PersistenceDB {
    return PersistenceDB{
		Queries: database.New(pool),
		Pool: pool,

}
}

func (db *PersistenceDB) Close() error {
	db.Close()
	return nil
}

