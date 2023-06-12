package services

import (
	"github.com/iamuditg/models"
	"github.com/iamuditg/repository"
	"log"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

func (s *ProductService) GetProduct(name string) ([]models.Product, error) {
	products, err := s.ProductRepo.GetProduct(name)
	if err != nil {
		log.Printf("failed to retreive products from repository: %v", err)
		return nil, err
	}
	return products, nil
}

func (s *ProductService) CreateProductService(product *models.Product) (*models.Product, error) {
	// Call the repository to create the product
	createdProduct, err := s.ProductRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}
