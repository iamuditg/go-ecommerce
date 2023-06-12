package repository

import (
	"database/sql"
	"github.com/iamuditg/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (c *CategoryRepository) CreateCategory(category models.Category) {

}
