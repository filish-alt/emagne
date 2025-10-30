package foundation

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitDB(url string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal("Unable to parse DATABASE_URL", err)
	}
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal(fmt.Sprintf("Failed to ping database: %v", err))
	}
	return conn
}