package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamuditg/config"
	"github.com/iamuditg/handlers"
	"github.com/iamuditg/repository"
	"github.com/iamuditg/services"
	"github.com/iamuditg/utils"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
	connStr := config.GetDBConnectionString(cfg)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Execute the SQL file
	if err := utils.ExecuteSQLFile(db, "products/schema.sql"); err != nil {
		log.Fatalf("failed to execute SQL file: %v", err)
	}

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Create instance of the repository service and handler layers
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Register the Http Routes for Products
	router.HandleFunc("/products", productHandler.GetProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", productHandler.CreateProductHandler).Methods(http.MethodPost)

	//Start the server
	port := viper.GetString("port")
	if port == "" {
		port = "8000"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
