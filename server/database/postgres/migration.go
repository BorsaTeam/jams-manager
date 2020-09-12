package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/BorsaTeam/jams-manager/server/database"
)

var migrationPath = os.Getenv("DATABASE_MIGRATION_PATH")

var _ database.Migration = Migration{}

type Migration struct {
}

func NewMigration() Migration {
	return Migration{}
}

func (mi Migration) Apply() {
	m, err := migrate.New(migrationPath, databaseUrl())
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func databaseUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslMode)
}
