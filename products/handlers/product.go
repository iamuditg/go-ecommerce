package handlers

import (
	"encoding/json"
	"github.com/iamuditg/models"
	"github.com/iamuditg/services"
	"log"
	"net/http"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

func (h *ProductHandler) GetProductHandler(writer http.ResponseWriter, request *http.Request) {

	search := request.URL.Query().Get("search")

	products, err := h.ProductService.GetProduct(search)
	if err != nil {
		log.Printf("failed to retreive products from service: %v", err)
		http.Error(writer, "failed to retrieve products", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(products)
	if err != nil {
		log.Printf("failed to marshal products to JSON: %v", err)
		http.Error(writer, "failed to retrieve products", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}

func (h *ProductHandler) CreateProductHandler(writer http.ResponseWriter, request *http.Request) {
	// Parse the JSON request body into a handler object
	var product models.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		http.Error(writer, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = product.Validation()
	if err != nil {
		http.Error(writer, "failed to validate the request body", http.StatusBadRequest)
		return
	}

	// Call the service to create the product
	createdProduct, err := h.ProductService.CreateProductService(&product)
	if err != nil {
		http.Error(writer, "Failed to create product", http.StatusInternalServerError)
		return
	}

	// Convert the Created Product to JSON
	createdProductJSON, err := json.Marshal(createdProduct)
	if err != nil {
		http.Error(writer, "Failed to marshal created product data", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(createdProductJSON)
}
