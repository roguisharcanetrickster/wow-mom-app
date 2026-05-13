package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFileName    = "wowmom.db"
	schemaFilePath = "./database/schema.sql"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	fmt.Println("Current working directory:", wd)

	// Construct absolute path for the database file
	dbPath := filepath.Join(wd, dbFileName)
	fmt.Println("Database path:", dbPath)

	// Initialize database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Ping the database to ensure connection is established
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Read and execute schema.sql
	schema, err := os.ReadFile(schemaFilePath)
	if err != nil {
		log.Fatalf("Failed to read schema file %s: %v", schemaFilePath, err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		log.Fatalf("Failed to execute schema from %s: %v", schemaFilePath, err)
	}

	log.Println("Database initialized and schema applied successfully.")
	fmt.Println("Database setup complete.")
	os.Exit(0)
}
