package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/morshedulmunna/go-curd/config"
)

func main() {
	config.ConnectDB()

	driver, err := postgres.WithInstance(config.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Check migration argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration command: up / down")
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Database migrated up successfully!")
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database rolled back successfully!")
	default:
		log.Fatal("Invalid command! Use 'up' or 'down'.")
	}
}
