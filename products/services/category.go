package services

import (
	"github.com/iamuditg/models"
	"github.com/iamuditg/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService(category *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: category}
}

func (c *CategoryService) CreateCategoryService(category models.Category) models.Category {
	return models.Category{}
}
