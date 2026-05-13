package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFileName    = "wow-mom-app/wowmom.db"
	schemaFilePath = "./wow-mom-app/database/schema.sql"
	port          = ":30111" // As requested, port 30111
)

var db *sql.DB

func main() {
	// Print current working directory for debugging
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	fmt.Println("Current working directory:", wd)

	// Initialize database
	initDB()
	defer db.Close()

	// Set up router
	router := mux.NewRouter()

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./wow-mom-app/static"))))

	// Routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	// TODO: Add more routes for authentication, mothers, leaders, groups, applications, notifications

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, router))
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

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
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	// For now, a simple response. Later, this will render an HTML template.
	fmt.Fprintf(w, "Welcome to WOW Mom App!")
}
