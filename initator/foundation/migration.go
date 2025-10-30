package foundation

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitiateMigration(path, conn string) *migrate.Migrate {
	m, err := migrate.New(fmt.Sprintf("file://%s", path), conn)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	m.LockTimeout = 10 * time.Minute
	return m
}

func UpMigration(m *migrate.Migrate) {
	defer func() {
		if sourceErr, dbErr := m.Close(); sourceErr != nil || dbErr != nil {
			log.Fatalf("Failed to close migrate instance: sourceErr=%v, dbErr=%v", sourceErr, dbErr)
		}
	}()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration up failed: %v", err)
	}
}

func DownMigration(m *migrate.Migrate) {
	defer func() {
		if sourceErr, dbErr := m.Close(); sourceErr != nil || dbErr != nil {
			log.Fatalf("Failed to close migrate instance: sourceErr=%v, dbErr=%v", sourceErr, dbErr)
		}
	}()
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration down failed: %v", err)
	}
}

func ForceMigration(m *migrate.Migrate, version int) {
	defer func() {
		if sourceErr, dbErr := m.Close(); sourceErr != nil || dbErr != nil {
			log.Fatalf("Failed to close migrate instance: sourceErr=%v, dbErr=%v", sourceErr, dbErr)
		}
	}()
	if err := m.Force(version); err != nil {
		log.Fatalf("Migration force failed: %v", err)
	}
}