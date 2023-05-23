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

func (s *ProductService) GetProduct() ([]models.Product, error) {
	products, err := s.ProductRepo.GetProduct()
	if err != nil {
		log.Printf("failed to retreive products from repository: %v", err)
		return nil, err
	}
	return products, nil
}
