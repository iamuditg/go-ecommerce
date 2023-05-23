package handlers

import (
	"encoding/json"
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
	products, err := h.ProductService.GetProduct()
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
