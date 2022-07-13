package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) Migrate() error {
	fmt.Println("migrating the database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not migrate the database: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("could not migrate the database: %w", err)
	}

	fmt.Println("successfully migrated the database")
	return nil
}
