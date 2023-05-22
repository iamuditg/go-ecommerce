package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/iamuditg/models"
	"log"
	"net/http"
)

type ProductConfig struct {
	DB *sql.DB
}

func (p ProductConfig) ProductHandler(writer http.ResponseWriter, request *http.Request) {
	rows, err := p.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Printf("failed to query products: %v", err)
		http.Error(writer, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the products
	products := make([]models.Product, 0)

	// Iterate over the rows and populate the products slice
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Printf("failed to scan product row: %v", err)
			http.Error(writer, "failed to retrieve products", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("error iterating over product rows: %v", err)
		http.Error(writer, "failed to retrieve products", http.StatusInternalServerError)
		return
	}

	// Convert the products slice to JSON
	jsonData, err := json.Marshal(products)
	if err != nil {
		log.Printf("failed to marshal products to JSON: %v", err)
		http.Error(writer, "failed to retrieve products", http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)

}
