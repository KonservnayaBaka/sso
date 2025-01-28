package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"os"
	"sso/internal/config"
)

func main() {
	var migrationsPath, migrationsTable string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	godotenv.Load("env/.env")
	cfg := config.MustLoad()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User, os.Getenv("DB_PASSWORD"),
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.DBName, cfg.Database.SslMode,
	)

	m, err := migrate.New(
		"file://"+migrationsPath, dsn)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
