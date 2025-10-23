package persistancedb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/filagot/emagne/internal/database"
	_ "github.com/lib/pq"
)

type PersistenceDB struct {
	*database.Queries
	DB *sql.DB
}

func New(dbSource string) (*PersistenceDB, error) {
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PersistenceDB{
		Queries: database.New(db),
		DB:      db,
	}, nil
}

func (p *PersistenceDB) Close() error {
	return p.DB.Close()
}

func (p *PersistenceDB) Ping(ctx context.Context) error {
	return p.DB.PingContext(ctx)
}
