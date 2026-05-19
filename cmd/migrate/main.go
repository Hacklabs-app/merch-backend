package main

import (
	"log"
	"os"

	"github.com/Hacklabs-app/merch-backend/internal/repository/postgres"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	db, err := postgres.Connect(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = postgres.RunMigrations(db, "migrations")
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
