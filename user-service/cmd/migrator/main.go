package main

import (
	"errors"
	"flag"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/protomem/socnet/user-service/assets"
)

const (
	_migrateUp   = "up"
	_migrateDown = "down"
	_migrateDrop = "drop"
)

var (
	_migrationsAction = flag.String("migrations-action", _migrateUp, "migrations action")
	_databaseURL      = flag.String("database-url", "", "database url")
)

func init() {
	flag.Parse()

	if *_databaseURL == "" {
		*_databaseURL = os.Getenv("DATABASE_URL")
	}
}

func main() {
	var err error

	if *_migrationsAction == "" {
		panic("migrations action is required")
	}
	if *_databaseURL == "" {
		panic("database url is required")
	}

	source, err := iofs.New(assets.Assets, "migrations")
	if err != nil {
		panic("failed to create migrations source: " + err.Error())
	}

	m, err := migrate.NewWithSourceInstance("iofs", source, *_databaseURL)
	if err != nil {
		panic("failed to create migrations instance: " + err.Error())
	}
	defer func() { _, _ = m.Close() }()

	switch *_migrationsAction {
	case _migrateUp:
		err = m.Up()
	case _migrateDown:
		err = m.Down()
	case _migrateDrop:
		err = m.Drop()
	default:
		panic("migrations action must be 'up' or 'down' or 'drop'")
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic("failed to run " + *_migrationsAction + " migrations: " + err.Error())
	}
}
