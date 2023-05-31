package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Migrations -> Migration Struct
type Migrations struct {
	env Env
}

// NewMigrations -> return new Migrations struct
func NewMigrations(
	env Env,
) Migrations {
	return Migrations{
		env: env,
	}
}

func (mm Migrations) Migrate() error {
	fmt.Println("migrating our database")
	db, _ := sql.Open("postgres", "postgres://admin:password123@postgresdb_container:5432/golang_gin_jwt?slmode=enable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("fail to create postgres driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrations : %w", err)
		}
	}
	fmt.Println("successfully migrate the database")
	return nil
}
