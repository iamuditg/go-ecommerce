package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamuditg/config"
	"github.com/iamuditg/handlers"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to lead config: %v", err)
	}

	// Establish database connection
	connStr := getDBConnectionString(cfg)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Execute the SQL file
	//if err := executeSQLFile(db, "products/schema.sql"); err != nil {
	//	log.Fatalf("failed to execute SQL file: %v", err)
	//}

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define the product endpoint
	handler := handlers.ProductConfig{db}
	router.HandleFunc("/products", handler.ProductHandler).Methods(http.MethodGet)

	//Start the server
	port := viper.GetString("port")
	if port == "" {
		port = "8000"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func executeSQLFile(db *sql.DB, filename string) error {
	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Failed to read SQL file: %v", err)
	}

	// Execute the SQL statements
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("failed to execute SQL statements: %v", err)
	}

	return nil
}

func getDBConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
}
