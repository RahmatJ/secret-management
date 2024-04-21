package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"secret-management/app"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp() {
	env, err := app.GetEnv()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file. e: %+v", err))
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DbUser, env.DbPass, env.DbHost, env.DbPort, env.DbName)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cannot open connection to database")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("cannot create postgres driver", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration/scripts",
		env.DbName,
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}
