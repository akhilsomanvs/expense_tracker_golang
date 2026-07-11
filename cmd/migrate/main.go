package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	m, err := migrate.New(
		"file://migrations",
		dsn,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil {

		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("Database is already up-to-date.")
			return
		}

		log.Fatal(err)
	}

	fmt.Println("Migrations completed successfully.")
}
